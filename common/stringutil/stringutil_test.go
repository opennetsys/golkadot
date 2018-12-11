package stringutil

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	for i, tt := range []struct {
		in, out string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := Reverse(tt.in)
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestToCamecalCase(t *testing.T) {
	for i, tt := range []struct {
		in, out string
	}{
		{"snake_case", "snakeCase"},
		{"SnakeCase", "snakeCase"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := ToCamelCase(tt.in)
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestShorten(t *testing.T) {
	type input struct {
		value        string
		prefixLength int
	}
	for i, tt := range []struct {
		in  input
		out string
	}{
		{input{"0123456789", 4}, "0123456789"},
		{input{"0123456789", 3}, "012..789"},
		{input{"12345678", 2}, "12..78"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := Shorten(tt.in.value, tt.in.prefixLength)
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestLowerFirst(t *testing.T) {
	for i, tt := range []struct {
		in  string
		out string
	}{
		{"ABC", "aBC"},
		{"abc", "abc"},
		{"", ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := LowerFirst(tt.in)
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestUpperFirst(t *testing.T) {
	for i, tt := range []struct {
		in  string
		out string
	}{
		{"abc", "Abc"},
		{"ABC", "ABC"},
		{"", ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := UpperFirst(tt.in)
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestToUint8Slice(t *testing.T) {
	for i, tt := range []struct {
		in  string
		out []uint8
	}{
		{"Привет, мир!", []uint8{208, 159, 209, 128, 208, 184, 208, 178, 208, 181, 209, 130, 44, 32, 208, 188, 208, 184, 209, 128, 33}},
		{"", []uint8{}},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := ToUint8Slice(tt.in)
			if !reflect.DeepEqual(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}
