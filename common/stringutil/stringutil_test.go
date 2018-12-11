package stringutil

import "testing"

func TestReverse(t *testing.T) {
	for _, tt := range []struct {
		in, out string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	} {
		result := Reverse(tt.in)
		if result != tt.out {
			t.Errorf("want %v; got %v", tt.out, result)
		}
	}
}
