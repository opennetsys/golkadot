package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestText(t *testing.T) {
	t.Run("decode", func(t *testing.T) {
		for i, tt := range []struct {
			in  interface{}
			out string
		}{
			{"foo", "foo"},
			{[]uint8{12, 102, 111, 111}, "foo"},
			{NewTextFromString("foo"), "foo"},
		} {
			t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
				result := NewText(tt.in)
				if result.String() != tt.out {
					t.Errorf("want %v; got %v", tt.out, result)
				}
			})
		}
	})

	t.Run("encode", func(t *testing.T) {
		text := NewTextFromString("foo")
		t.Run("to string", func(t *testing.T) {
			got := text.String()
			want := "foo"
			if got != want {
				t.Errorf("want %v; got %v", want, got)
				t.Fail()
			}
		})
		t.Run("to hex", func(t *testing.T) {
			got := text.Hex()
			want := "0x0c666f6f"
			if got != want {
				t.Errorf("want %v; got %v", want, got)
				t.Fail()
			}
		})
		t.Run("to bytes", func(t *testing.T) {
			got := text.Bytes()
			want := []uint8{12, 102, 111, 111}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("want %v; got %v", want, got)
				t.Fail()
			}
		})
	})

	t.Run("length", func(t *testing.T) {
		text := NewTextFromString("foo")
		got := text.Len()
		want := 3
		if got != want {
			t.Errorf("want %v; got %v", want, got)
			t.Fail()
		}
	})

	t.Run("encoded length", func(t *testing.T) {
		text := NewTextFromString("foo")
		got := text.EncodedLen()
		want := 4
		if got != want {
			t.Errorf("want %v; got %v", want, got)
			t.Fail()
		}
	})
}
