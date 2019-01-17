package codec

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWriteBinary(t *testing.T) {
	for i, tt := range []struct {
		in  interface{}
		out []byte
	}{
		{
			[]byte("bazzing"),
			[]byte{98, 97, 122, 122, 105, 110, 103},
		},
		{
			int32(69),
			[]byte{69, 0, 0, 0},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			bin, err := writeBinary(tt.in)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(bin, tt.out) {
				t.Errorf("want %v; got %v", tt.out, bin)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	for i, tt := range []struct {
		in  interface{}
		out []byte
	}{
		{
			struct {
				Foo string
				Bar int
			}{
				Foo: "bazzing",
				Bar: 69,
			},
			[]byte{28, 98, 97, 122, 122, 105, 110, 103, 69, 0, 0, 0},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			enc, err := Encode(tt.in)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(enc, tt.out) {
				t.Errorf("want %v; got %v", tt.out, enc)
			}
		})
	}
}
