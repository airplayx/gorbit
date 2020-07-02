package gorbit

import (
	"strconv"
	"strings"
	"time"
)

func SetOrderSn(sum int) string {
	str := time.Now().Format("06")
	days := strconv.Itoa(DaysInYear())
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

func DaysInYear() int {
	now := time.Now()
	total := 0
	arr := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y, month, d := now.Date()
	m := int(month)
	for i := 0; i < m-1; i++ {
		total = total + arr[i]
	}
	if (y%400 == 0 || (y%4 == 0 && y%100 != 0)) && m > 2 {
		total = total + d + 1
	} else {
		total = total + d
	}
	return total
}
