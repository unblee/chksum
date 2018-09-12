package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-colorable"
)

const (
	terminalColorRed       = "\x1b[31m"
	terminalColorGreen     = "\x1b[32m"
	terminalColorReset     = "\x1b[m"
	terminalCursorShow     = "\x1b[?25h"
	terminalCursorHide     = "\x1b[?25l"
	terminalClearRight     = "\x1b[K"
	terminalCarriageReturn = "\r"
)

const (
	exitCodeOK = iota
	exitCodeErr
)

const USAGE = `A small tool for checking md5/sha1/sha256/sha512 checksums of a file.
Usage:
  chksum <file> <checksum>
  chksum (-h | --help)
Options:
  -h --help     Show this message.
  -V --version  Show version.`

var (
	helpFlag         = flag.Bool("help", false, "")
	helpShortFlag    = flag.Bool("h", false, "")
	versionFlag      = flag.Bool("version", false, "")
	versionShortFlag = flag.Bool("V", false, "")
)

type cli struct {
	stdout, stderr io.Writer
}

func newCLI(stdout, stderr io.Writer) *cli {
	return &cli{
		stdout: stdout,
		stderr: stderr,
	}
}

func (cl *cli) run(args []string) int {
	ret, err := cl.main(args)
	if err != nil {
		fmt.Fprintln(cl.stderr, err.Error())
		return exitCodeErr
	}
	fmt.Fprintln(cl.stdout, ret)
	return exitCodeOK
}

func (cl *cli) main(args []string) (string, error) {
	flag.Parse()
	if *helpFlag || *helpShortFlag {
		return USAGE, nil
	}
	if *versionFlag || *versionShortFlag {
		return version, nil
	}

	if len(args[1:]) < 2 {
		errStr := fmt.Sprintf("please input arguments\n\n%s", USAGE)
		return "", errors.New(errStr)
	}

	md5hash := md5.New()
	sha1hash := sha1.New()
	sha256hash := sha256.New()
	sha512hash := sha512.New()
	h := io.MultiWriter(md5hash, sha1hash, sha256hash, sha512hash)

	targetFile := args[1]
	f, err := os.Open(targetFile)
	if err != nil {
		f.Close()
		return "", err
	}
	defer f.Close()

	fmt.Fprint(cl.stdout, terminalCursorHide)
	defer fmt.Fprint(cl.stdout, terminalCursorShow)
	fmt.Fprint(cl.stdout, "calculating checksum...")
	if _, err = io.Copy(h, f); err != nil {
		return "", err
	}
	fmt.Fprint(cl.stdout, terminalCarriageReturn)
	fmt.Fprint(cl.stdout, terminalClearRight)

	md5sum := hex.EncodeToString(md5hash.Sum(nil))
	sha1sum := hex.EncodeToString(sha1hash.Sum(nil))
	sha256sum := hex.EncodeToString(sha256hash.Sum(nil))
	sha512sum := hex.EncodeToString(sha512hash.Sum(nil))

	targetChecksum := args[2]
	switch targetChecksum {
	case md5sum:
		return fmt.Sprintf("%s%s:md5 OK!%s", terminalColorGreen, md5sum, terminalColorReset), nil
	case sha1sum:
		return fmt.Sprintf("%s%s:sha1 OK!%s", terminalColorGreen, sha1sum, terminalColorReset), nil
	case sha256sum:
		return fmt.Sprintf("%s%s:sha256 OK!%s", terminalColorGreen, sha256sum, terminalColorReset), nil
	case sha512sum:
		return fmt.Sprintf("%s%s:sha512 OK!%s", terminalColorGreen, sha512sum, terminalColorReset), nil
	default:
		return "", fmt.Errorf("%s%s:md5/sha1/sha256/sha512 NG!%s", terminalColorRed, sha512sum, terminalColorReset)
	}
}

func main() {
	stdout := colorable.NewColorableStdout()
	stderr := colorable.NewColorableStderr()
	cli := newCLI(stdout, stderr)
	os.Exit(cli.run(os.Args))
}
