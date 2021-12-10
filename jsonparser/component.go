package jsonparser

import (
	"io"
	"net"
	"net/url"
	"strconv"
	"strings"
	"text/template"
	"time"
	"unicode"
	"unicode/utf8"
)

func getName(u string) string {
	p, err := url.Parse(u)
	if err != nil {
		return "Data"
	}
	s := strings.Split(p.Path, "/")
	if len(s) < 1 {
		return "Data"
	}
	return strings.Title(s[len(s)-1])
}

func printTpl(w io.Writer, tplData string, name string, isArray bool) {
	tmpl, err := template.New("test").Parse(tplData)
	if err != nil {
		panic(err)
	}

	data := struct {
		Name    string
		IsArray bool
	}{
		name,
		isArray,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func parseType(value string) (string, bool) {
	if _, err := time.Parse(time.RFC3339, value); err == nil {
		return "time.Time", false
	} else if ip := net.ParseIP(value); ip != nil {
		return "net.IP", false
	} else if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		return "int64", true
	} else if _, err := strconv.ParseFloat(value, 64); err == nil {
		return "float64", true
	} else if _, err := strconv.ParseBool(value); err == nil {
		return "bool", true
	} else {
		return "string", false
	}
}

func filter(name string) string {
	if name == "" {
		name = "Data"
	}
	newString := ""
	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			newString += string(r)
		} else {
			newString += " "
		}
	}
	newString = strings.Title(newString)
	newString = strings.Replace(newString, " ", "", -1)
	newString = strings.Replace(newString, "Url", "URL", -1)
	newString = strings.Replace(newString, "Uri", "URI", -1)
	newString = strings.Replace(newString, "Id", "ID", -1)

	r, _ := utf8.DecodeRuneInString(name)
	if !unicode.IsLetter(r) && !(r == '_') {
		newString = "_" + newString
	}

	return newString
}
