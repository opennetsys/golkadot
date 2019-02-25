package wasm

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestExecute(t *testing.T) {
	t.Skip()
	input, err := ioutil.ReadFile("wasm_test_fib.wasm")
	if err != nil {
		t.Fatalf("[wasm] error reading file; %s", err)
	}

	vm := NewVM(&Config{
		Input: input,
	})

	var tests = []struct {
		in  int64
		out int64
	}{
		{4, 3},
		{15, 610},
		{20, 6765},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			ret, err := vm.Execute("fib", tt.in)
			if err != nil {
				t.Fatal(err)
			}

			v, ok := ret.(int64)
			if !ok {
				t.Fatal("error casting to int64")
			}

			if v != tt.out {
				t.Errorf("got %d, want %d", v, tt.out)
			}
		})
	}
}
