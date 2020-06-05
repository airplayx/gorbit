package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
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
