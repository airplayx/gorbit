package gorbit

import (
	"errors"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func RandomStr(l int, isNum bool) (s string, err error) {
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

func FileUpTime(file string) (int64, error) {
	f, err := os.Open(file)
	if err != nil {
		return time.Now().Unix(), err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return time.Now().Unix(), err
	}
	return fi.ModTime().Unix(), nil
}

func SetVersion(ver string, upTime time.Time) string {
	total := 0
	arr := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y, month, d := upTime.Date()
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
	return ver + upTime.Format("06") + days
}

func IsExistItem(key, array interface{}) bool {
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
