package main

import (
	"syscall/js"
)

func add(i []js.Value) {
	js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
	println(js.ValueOf(i[0].Int() + i[1].Int()).String())
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallBack(add))
}

func main() {
	c := make(chan struct{}, 0)

	println("Go WebAssembly")
	registerCallbacks()

	<-c
}
