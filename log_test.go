package gorbit

import "testing"

func TestDebug(t *testing.T) {
	Debug(`test.me`)
	Info(`test.me`)
	Info(`test.me`)
	Warn(`test.me`)
	Error(`test.me`)
}
