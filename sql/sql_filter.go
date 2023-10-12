package sql

import (
	"regexp"
	"strings"
)

const unsafeKeys = `(?:')|(?:\%)|(?:\\)|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`

func SafeString(matchStr string, exactly bool) string {
	ruler := regexp.MustCompile(unsafeKeys)
	return ruler.ReplaceAllStringFunc(matchStr, func(dst string) string {
		if dst == `\` && !exactly {
			return strings.ReplaceAll(dst, dst, `\\\`+dst)
		}
		return strings.ReplaceAll(dst, dst, `\`+dst)
	})
}
