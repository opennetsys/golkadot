package address

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	for i, tt := range []struct {
		in     []byte
		prefix PrefixEnum
		toPass bool
		out    string
	}{
		{
			[]byte{209, 114, 167, 76, 218, 76, 134, 89, 18, 195, 43, 160, 168, 10, 87, 174, 105, 171, 174, 65, 14, 92, 203, 89, 222, 232, 78, 47, 68, 50, 219, 79},
			nil,
			true,
			"5GoKvZWG5ZPYL1WUovuHW3zJBWBP5eT8CbqjdRY4Q6iMaDtZ",
		}, {
			[]byte{215, 86, 142, 95, 10, 126, 218, 103, 168, 38, 145, 255, 55, 154, 196, 187, 164, 249, 201, 184, 89, 254, 119, 155, 93, 70, 54, 59, 97, 173, 45, 185},
			nil,
			true,
			"5Gw3s7q4QLkSWwknsiPtjujPv3XM4Trxi5d4PgKMMk3gfGTE",
		}, {
			[]byte{163, 155, 255, 175, 56, 21, 16, 70, 7, 55, 3, 233, 117, 111, 140, 228, 158, 235, 160, 83, 65, 81, 247, 227, 251, 198, 170, 24, 95, 6, 179, 164},
			nil,
			true,
			"5FmE1Adpwp1bT1oY95w59RiSPVu9QwzBGjKsE2hxemD2AFs8",
		}, {
			[]byte{191, 200, 35, 170, 117, 195, 0, 88, 238, 236, 33, 171, 226, 194, 214, 183, 36, 116, 24, 164, 175, 137, 214, 122, 32, 132, 194, 172, 134, 77, 160, 128},
			nil,
			true,
			"5GQATTPFqkfze7kbGuvezhV921FwSp6Xyr8FMW1pmi6LjDuu",
		}, {
			[]byte{13, 113, 209, 169, 202, 214, 242, 171, 119, 52, 53, 167, 222, 193, 186, 192, 25, 153, 77, 5, 209, 221, 94, 179, 16, 130, 17, 220, 242, 92, 157, 30},
			nil,
			true,
			"5CNLHq4doqBbrrxLCxAakEgaEvef5tjSrN7QqJwcWzNd7W7k",
		}, {
			[]byte{210, 222, 115, 148, 174, 4, 122, 85, 2, 173, 154, 219, 156, 198, 159, 246, 254, 72, 64, 51, 191, 206, 135, 77, 119, 93, 169, 71, 72, 124, 216, 50},
			nil,
			true,
			"5GqBzeuVYJBorP3oP7FgheoP5nb2twFeDUFZhoBhX7ExYsia",
		},
		{
			[]byte{209, 114, 167, 76, 218, 76, 134, 89, 18, 195, 43, 160, 168, 10, 87, 174, 105, 171, 174, 65, 14, 92, 203, 89, 222, 232, 78, 47, 68, 50, 219, 79}[0:30],
			nil,
			false,
			"",
		},
		{
			[]byte{1},
			nil,
			true,
			"F7L6",
		},
		{
			[]byte{1},
			SixtyEight,
			true,
			"Pqt7",
		},
		{
			[]byte{0, 1},
			SixtyEight,
			true,
			"2jpAJz",
		},
		{
			[]byte{1, 2, 3, 4},
			SixtyEight,
			true,
			"as7QnGQ7",
		},
		{
			[]byte{42, 44, 10, 0, 0, 0, 0, 0},
			SixtyEight,
			true,
			"4q7qY5RBG7Z4xX",
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := Encode(tt.in, tt.prefix)
			if !tt.toPass && err == nil {
				t.Error("expected err != nil")
			}
			if tt.toPass && err != nil {
				t.Error(err)
			}
			if tt.toPass && result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	for i, tt := range []struct {
		in     string
		prefix PrefixEnum
		toPass bool
		out    []byte
	}{
		{
			"5GoKvZWG5ZPYL1WUovuHW3zJBWBP5eT8CbqjdRY4Q6iMaDtZ",
			nil,
			true,
			[]byte{209, 114, 167, 76, 218, 76, 134, 89, 18, 195, 43, 160, 168, 10, 87, 174, 105, 171, 174, 65, 14, 92, 203, 89, 222, 232, 78, 47, 68, 50, 219, 79},
		}, {
			"5Gw3s7q4QLkSWwknsiPtjujPv3XM4Trxi5d4PgKMMk3gfGTE",
			nil,
			true,
			[]byte{215, 86, 142, 95, 10, 126, 218, 103, 168, 38, 145, 255, 55, 154, 196, 187, 164, 249, 201, 184, 89, 254, 119, 155, 93, 70, 54, 59, 97, 173, 45, 185},
		}, {
			"5FmE1Adpwp1bT1oY95w59RiSPVu9QwzBGjKsE2hxemD2AFs8",
			nil,
			true,
			[]byte{163, 155, 255, 175, 56, 21, 16, 70, 7, 55, 3, 233, 117, 111, 140, 228, 158, 235, 160, 83, 65, 81, 247, 227, 251, 198, 170, 24, 95, 6, 179, 164},
		}, {
			"5GQATTPFqkfze7kbGuvezhV921FwSp6Xyr8FMW1pmi6LjDuu",
			nil,
			true,
			[]byte{191, 200, 35, 170, 117, 195, 0, 88, 238, 236, 33, 171, 226, 194, 214, 183, 36, 116, 24, 164, 175, 137, 214, 122, 32, 132, 194, 172, 134, 77, 160, 128},
		}, {
			"5CNLHq4doqBbrrxLCxAakEgaEvef5tjSrN7QqJwcWzNd7W7k",
			nil,
			true,
			[]byte{13, 113, 209, 169, 202, 214, 242, 171, 119, 52, 53, 167, 222, 193, 186, 192, 25, 153, 77, 5, 209, 221, 94, 179, 16, 130, 17, 220, 242, 92, 157, 30},
		}, {
			"5GqBzeuVYJBorP3oP7FgheoP5nb2twFeDUFZhoBhX7ExYsia",
			nil,
			true,
			[]byte{210, 222, 115, 148, 174, 4, 122, 85, 2, 173, 154, 219, 156, 198, 159, 246, 254, 72, 64, 51, 191, 206, 135, 77, 119, 93, 169, 71, 72, 124, 216, 50},
		},
		{
			"",
			nil,
			false,
			[]byte{209, 114, 167, 76, 218, 76, 134, 89, 18, 195, 43, 160, 168, 10, 87, 174, 105, 171, 174, 65, 14, 92, 203, 89, 222, 232, 78, 47, 68, 50, 219, 79}[0:30],
		},
		{
			"F7L6",
			nil,
			true,
			[]byte{1},
		},
		{
			"Pqt7",
			SixtyEight,
			true,
			[]byte{1},
		},
		{
			"2jpAJz",
			SixtyEight,
			true,
			[]byte{0, 1},
		},
		{
			"as7QnGQ7",
			SixtyEight,
			true,
			[]byte{1, 2, 3, 4},
		},
		{
			"4q7qY5RBG7Z4xX",
			SixtyEight,
			true,
			[]byte{42, 44, 10, 0, 0, 0, 0, 0},
		},
		{
			"0x01020304",
			nil,
			true,
			[]byte{1, 2, 3, 4},
		},
		// TODO...
		//{
		//// invalid prefix
		//"6GfvWUvHvU8otbZ7sFhXH4eYeMcKdUkL61P3nFy52efEPVUx",
		//nil,
		//false,
		//nil,
		//},
		{
			// invalid length
			"y9EMHt34JJo4rWLSaxoLGdYXvjgSXEd4zHUnQgfNzwES8b",
			nil,
			false,
			nil,
		},
		{
			// invalid checksum
			"5GoKvZWG5ZPYL1WUovuHW3zJBWBP5eT8CbqjdRY4Q6iMa9cj",
			nil,
			false,
			nil,
		},
		{
			// invalid checksum
			"5GoKvZWG5ZPYL1WUovuHW3zJBWBP5eT8CbqjdRY4Q6iMaDwU",
			nil,
			false,
			nil,
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := Decode(tt.in, tt.prefix)
			if !tt.toPass && err == nil {
				t.Error("expected err != nil")
			}
			if tt.toPass && err != nil {
				t.Error(err)
			}
			if tt.toPass && !reflect.DeepEqual(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}

}
