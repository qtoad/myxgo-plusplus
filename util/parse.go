package util

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	textTmpl "text/template"
	"time"
)

func ParseBool(s string) bool {
	b1, _ := strconv.ParseBool(s)
	return b1
}

func ParseFloat32(f string) float32 {
	s, _ := strconv.ParseFloat(f, 32)
	return float32(s)
}

func ParseTime(stringTime string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
	return theTime
}

func ParseHTML(name, tmpl string, param interface{}) template.HTML {
	t := template.New(name)
	t, err := t.Parse(tmpl)
	if err != nil {
		fmt.Println("util parseHTML error", err)
		return ""
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, param)
	if err != nil {
		fmt.Println("util parseHTML error", err)
		return ""
	}
	return template.HTML(buf.String())
}

func ParseText(name, tmpl string, param interface{}) string {
	t := textTmpl.New(name)
	t, err := t.Parse(tmpl)
	if err != nil {
		fmt.Println("util parseHTML error", err)
		return ""
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, param)
	if err != nil {
		fmt.Println("util parseHTML error", err)
		return ""
	}
	return buf.String()
}
