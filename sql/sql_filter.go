package sql

import (
	"log"
	"regexp"
	"strings"
)

var reStr = `(?:')|(?:\%)|(?:\\)|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`

func SafeString(matchStr string, exactly bool) string {
	re, err := regexp.Compile(reStr)
	if err != nil {
		log.Println(err)
		return ``
	}
	return re.ReplaceAllStringFunc(matchStr, func(dst string) string {
		if dst == `\` && !exactly {
			dst = strings.ReplaceAll(dst, dst, `\\\`+dst)
		} else {
			dst = strings.ReplaceAll(dst, dst, `\`+dst)
		}
		return dst
	})
}
