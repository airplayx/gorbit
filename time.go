package gorbit

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
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

func TimeDiff(t time.Time) (diffStr string) {
	var times = []float64{365 * 24 * 60 * 60, 24 * 60 * 60, 60 * 60, 60, 1}
	var units = []string{"年前", "天前", "小时前", "分钟前", "秒钟前"}
	diffTime := time.Now().Sub(t).Seconds()
	if diffTime <= 0 {
		return "刚刚"
	}
	for i, matTime := range times {
		if diffTime < matTime {
			continue
		}
		if temp := math.Floor(float64(diffTime / matTime)); temp > 0 {
			return fmt.Sprint(
				strconv.FormatFloat(temp, 'f', -1, 64),
				units[i],
			)
		}
		diffTime = math.Mod(diffTime, matTime)
	}
	return
}
