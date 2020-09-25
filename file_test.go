package gorbit

import (
	"os"
	"testing"
)

func TestCheckFileExt(t *testing.T) {
	t.Parallel()
	t.Log(CheckFileExt(os.Args[0], []string{".exe", ""}))
}

func TestPathname(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pathname(tt.args.basePath); got != tt.want {
				t.Errorf("Pathname() = %v, want %v", got, tt.want)
			}
		})
	}
}
