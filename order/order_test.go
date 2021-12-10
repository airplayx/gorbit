package order

import "testing"

func TestSetOrderSn(t *testing.T) {
	t.Parallel()
	t.Log(SnCode(0))
	t.Log(SnCode(25))
	t.Log(SnCode(-1))
}
