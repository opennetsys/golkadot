package types

import (
	"reflect"
	"testing"
)

func TestOption(t *testing.T) {
	t.Run("can decode a type", func(t *testing.T) {
		t.Run("string (with)", func(t *testing.T) {
			o := NewOption(NewText("foo"))

			expected := "foo"
			if o.String() != expected {
				t.Fail()
			}
			if o.IsNone() == (len(expected) > 0) {
				t.Fail()
			}
		})

		t.Run("uint8 array (with)", func(t *testing.T) {
			o := NewOption(NewText([]uint8{1, 12, 102, 111, 111}))

			expected := "foo"
			if o.String() != expected {
				t.Fail()
			}
		})

		t.Run("uint8 array (without)", func(t *testing.T) {
			o := NewOption(NewText([]uint8{0}))

			expected := ""
			if o.String() != expected {
				t.Fail()
			}
		})
	})

	t.Run("can encode a type", func(t *testing.T) {
		t.Run("can encode to hex", func(t *testing.T) {
			o := NewOption(NewText("foo"))

			expected := "0x010c666f6f"
			if o.Hex() != expected {
				t.Fail()
			}
		})

		t.Run("can encode to string", func(t *testing.T) {
			o := NewOption(NewText("foo"))

			expected := "foo"
			if o.String() != expected {
				t.Fail()
			}
		})

		t.Run("can encode to uint8 slice", func(t *testing.T) {
			o := NewOption(NewText("foo"))

			expected := []byte{1, 12, 102, 111, 111}
			if !reflect.DeepEqual(o.ToU8a(false), expected) {
				t.Fail()
			}
		})
	})

	t.Run("has empty String() (empty)", func(t *testing.T) {
		o := NewOption(NewText(""))

		expected := ""
		if o.String() != expected {
			t.Fail()
		}
	})

	t.Run("has value String() (provided)", func(t *testing.T) {
		o := NewOption(NewText([]uint8{1, 4 << 2, 49, 50, 51, 52}))

		expected := "1234"
		if o.String() != expected {
			t.Fail()
		}
	})

	t.Run("converts ToU8a() with", func(t *testing.T) {
		o := NewOption(NewText("1234"))

		expected := []uint8{1, 4 << 2, 49, 50, 51, 52}
		if !reflect.DeepEqual(o.ToU8a(false), expected) {
			t.Fail()
		}
	})

	t.Run("converts ToU8a() without", func(t *testing.T) {
		o := NewOption(NewText(nil))

		expected := []uint8{0}
		if !reflect.DeepEqual(o.ToU8a(false), expected) {
			t.Fail()
		}
	})

	t.Run("compare against other option", func(t *testing.T) {
		o := NewOption(NewText("1234"))

		if !o.Equals(NewOption(NewText("1234"))) {
			t.Fail()
		}
	})

	t.Run("compare against raw value", func(t *testing.T) {
		o := NewOption(NewText("1234"))

		if !o.Equals("1234") {
			t.Fail()
		}
	})
}
