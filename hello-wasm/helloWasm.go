package main

import (
	"fmt"
	"github.com/bytecodealliance/wasmtime-go"
	"github.com/wasmerio/wasmer-go/wasmer"
	"log"
	"os"
)

func main() {
	engine := wasmtime.NewEngine()
	store := wasmtime.NewStore(engine)
	file, err := wasmtime.NewModuleFromFile(engine, "gcd.wat")
	if err != nil {
		log.Fatal(err)
	}

	instance, err := wasmtime.NewInstance(store, file, []wasmtime.AsExtern{})
	if err != nil {
		log.Fatal(err)
	}

	gcd := instance.GetExport(store, "gcd").Func()
	val, err := gcd.Call(store, 6, 27)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result: ", val.(int32))

	useWasmer()
}

func useWasmer() {
	wasmBytes, err := os.ReadFile("gcd.wat")
	if err != nil {
		log.Fatal("Read file error:", err)
	}

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)
	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		log.Fatal(err)
	}

	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		log.Fatal(err)
	}

	sum, err := instance.Exports.GetFunction("gcd")
	if err != nil {
		log.Fatal(err)
	}

	result, err := sum(6, 27)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result of sum:", result)
}
