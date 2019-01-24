package mnemonic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	const numTests int = 5
	for i := 0; i < numTests; i++ {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			mn, err := Generate(nil)
			if err != nil {
				t.Error(err)
			}

			if ok := Validate(mn); !ok {
				t.Errorf("%s is invalid mnemonic", mn)
			}
		})
	}
}

func TestToSecret(t *testing.T) {
	for i, tt := range []struct {
		mn       string
		password string
		out      []byte
	}{
		{
			"basket actual",
			"",
			[]byte{92, 242, 212, 168, 176, 53, 94, 144, 41, 91, 223, 197, 101, 160, 34, 164, 9, 175, 6, 61, 83, 101, 187, 87, 191, 116, 217, 82, 143, 73, 75, 250, 68, 0, 245, 61, 131, 73, 184, 15, 218, 228, 64, 130, 215, 249, 84, 30, 29, 186, 43, 0, 59, 207, 236, 157, 13, 83, 120, 28, 166, 118, 101, 31},
		},
		{
			"foo",
			"bar",
			[]byte{52, 240, 234, 234, 147, 188, 0, 49, 5, 117, 253, 10, 137, 26, 15, 134, 167, 162, 183, 62, 215, 236, 31, 58, 80, 160, 40, 132, 157, 105, 148, 243, 181, 44, 146, 137, 78, 171, 209, 175, 75, 14, 113, 127, 28, 134, 178, 10, 49, 234, 133, 153, 176, 108, 195, 186, 22, 95, 206, 105, 49, 82, 122, 169},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			secret, err := ToSecret(tt.mn, tt.password)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(secret, tt.out) {
				t.Errorf("expected %v, received %v", tt.out, secret)
			}
		})
	}
}

func TestToSeed(t *testing.T) {
	for i, tt := range []struct {
		mn       string
		password string
		out      []byte
	}{
		{
			"basket actual",
			"",
			[]byte{92, 242, 212, 168, 176, 53, 94, 144, 41, 91, 223, 197, 101, 160, 34, 164, 9, 175, 6, 61, 83, 101, 187, 87, 191, 116, 217, 82, 143, 73, 75, 250, 68, 0, 245, 61, 131, 73, 184, 15, 218, 228, 64, 130, 215, 249, 84, 30, 29, 186, 43, 0, 59, 207, 236, 157, 13, 83, 120, 28, 166, 118, 101, 31},
		},
		{
			"foo",
			"bar",
			[]byte{52, 240, 234, 234, 147, 188, 0, 49, 5, 117, 253, 10, 137, 26, 15, 134, 167, 162, 183, 62, 215, 236, 31, 58, 80, 160, 40, 132, 157, 105, 148, 243, 181, 44, 146, 137, 78, 171, 209, 175, 75, 14, 113, 127, 28, 134, 178, 10, 49, 234, 133, 153, 176, 108, 195, 186, 22, 95, 206, 105, 49, 82, 122, 169},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			seed, err := ToSeed(tt.mn, tt.password)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(seed, tt.out[:32]) {
				t.Errorf("expected %v, received %v", tt.out[:32], seed)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	for i, tt := range []struct {
		in  string
		out bool
	}{
		{
			"seed sock milk update focus rotate barely fade car face mechanic mercy",
			true,
		},
		// TODO: this fails in the js lib, but not here, why?
		// https://github.com/polkadot-js/common/blob/0f53f0ebf0e77fad949d1c7db146f3bef6b6ec2a/packages/util-crypto/src/mnemonic/validate.spec.js#L16
		//{
		//"wine photo extra cushion basket dwarf humor cloud truck job boat submit",
		//false,
		//},
		{
			"",
			false,
		},
		{
			"foo",
			false,
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if ok := Validate(tt.in); ok != tt.out {
				t.Errorf("expected %v, received %v", tt.out, ok)
			}
		})
	}
}
