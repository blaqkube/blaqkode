package wasm

import (
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	wasmtime "github.com/bytecodealliance/wasmtime-go"
	"github.com/stretchr/testify/assert"
	wasmer "github.com/wasmerio/go-ext-wasm/wasmer"
)

const MaxLoop int = 1000000

func TestRunWasmer(t *testing.T) {
	bytes, _ := wasmer.ReadBytes("simple/target/wasm32-unknown-unknown/release/simple.wasm")

	instance, _ := wasmer.NewInstance(bytes)
	defer instance.Close()

	sum := instance.Exports["sum"]

	output := int64(0)
	for i := 0; i < MaxLoop; i++ {
		result, _ := sum(5, 37)
		if i%100000 == 0 {
			fmt.Printf(".")
		}
		output = result.ToI64()
	}
	fmt.Printf("\n")
	assert.Equal(t, int64(42), output, "script should return 42")
}

func TestRunWasmtime(t *testing.T) {
	wasm, _ := ioutil.ReadFile("simple/target/wasm32-unknown-unknown/release/simple.wasm")
	store := wasmtime.NewStore(wasmtime.NewEngine())
	module, err := wasmtime.NewModule(store.Engine, wasm)
	assert.Equal(t, nil, err, "should succeed")

	instance, err := wasmtime.NewInstance(store, module, []*wasmtime.Extern{})
	assert.Equal(t, nil, err, "should succeed")

	sum := instance.GetExport("sum").Func()
	output := int32(0)
	for i := 0; i < MaxLoop; i++ {
		result, _ := sum.Call(5, 37)
		if i%100000 == 0 {
			fmt.Printf(".")
		}
		switch i := result.(type) {
		case int32:
			output = i
		default:
			panic(errors.New("wrong type"))
		}
	}
	assert.Equal(t, int32(42), output, "script should return int32(42)")
}
