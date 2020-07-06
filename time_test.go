package gorbit

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
