package gorbit

import (
	"path"
	"path/filepath"
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

func Pathname(basePath string) string {
	basePath = filepath.Base(basePath)
	for i := len(basePath) - 1; i >= 0; i-- {
		if basePath[i] == '.' {
			return basePath[:i]
		}
	}
	return ""
}
