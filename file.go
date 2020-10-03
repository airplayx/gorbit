package gorbit

import (
	"path"
	"path/filepath"
	"strings"
)

func Ext(fileName string, allows []string) bool {
	ext := path.Ext(fileName)
	for _, allowExt := range allows {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func CleanName(filePath string) string {
	filePath = filepath.Base(filePath)
	for i := len(filePath) - 1; i >= 0; i-- {
		if filePath[i] == '.' {
			return filePath[:i]
		}
	}
	return filePath
}
