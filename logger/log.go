package logger

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
	f *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  string
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR"}
	SavePath   = ".logs/"
	FileExt    = "log"
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
)

func Debug(v ...interface{}) {
	file := setPrefix(DEBUG)
	defer file.Close()
	logger.Println(v...)
}

func Info(v ...interface{}) {
	file := setPrefix(INFO)
	defer file.Close()
	logger.Println(v...)
}

func Warn(v ...interface{}) {
	file := setPrefix(WARNING)
	defer file.Close()
	logger.Println(v...)
}

func Error(v ...interface{}) {
	file := setPrefix(ERROR)
	defer file.Close()
	logger.Println(v...)
}

func setPrefix(level Level) *os.File {
	filePath := getLogFileFullPath(level)
	f, _ = openLogFile(filePath)
	logger = log.New(f, DefaultPrefix, log.LstdFlags)
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s] %s:%d ", levelFlags[level], filepath.Clean(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s] ", levelFlags[level])
	}
	if level != INFO {
		logger.SetPrefix(logPrefix)
	}
	return f
}

func getLogFilePath() string {
	return fmt.Sprintf("%s/", SavePath)
}

func getLogFileFullPath(level Level) string {
	suffixPath := fmt.Sprintf("%s.%s", levelFlags[level]+"-"+time.Now().Format("2006-01-02"), FileExt)
	return fmt.Sprintf("%s%s", getLogFilePath(), suffixPath)
}

func openLogFile(filePath string) (*os.File, error) {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		dir, _ := os.Getwd()
		_ = os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	case os.IsPermission(err):
		return nil, err
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return handle, nil
}
