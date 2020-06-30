package gorbit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
)

type err struct {
	Message string      `json:"message,omitempty"`
	File    string      `json:"file,omitempty"`
	Line    int         `json:"line,omitempty"`
	Func    string      `json:"func,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (e *err) Error() string {
	var buf bytes.Buffer
	if e.File != "" {
		buf.WriteString(fmt.Sprintf("[%s - %s : %d] ", e.File, e.Func, e.Line))
	}
	buf.WriteString(e.Message)
	return buf.String()
}

func (e err) Format(args ...interface{}) *err {
	e.Message = fmt.Sprintf(e.Message, args...)
	return &e
}

func (e err) Location() *err {
	pc, file, line, ok := runtime.Caller(1)
	if ok == false {
		file = "???"
		line = -1
	}
	f := runtime.FuncForPC(pc)
	e.File = file
	e.Line = line
	e.Func = f.Name()
	return &e
}

func (e err) WithData(data interface{}) *err {
	e.Data = data
	return &e
}

func Parse(s string) *err {
	var e *err
	if err := json.Unmarshal([]byte(s), &e); err != nil {
		e.Message = s
		return e
	}
	return e
}
