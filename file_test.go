package gorbit

import (
	"testing"
)

func TestFileCleanName(t *testing.T) {
	type args struct {
		basePath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check path name1",
			args: args{
				basePath: "css.html",
			},
			want: "css",
		},
		{
			name: "check path name2",
			args: args{
				basePath: "/////////////css.html",
			},
			want: "css",
		},
		{
			name: "check path name3",
			args: args{
				basePath: "//ccc//..///css.....html",
			},
			want: "css....",
		},
		{
			name: "check path name4",
			args: args{
				basePath: "...../css_(:з」∠)_.html",
			},
			want: "css_(:з」∠)_",
		},
		{
			name: "check path name5",
			args: args{
				basePath: "....",
			},
			want: "...",
		},
		{
			name: "check path name6",
			args: args{
				basePath: "",
			},
			want: "",
		},
		{
			name: "check path name7",
			args: args{
				basePath: "css",
			},
			want: "css",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanName(tt.args.basePath); got != tt.want {
				t.Errorf("CleanName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckFileExt(t *testing.T) {
	type args struct {
		fileName string
		allows   []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestCheckFileExt 1",
			args: args{
				fileName: "...dds.exe",
				allows:   []string{".exe", ".jpg"},
			},
			want: true,
		},
		{
			name: "TestCheckFileExt 2",
			args: args{
				fileName: "....exe",
				allows:   []string{".exe", ".jpg"},
			},
			want: true,
		},
		{
			name: "TestCheckFileExt 3",
			args: args{
				fileName: "....",
				allows:   []string{".exe", ".jpg"},
			},
			want: false,
		},
		{
			name: "TestCheckFileExt 4",
			args: args{
				fileName: "......_(:з」∠)_",
				allows:   []string{"._(:з」∠)_", ".jpg"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ext(tt.args.fileName, tt.args.allows); got != tt.want {
				t.Errorf("Ext() = %v, want %v", got, tt.want)
			}
		})
	}
}
