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

func TestEncodeDecode(t *testing.T) {
	var (
		outU64 uint64
		outU32 uint32
		outU8  uint8
		outI64 int64
		outI32 int32
		outI8  int8
		enc    []byte
		err    error
	)

	// u64
	for i, tt := range []struct {
		in uint64
	}{
		{
			63,
		}, {
			64,
		}, {
			16383,
		}, {
			16384,
		}, {
			1073741823,
		}, {
			1073741824,
		}, {
			(1 << 32) - 1,
		}, {
			1 << 32,
		}, {
			1 << 40,
		}, {
			1 << 48,
		}, {
			(1 << 56) - 1,
		}, {
			1 << 56,
		},
	} {
		t.Run(fmt.Sprintf("uint64 test %v", i), func(t *testing.T) {
			enc, err = Encode(tt.in)
			if err != nil {
				t.Error(err)
				return
			}

			if err = Decode(enc, &outU64); err != nil {
				t.Errorf("err decoding %v\n%v", enc, err)
				return
			}

			if tt.in != outU64 {
				t.Errorf("expected %v, received %v", tt.in, outU64)
			}
		})
	}

	// u32
	for i, tt := range []struct {
		in uint32
	}{
		{
			63,
		}, {
			64,
		}, {
			16383,
		}, {
			16384,
		}, {
			1073741823,
		}, {
			1073741824,
		},
	} {
		t.Run(fmt.Sprintf("uint32 test %v", i), func(t *testing.T) {
			enc, err = Encode(tt.in)
			if err != nil {
				t.Error(err)
				return
			}

			if err = Decode(enc, &outU32); err != nil {
				t.Errorf("err decoding %v\n%v", enc, err)
				return
			}

			if tt.in != outU32 {
				t.Errorf("expected %v, received %v", tt.in, outU32)
			}
		})
	}

	// u8
	for i, tt := range []struct {
		in uint8
	}{
		{
			63,
		}, {
			64,
		}, {
			255,
		},
	} {
		t.Run(fmt.Sprintf("uint8 test %v", i), func(t *testing.T) {
			enc, err = Encode(tt.in)
			if err != nil {
				t.Error(err)
				return
			}

			if err = Decode(enc, &outU8); err != nil {
				t.Errorf("err decoding %v\n%v", enc, err)
				return
			}

			if tt.in != outU8 {
				t.Errorf("expected %v, received %v", tt.in, outU8)
			}
		})
	}

	// i64
	for i, tt := range []struct {
		in int64
	}{
		{
			63,
		}, {
			-63,
		}, {
			64,
		}, {
			-64,
		}, {
			16383,
		}, {
			-16383,
		}, {
			16384,
		}, {
			-16384,
		}, {
			1073741823,
		}, {
			-1073741823,
		}, {
			1073741824,
		}, {
			-1073741824,
		}, {
			(1 << 32) - 1,
		}, {
			-1 * ((1 << 32) - 1),
		}, {
			1 << 32,
		}, {
			-1 * (1 << 32),
		}, {
			1 << 40,
		}, {
			-1 * (1 << 40),
		}, {
			1 << 48,
		}, {
			-1 * (1 << 48),
		}, {
			(1 << 56) - 1,
		}, {
			-1 * ((1 << 56) - 1),
		}, {
			1 << 56,
		}, {
			-1 * (1 << 56),
		},
	} {
		t.Run(fmt.Sprintf("int64 test %v", i), func(t *testing.T) {
			enc, err = Encode(tt.in)
			if err != nil {
				t.Error(err)
				return
			}

			if err = Decode(enc, &outI64); err != nil {
				t.Errorf("err decoding %v\n%v", enc, err)
				return
			}

			if tt.in != outI64 {
				t.Errorf("expected %v, received %v", tt.in, outI64)
			}
		})
	}

	// i32
	for i, tt := range []struct {
		in int32
	}{
		{
			63,
		}, {
			-63,
		}, {
			64,
		}, {
			-64,
		}, {
			16383,
		}, {
			-16383,
		}, {
			16384,
		}, {
			-16384,
		}, {
			1073741823,
		}, {
			-1073741823,
		}, {
			1073741824,
		}, {
			-1073741824,
		},
	} {
		t.Run(fmt.Sprintf("int32 test %v", i), func(t *testing.T) {
			enc, err = Encode(tt.in)
			if err != nil {
				t.Error(err)
				return
			}

			if err = Decode(enc, &outI32); err != nil {
				t.Errorf("err decoding %v\n%v", enc, err)
				return
			}

			if tt.in != outI32 {
				t.Errorf("expected %v, received %v", tt.in, outI32)
			}
		})
	}

	// i8
	for i, tt := range []struct {
		in int8
	}{
		{
			63,
		}, {
			-63,
		}, {
			64,
		}, {
			-64,
		}, {
			127,
		}, {
			-127,
		},
	} {
		t.Run(fmt.Sprintf("int8 test %v", i), func(t *testing.T) {
			enc, err = Encode(tt.in)
			if err != nil {
				t.Error(err)
				return
			}

			if err = Decode(enc, &outI8); err != nil {
				t.Errorf("err decoding %v\n%v", enc, err)
				return
			}

			if tt.in != outI8 {
				t.Errorf("expected %v, received %v", tt.in, outI8)
			}
		})
	}
}
