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
			name: "when input valid md5",
			fields: fields{
				stdout: new(bytes.Buffer),
				stderr: new(bytes.Buffer),
			},
			args: args{
				args: []string{
					"chksum",
					"testdata/testdata.txt",
					"34bbf27b28c63465e465145468f44679",
				},
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
			name: "when input valid sha1",
			fields: fields{
				stdout: new(bytes.Buffer),
				stderr: new(bytes.Buffer),
			},
			args: args{
				args: []string{
					"chksum",
					"testdata/testdata.txt",
					"cb8d0af2d9b01e20d0ce0c070fbdac38f499e43e",
				},
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
			name: "when input valid sha256",
			fields: fields{
				stdout: new(bytes.Buffer),
				stderr: new(bytes.Buffer),
			},
			args: args{
				args: []string{
					"chksum",
					"testdata/testdata.txt",
					"a9b28957d417ef3d18180bbf7c220916d45e9250845569c05fb2e79b42b7d750",
				},
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
			name: "when input valid sha512",
			fields: fields{
				stdout: new(bytes.Buffer),
				stderr: new(bytes.Buffer),
			},
			args: args{
				args: []string{
					"chksum",
					"testdata/testdata.txt",
					"d7598bc31339787f29bb90d633ab72f97a73b0126edaaf68831f76d4b5afff2635be1ea4c695438ddf3f47d018c67823cd182fdfdb75dac6f32e7aaba513cb25",
				},
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
			fields: fields{
				stdout: new(bytes.Buffer),
				stderr: new(bytes.Buffer),
			},
			args: args{
				args: []string{
					"chksum",
					"testdata/testdata.txt",
					"570462a0a1725a0d0b41b443fe661d2b",
				},
			},
			want:    "",
			wantErr: true,
		},
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
