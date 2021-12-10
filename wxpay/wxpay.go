package wxpay

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func PaySign(params map[string]string, key string) string {
	var keys []string
	var sorted []string
	for k, v := range params {
		if k != "sign" && v != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		sorted = append(sorted, fmt.Sprintf("%s=%s", k, params[k]))
	}
	str := strings.Join(sorted, "&")
	str += "&key=" + key
	return fmt.Sprintf("%X", md5.Sum([]byte(str)))
}

func NonceStr(n uint) string {
	var value []byte
	chars := []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	m := len(chars)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := uint(0); i < n; i++ {
		value = append(value, chars[r.Intn(m)])
	}
	return string(value)
}
