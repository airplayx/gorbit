package gorbit

import (
	"path"
	"strings"
)

func CheckFileExt(fileName string, allows []string) bool {
	ext := path.Ext(fileName)
	for _, allowExt := range allows {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}
