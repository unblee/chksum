package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

func Test_cli_main(t *testing.T) {
	type fields struct {
		stdout io.Writer
		stderr io.Writer
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "when arguments insufficient",
			fields: fields{
				stdout: new(bytes.Buffer),
				stderr: new(bytes.Buffer),
			},
			args: args{
				args: []string{
					"chksum",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "when generating checksums",
			fields: fields{
				stdout: new(bytes.Buffer),
				stderr: new(bytes.Buffer),
			},
			args: args{
				args: []string{
					"chksum",
					"testdata/testdata.txt",
				},
			},
			want: `md5sum: 34bbf27b28c63465e465145468f44679  testdata/testdata.txt
sha1sum: cb8d0af2d9b01e20d0ce0c070fbdac38f499e43e  testdata/testdata.txt
sha256sum: a9b28957d417ef3d18180bbf7c220916d45e9250845569c05fb2e79b42b7d750  testdata/testdata.txt
sha512sum: d7598bc31339787f29bb90d633ab72f97a73b0126edaaf68831f76d4b5afff2635be1ea4c695438ddf3f47d018c67823cd182fdfdb75dac6f32e7aaba513cb25  testdata/testdata.txt
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &cli{
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
			}

			got, err := cl.main(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("cli.main() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// clean unnecessary strings
			r := strings.NewReplacer(
				"calculating checksum...", "",
				terminalCursorHide, "",
				terminalCursorShow, "",
				terminalCarriageReturn, "",
				terminalClearRight, "",
			)

			got = r.Replace(got)
			if got != tt.want {
				t.Errorf("cli.main() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_check(t *testing.T) {
	type args struct {
		target    string
		md5sum    string
		sha1sum   string
		sha256sum string
		sha512sum string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "when checking valid md5",
			args: args{
				target: "34bbf27b28c63465e465145468f44679",
				md5sum: "34bbf27b28c63465e465145468f44679",
			},
			want: fmt.Sprintf(
				"%s%s:md5 OK!%s",
				terminalColorGreen,
				"34bbf27b28c63465e465145468f44679",
				terminalColorReset,
			),
			wantErr: false,
		},
		{
			name: "when checking valid sha1",
			args: args{
				target:  "cb8d0af2d9b01e20d0ce0c070fbdac38f499e43e",
				sha1sum: "cb8d0af2d9b01e20d0ce0c070fbdac38f499e43e",
			},
			want: fmt.Sprintf(
				"%s%s:sha1 OK!%s",
				terminalColorGreen,
				"cb8d0af2d9b01e20d0ce0c070fbdac38f499e43e",
				terminalColorReset,
			),
			wantErr: false,
		},
		{
			name: "when checking valid sha256",
			args: args{
				target:    "a9b28957d417ef3d18180bbf7c220916d45e9250845569c05fb2e79b42b7d750",
				sha256sum: "a9b28957d417ef3d18180bbf7c220916d45e9250845569c05fb2e79b42b7d750",
			},
			want: fmt.Sprintf(
				"%s%s:sha256 OK!%s",
				terminalColorGreen,
				"a9b28957d417ef3d18180bbf7c220916d45e9250845569c05fb2e79b42b7d750",
				terminalColorReset,
			),
			wantErr: false,
		},
		{
			name: "when checking valid sha512",
			args: args{
				target:    "d7598bc31339787f29bb90d633ab72f97a73b0126edaaf68831f76d4b5afff2635be1ea4c695438ddf3f47d018c67823cd182fdfdb75dac6f32e7aaba513cb25",
				sha512sum: "d7598bc31339787f29bb90d633ab72f97a73b0126edaaf68831f76d4b5afff2635be1ea4c695438ddf3f47d018c67823cd182fdfdb75dac6f32e7aaba513cb25",
			},
			want: fmt.Sprintf(
				"%s%s:sha512 OK!%s",
				terminalColorGreen,
				"d7598bc31339787f29bb90d633ab72f97a73b0126edaaf68831f76d4b5afff2635be1ea4c695438ddf3f47d018c67823cd182fdfdb75dac6f32e7aaba513cb25",
				terminalColorReset,
			),
			wantErr: false,
		},
		{
			name: "when input invalid checksum",
			args: args{
				target:    "1f21495e59ed694f2f4ea5edc117a1bf5429b88fe5b039080044496ac202cf121da8ae70f95a31c8f02b2bf88cb583c06779b75ed050dfc4e29e711aa1a058c0",
				sha512sum: "d7598bc31339787f29bb90d633ab72f97a73b0126edaaf68831f76d4b5afff2635be1ea4c695438ddf3f47d018c67823cd182fdfdb75dac6f32e7aaba513cb25",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := check(tt.args.target, tt.args.md5sum, tt.args.sha1sum, tt.args.sha256sum, tt.args.sha512sum)
			if (err != nil) != tt.wantErr {
				t.Errorf("check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
