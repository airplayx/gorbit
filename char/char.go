package char

import (
	"errors"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Random(l int, isNum bool) (s string, err error) {
	if l <= 0 {
		return "", errors.New("the length must > 0")
	}
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if isNum {
		str = "0123456789"
	}
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result), nil
}

func Version(ver string, t time.Time) string {
	total := 0
	arr := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y, month, d := t.Date()
	m := int(month)
	for i := 0; i < m-1; i++ {
		total = total + arr[i]
	}
	if (y%400 == 0 || (y%4 == 0 && y%100 != 0)) && m > 2 {
		total = total + d + 1
	} else {
		total = total + d
	}
	days := strconv.Itoa(total)
	count := len(days)
	if count < 3 {
		days = strings.Repeat("0", 3-count) + days
	}
	return ver + t.Format("06") + days
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

func Exist(key, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(key, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}
