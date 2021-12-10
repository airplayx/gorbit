package char

import (
	"testing"
	"time"
)

func TestRandomStr(t *testing.T) {
	t.Parallel()
	var testChar = []map[int]bool{
		{-1: false}, {-1: true},
		{0: false}, {0: true},
		{1: false}, {1: true},
		{99: false}, {99: true},
		{255: true}, {255: false},
	}
	for _, v := range testChar {
		for key, value := range v {
			result, err := Random(key, value)
			if err != nil {
				t.Logf("RandomStr fail: [%d , %t] => %s", key, value, err.Error())
				continue
			}
			t.Logf("RandomStr ok: [%d , %t] => %s", key, value, result)
		}
	}
}

func TestSetVersion(t *testing.T) {
	t.Parallel()
	var testVer = []map[string]time.Time{
		{"": time.Now()},
		{"test_": time.Now()},
	}
	for _, v := range testVer {
		for key, value := range v {
			result := Version(key, value)
			if result == "" {
				t.Errorf("RandomStr fail: [%s , %v] => ''", key, value)
			}
			t.Logf("RandomStr ok: [%s , %v] => %s", key, value, result)
		}
	}
}

func TestIsExistItem(t *testing.T) {
	type args struct {
		key   interface{}
		array interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "key one",
			args: args{
				key:   "aa",
				array: []string{"aa", "bb", "cc"},
			},
			want: true,
		},
		{
			name: "key two",
			args: args{
				key:   1,
				array: []string{"aa", "bb", "cc"},
			},
			want: false,
		},
		{
			name: "key three",
			args: args{
				key:   uint(1),
				array: []uint{1, 2, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exist(tt.args.key, tt.args.array); got != tt.want {
				t.Errorf("IsExistItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDaysInYear(t *testing.T) {
	t.Parallel()
	t.Log(DaysInYear())
}
