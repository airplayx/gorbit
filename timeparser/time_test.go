package timeparser

import (
	"testing"
	"time"
)

func TestDay0(t *testing.T) {
	t.Parallel()
	t.Log(time.Now())
	t.Log(Day0(-11))
	t.Log(Day0(0))
	t.Log(Day0(123))
}

func TestTimeDiff(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test timeDiff 0",
			args: args{
				t: time.Now().AddDate(-1, 0, 0),
			},
			want: "1年前",
		},
		{
			name: "test timeDiff 1",
			args: args{
				t: time.Now().Add(-time.Second * 60),
			},
			want: "1分钟前",
		},
		{
			name: "test timeDiff 1",
			args: args{
				t: time.Now().Add(-time.Second * 60 * 60),
			},
			want: "1小时前",
		},
		{
			name: "test timeDiff 1",
			args: args{
				t: time.Now().Add(-time.Second * 60 * 60 * 24),
			},
			want: "1天前",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.args.t); got != tt.want {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
