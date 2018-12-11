package ext

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewExtError(t *testing.T) {
	type input struct {
		message string
		code    int
	}

	for i, tt := range []struct {
		in  input
		out *Error
	}{
		{
			input{"test message", 100},
			&Error{
				Message: "test message",
				Code:    100,
			},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := NewExtError(tt.in.message, tt.in.code)
			if !reflect.DeepEqual(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}
