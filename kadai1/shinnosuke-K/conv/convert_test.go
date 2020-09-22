package conv

import (
	"testing"

	"github.com/shinnosuke-K/gopherdojo-studyroom/kadai1/shinnosuke-K/file"
)

func Test_checkOpt(t *testing.T) {
	type args struct {
		ex string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "拡張子が違う（小文字）",
			args:    args{ex: "d"},
			wantErr: true,
		},
		{
			name:    "拡張子が違う（大文字）",
			args:    args{ex: "D"},
			wantErr: true,
		},
		{
			name:    "拡張子が正しい（png）",
			args:    args{ex: "png"},
			wantErr: false,
		},
		{
			name:    "拡張子が正しい（jpg）",
			args:    args{ex: "jpg"},
			wantErr: false,
		},
		{
			name:    "拡張子が正しい（jpeg）",
			args:    args{ex: "jpeg"},
			wantErr: false,
		},
		{
			name:    "拡張子が正しい（gif）",
			args:    args{ex: "gif"},
			wantErr: false,
		},
		{
			name:    "拡張子が正しい（PNG）",
			args:    args{ex: "PNG"},
			wantErr: false,
		},
		{
			name:    "拡張子が正しい（JPG）",
			args:    args{ex: "JPG"},
			wantErr: false,
		},
		{
			name:    "拡張子が正しい（JPEG）",
			args:    args{ex: "JPEG"},
			wantErr: false,
		},
		{
			name:    "拡張子が正しい（GIF）",
			args:    args{ex: "GIF"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkOpt(tt.args.ex); (err != nil) != tt.wantErr {
				t.Errorf("checkOpt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_convert(t *testing.T) {
	type args struct {
		afterExt string
		f        file.File
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// pngへ変換
		{
			name: "拡張子が連続（jpeg）",
			args: args{
				afterExt: "png",
				f: struct {
					Dir       string
					Name      string
					Extension string
				}{Dir: "../dummyImg/jpeg", Name: "1.jpeg.jpeg", Extension: ".jpeg"},
			},
			wantErr: false,
		},
		{
			name: "拡張子が連続（jpg）",
			args: args{
				afterExt: "png",
				f: struct {
					Dir       string
					Name      string
					Extension string
				}{Dir: "../dummyImg/jpg", Name: "1.jpg.jpg", Extension: ".jpg"}},
			wantErr: false,
		},
		{
			name: "拡張子が連続（gif）",
			args: args{
				afterExt: "png",
				f: struct {
					Dir       string
					Name      string
					Extension string
				}{Dir: "../dummyImg/", Name: "1.gif.gif", Extension: ".gif"}},
			wantErr: false,
		},

		// jpgへ変換
		{
			name: "拡張子が連続（png）",
			args: args{
				afterExt: "jpg",
				f: struct {
					Dir       string
					Name      string
					Extension string
				}{Dir: "../dummyImg/png", Name: "1.png.png", Extension: ".png"}},
			wantErr: false,
		},
		{
			name: "拡張子が連続（gif）",
			args: args{
				afterExt: "jpg",
				f: struct {
					Dir       string
					Name      string
					Extension string
				}{Dir: "../dummyImg/", Name: "1.gif.gif", Extension: ".gif"}},
			wantErr: false,
		},

		// gifへ変換
		{
			name: "拡張子が連続（jpeg）",
			args: args{
				afterExt: "gif",
				f: struct {
					Dir       string
					Name      string
					Extension string
				}{Dir: "../dummyImg/jpeg", Name: "1.jpeg.jpeg", Extension: ".jpeg"},
			},
			wantErr: false,
		},
		{
			name: "拡張子が連続（jpg）",
			args: args{
				afterExt: "gif",
				f: struct {
					Dir       string
					Name      string
					Extension string
				}{Dir: "../dummyImg/jpg", Name: "1.jpg.jpg", Extension: ".jpg"}},
			wantErr: false,
		},
		{
			name: "拡張子が連続（png）",
			args: args{
				afterExt: "gif",
				f: struct {
					Dir       string
					Name      string
					Extension string
				}{Dir: "../dummyImg/png", Name: "1.png.png", Extension: ".png"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := convert(tt.args.afterExt, tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("convert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
