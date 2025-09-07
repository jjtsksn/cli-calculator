package splitter

import (
	"strings"
	"unicode"
)

type StringSplitter interface {
	Split(string) []string
}

type SmartStringSplitter struct{}

func NewBasicStringSplitter() *SmartStringSplitter {
	return &SmartStringSplitter{}
}

func (sss SmartStringSplitter) Split(str string) []string {
	str = strings.ReplaceAll(str, " ", "")
	var result []string
	var current strings.Builder

	for _, char := range str {
		if unicode.IsDigit(char) || char == '.' || char == ',' {
			if char == ',' {
				char = '.'
			}
			current.WriteRune(char)
		} else {
			if current.Len() > 0 {
				result = append(result, current.String())
				current.Reset()
			}
			result = append(result, string(char))
		}
	}
	if current.Len() > 0 {
		result = append(result, current.String())
	}
	return result
}
