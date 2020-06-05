package gorbit

import "time"

func Day0(diffDay int) time.Time {
	t := time.Now()
	timeToday := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return timeToday.AddDate(0, 0, diffDay)
}
