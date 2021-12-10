package jsonparser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	goDefaultTpl = `type {{.Name}} {{if .IsArray}}[]{{end}}struct {
`
	goArrayTpl = `_ {{if .IsArray}}[]{{end}}struct {
`
	goATplEmpty = `_ {{if .IsArray}}[]{{end}}struct {
}
`
)

//Model ...
type Model struct {
	Writer      io.Writer
	Name        string
	Data        interface{}
	WithExample bool
	Format      bool
	Convert     bool
}

//New ...
func New(byte []byte, name string) (m *Model, err error) {
	var data interface{}
	if err = json.Unmarshal(byte, &data); err != nil {
		return
	}
	return &Model{
		Writer: os.Stdout,
		Data:   data,
		Name:   filter(name),
		Format: true,
	}, nil
}

//Get ...
func Get(url string) ([]byte, string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Add("Accept", "application/json")
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer func() {
		_ = r.Body.Close()
	}()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, "", err
	}
	return b, getName(url), err
}

//WriteGo ...
func (m *Model) WriteGo() (b []byte, err error) {
	if m.Format {
		var buf bytes.Buffer
		m.Writer = &buf
		m.print(func(ms map[string]interface{}) {
			m.parseMap(ms)
		})
		b, err = format.Source(buf.Bytes())
		org := m.Writer
		if err == nil {
			//_, _ = org.Write(b)
		} else {
			_, _ = io.Copy(org, &buf)
		}
		m.Writer = org
	} else {
		m.print(func(ms map[string]interface{}) {
			m.parseMap(ms)
		})
	}
	return
}

func (m *Model) print(convert func(map[string]interface{})) {
	structNum := 0
	defer func() {
		for i := 0; i < structNum; i++ {
			_, _ = fmt.Fprintln(m.Writer, "}")
		}
		_, _ = fmt.Fprintln(m.Writer, "}")
	}()

	switch v := m.Data.(type) {
	case []interface{}:
		printTpl(m.Writer, goDefaultTpl, m.Name, true)
	init:
		for _, node := range v {
			switch node.(type) {
			case []interface{}:
				if len(node.([]interface{})) == 0 {
					printTpl(m.Writer, goATplEmpty, m.Name, true)
				} else {
					v = node.([]interface{})
					printTpl(m.Writer, goArrayTpl, m.Name, true)
					structNum++
					goto init
				}
			default:
				if len(node.(map[string]interface{})) == 0 {
					printTpl(m.Writer, goATplEmpty, m.Name, false)
				}
				convert(node.(map[string]interface{}))
			}
		}
	case float64:
		break
	default:
		printTpl(m.Writer, goDefaultTpl, m.Name, false)
		convert(m.Data.(map[string]interface{}))
	}
}

func (m *Model) parseMap(ms map[string]interface{}) {
	keys := getSortedKeys(ms)
	for _, k := range keys {
		m.parse(ms[k], k)
	}
}

func (m *Model) parse(data interface{}, k string) {
	switch v := data.(type) {
	case string:
		if m.Convert {
			t, converted := parseType(v)
			m.printType(k, v, t, converted)
		} else {
			m.printType(k, v, "string", false)
		}
	case bool:
		m.printType(k, v, "bool", false)
	case float64:
		//json parser always returns a float for number values, check if it is an int value
		if float64(int64(v)) == v {
			m.printType(k, v, "int64", false)
		} else {
			m.printType(k, v, "float64", false)
		}
	case int64:
		m.printType(k, v, "int64", false)
	case []interface{}:
		if len(v) > 0 {
			switch vv := v[0].(type) {
			case string:
				m.printType(k, v[0], "[]string", false)
			case float64:
				//json parser always returns a float for number values, check if it is an int value
				if float64(int64(v[0].(float64))) == v[0].(float64) {
					m.printType(k, v[0], "[]int64", false)
				} else {
					m.printType(k, v[0], "[]float64", false)
				}
			case bool:
				m.printType(k, v[0], "[]bool", false)
			case []interface{}:
				m.parse(vv[0], k)
				//m.printObject(k, "[]struct", func() { m.parse(vv[0], k) })
			case map[string]interface{}:
				m.printObject(k, "[]struct", func() { m.parseMap(vv) })
			default:
				//fmt.Printf("unknown type: %T", vv)
				m.printType(k, nil, "interface{}", false)
			}
		} else {
			m.printType(k, nil, "[]interface{}", false)
		}
	case map[string]interface{}:
		m.printObject(k, "struct", func() { m.parseMap(v) })
	default:
		m.printType(k, nil, "interface{}", false)
	}
}

func (m *Model) printType(key string, value interface{}, t string, converted bool) {
	name := filter(key)
	if converted {
		key += ",string"
	}
	if m.WithExample {
		_, _ = fmt.Fprintf(m.Writer, "%s %s `json:\"%s\"` // %v\n", name, t, key, value)
	} else {
		_, _ = fmt.Fprintf(m.Writer, "%s %s `json:\"%s\"`\n", name, t, key)
	}
}

func (m *Model) printObject(n string, t string, f func()) {
	_, _ = fmt.Fprintf(m.Writer, "%s %s {\n", filter(n), t)
	f()
	_, _ = fmt.Fprintf(m.Writer, "} `json:\"%s\"`\n", n)
}
