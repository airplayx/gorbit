package gorbit

import (
	"encoding/json"
	"time"
)

type MyTime time.Time

func (mt *MyTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	if err != nil {
		return err
	}
	*mt = MyTime(t)
	return nil
}

func (mt *MyTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len("2006-01-02 15:04:05")+2)
	b = append(b, '"')
	b = time.Time(*mt).AppendFormat(b, "2006-01-02 15:04:05")
	b = append(b, '"')
	return b, nil
}

func (mt *MyTime) Time() time.Time {
	return time.Time(*mt)
}

func (mt *MyTime) String() string {
	return time.Time(*mt).Format("2006-01-02 15:04:05")
}

func Day0(diffDay int) time.Time {
	t := time.Now()
	timeToday := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return timeToday.AddDate(0, 0, diffDay)
}
