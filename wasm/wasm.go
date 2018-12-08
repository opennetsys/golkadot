package wasm

import (
	"errors"

	"github.com/c3systems/go-substrate/log"
	"github.com/perlin-network/life/compiler"
	"github.com/perlin-network/life/exec"
)

// Config ...
type Config struct {
	Input []byte
}

// VM ...
type VM struct {
	client *exec.VirtualMachine
}

var (
	// ErrInvalidEntryName ...
	ErrInvalidEntryName = errors.New("Invalid entry name")
	// ErrEntryFunctionNotFound ...
	ErrEntryFunctionNotFound = errors.New("Entry function not found")
	// ErrCasting ...
	ErrCasting = errors.New("Error casting value")
)

// NewVM ...
func NewVM(config *Config) *VM {
	gasPolicy := &compiler.SimpleGasPolicy{
		GasPerInstruction: int64(1),
	}

	vm, err := exec.NewVirtualMachine(config.Input, exec.VMConfig{}, &exec.NopResolver{}, gasPolicy)
	if err != nil {
		log.Fatalf("[wasm] bytecode is invalid; %s", err)
		panic(err)
	}

	return &VM{
		client: vm,
	}
}

// Execute ...
func (vm *VM) Execute(input ...interface{}) (interface{}, error) {
	fnName, ok := input[0].(string)
	if !ok {
		return nil, ErrInvalidEntryName
	}
	entryID, ok := vm.client.GetFunctionExport(fnName)
	if !ok {
		return nil, ErrEntryFunctionNotFound
	}

	var args []int64
	for _, in := range input[1:] {
		v, ok := in.(int64)
		if !ok {
			return nil, ErrCasting
		}

		args = append(args, int64(v))
	}

	ret, err := vm.client.Run(entryID, args...)
	if err != nil {
		log.Fatalf("[wasm] execution error; %s", err)
		vm.client.PrintStackTrace()
		return nil, err
	}

	return ret, nil
}
