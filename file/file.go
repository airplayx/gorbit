package file

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func Exist(fileName string, allows []string) bool {
	ext := path.Ext(fileName)
	for _, allowExt := range allows {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func UpdateTime(file string) (int64, error) {
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

func CleanName(filePath string) string {
	filePath = filepath.Base(filePath)
	for i := len(filePath) - 1; i >= 0; i-- {
		if filePath[i] == '.' {
			return filePath[:i]
		}
	}
	return filePath
}
