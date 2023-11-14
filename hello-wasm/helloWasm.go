package main

import (
	"fmt"
	"github.com/bytecodealliance/wasmtime-go"
	"log"
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
}
