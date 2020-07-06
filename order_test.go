package gorbit

import "testing"

func TestSetOrderSn(t *testing.T) {
	t.Parallel()
	t.Log(SetOrderSn(0))
	t.Log(SetOrderSn(25))
	t.Log(SetOrderSn(-1))
}

func TestDaysInYear(t *testing.T) {
	t.Parallel()
	t.Log(DaysInYear())
}
