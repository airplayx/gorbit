package gorbit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
)

type e struct {
	Message string      `json:"message,omitempty"`
	File    string      `json:"file,omitempty"`
	Line    int         `json:"line,omitempty"`
	Func    string      `json:"func,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (e *e) Error() string {
	var buf bytes.Buffer
	if e.File != "" {
		buf.WriteString(fmt.Sprintf("[%s - %s : %d] ", e.File, e.Func, e.Line))
	}
	buf.WriteString(e.Message)
	return buf.String()
}

func (e e) Format(args ...interface{}) *e {
	e.Message = fmt.Sprintf(e.Message, args...)
	return &e
}

func (e e) Location() *e {
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

func (e e) WithData(data interface{}) *e {
	e.Data = data
	return &e
}

func Parse(s string) *e {
	var e *e
	if err := json.Unmarshal([]byte(s), &e); err != nil {
		e.Message = s
		return e
	}
	return e
}
