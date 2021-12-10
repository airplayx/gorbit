package logger

import "testing"

func TestDebug(t *testing.T) {
	Debug(`test.me`)
	Warn(`test.me`)
	Info(`test.me`)
	Error(`test.me`)
}
