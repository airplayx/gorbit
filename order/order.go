package order

import (
	"github.com/airplayx/gorbit/char"
	"strconv"
	"strings"
	"time"
)

func SnCode(sum int) string {
	str := time.Now().Format("06")
	days := strconv.Itoa(char.DaysInYear())
	count := len(days)
	if count < 3 {
		days = strings.Repeat("0", 3-count) + days
	}
	str += days
	sum = sum - 5
	if sum < 1 {
		sum = 5
	}
	result := strconv.FormatInt(time.Now().UnixNano(), 10)
	count = len(result)
	if count < sum {
		result = strings.Repeat("0", sum-count) + result
	}
	str += result
	return str
}
