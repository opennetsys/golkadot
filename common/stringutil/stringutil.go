package stringutil

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

var numberSequence = regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
var numberReplacement = []byte(`$1 $2 $3`)

func addWordBoundariesToNumbers(s string) string {
	b := []byte(s)
	b = numberSequence.ReplaceAll(b, numberReplacement)
	return string(b)
}

// ToCamelCase converts a string to camel casing.
func ToCamelCase(s string) string {
	if s == "" {
		return s
	}
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}

	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	capNext := false
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' {
			capNext = true
		} else {
			capNext = false
		}
	}

	return n
}

// Shorten returns a string with maximum length.  Checks the string against the `prefixLength`, if longer than dopuble this, shortens it by placing `..` in the middle of it.
func Shorten(value string, prefixLength int) string {
	if len(value) <= 2+2*prefixLength {
		return value
	}

	start := value[0:prefixLength]
	end := value[len(value)-prefixLength : len(value)]
	return fmt.Sprintf("%s..%s", start, end)
}

// LowerFirst lowercase the first letter of a string
func LowerFirst(value string) string {
	for i, v := range value {
		return string(unicode.ToLower(v)) + value[i+1:]
	}
	return ""
}

// UpperFirst uppercase the first letter of a string
func UpperFirst(value string) string {
	for i, v := range value {
		return string(unicode.ToUpper(v)) + value[i+1:]
	}
	return ""
}

// ToUint8Slice creates an uint8 array from a utf-8 string
func ToUint8Slice(value string) []uint8 {
	return []uint8(value)
}
