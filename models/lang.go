package models

import (
	"bytes"
	"encoding/json"
)

// Language represents enum of programming languages
type Language int

// programming languages
const (
	LangCpp14GCC Language = 3003 // C++14 (GCC 5.4.1)
	LangNone     Language = 0
)

var langToStr = map[Language]string{
	LangCpp14GCC: "C++14 (GCC 5.4.1)",
	LangNone:     "<none>",
}

var strToLang = map[string]Language{
	"C++14 (GCC 5.4.1)": LangCpp14GCC,
	"<none>":            LangNone,
}

func (l Language) String() string {
	str, ok := langToStr[l]
	if !ok {
		str = langToStr[LangNone]
	}
	return str
}

func (l Language) Int() int {
	return int(l)
}

func NewLanguage(langStr string) Language {
	l, ok := strToLang[langStr]
	if !ok {
		l = LangNone
	}
	return l
}

func (l Language) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(l.String())
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (l Language) MarshalYAML() (interface{}, error) {
	return l.String(), nil
}

func (l *Language) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*l = NewLanguage(s)
	return nil
}

func (l *Language) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	*l = NewLanguage(s)
	return nil
}
