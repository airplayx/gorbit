package gorbit

import (
	"os"
	"testing"
)

func TestCheckFileExt(t *testing.T) {
	t.Parallel()
	t.Log(CheckFileExt(os.Args[0], []string{".exe", ""}))
}
