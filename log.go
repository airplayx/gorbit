package gorbit

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger      *log.Logger
	logPrefix   = ""
	levelFlags  = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	logSavePath = ".logs/"
	logFileExt  = "log"
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Debug(v ...interface{}) {
	file := setPrefix(DEBUG)
	defer func() {
		logger.Println(v)
		_ = file.Close()
	}()
}

func Info(v ...interface{}) {
	file := setPrefix(INFO)
	defer func() {
		logger.Println(v)
		_ = file.Close()
	}()
}

func Warn(v ...interface{}) {
	file := setPrefix(WARNING)
	defer func() {
		logger.Println(v)
		_ = file.Close()
	}()
}

func Error(v ...interface{}) {
	file := setPrefix(ERROR)
	defer func() {
		logger.Println(v)
		_ = file.Close()
	}()
}

func Fatal(v ...interface{}) {
	file := setPrefix(FATAL)
	defer func() {
		logger.Println(v)
		_ = file.Close()
	}()
}

func setPrefix(level Level) *os.File {
	filePath := getLogFileFullPath(level)
	F = openLogFile(filePath)
	logger = log.New(F, DefaultPrefix, log.LstdFlags)

	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s] %s:%d ", levelFlags[level], filepath.Clean(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s] ", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
	return F
}

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
