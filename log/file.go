package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	logSavePath = ".logs/"
	logFileExt  = "log"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s/", logSavePath)
}

func getLogFileFullPath(level Level) string {
	suffixPath := fmt.Sprintf("%s.%s", levelFlags[level]+"-"+time.Now().Format("2006-01-02"), logFileExt)
	return fmt.Sprintf("%s%s", getLogFilePath(), suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		dir, _ := os.Getwd()
		_ = os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	return handle
}
