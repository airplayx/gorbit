package gorbit

import (
	"os"
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
			result, err := RandomStr(key, value)
			if err != nil {
				t.Logf("RandomStr fail: [%d , %t] => %s", key, value, err.Error())
				continue
			}
			t.Logf("RandomStr ok: [%d , %t] => %s", key, value, result)
		}
	}
}

func TestFileUpTime(t *testing.T) {
	t.Parallel()
	t.Log(FileUpTime(os.Args[0]))
	t.Log(FileUpTime(""))
}

func TestSetVersion(t *testing.T) {
	t.Parallel()
	var testVer = []map[string]time.Time{
		{"": time.Now()},
		{"test_": time.Now()},
	}
	for _, v := range testVer {
		for key, value := range v {
			result := SetVersion(key, value)
			if result == "" {
				t.Errorf("RandomStr fail: [%s , %v] => ''", key, value)
			}
			t.Logf("RandomStr ok: [%s , %v] => %s", key, value, result)
		}
	}
}

func TestHaveFound(t *testing.T) {
	t.Parallel()
	t.Log(HaveFound([]string{}, ""))
	t.Log(HaveFound([]string{""}, ""))
	t.Log(HaveFound([]string{"1"}, "1"))
	t.Log(HaveFound([]string{"1"}, "2"))
	t.Log(HaveFound([]string{""}, "2"))
	t.Log(HaveFound([]string{"1"}, ""))
}
