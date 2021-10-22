package main

import (
	"syscall/js"
)

func main() {
	registerCallbacks()

	select {}
}

func registerCallbacks() {
	js.Global().Set("hello", js.FuncOf(Hello))
}

func Hello(this js.Value, args []js.Value) interface{} {
	return "Username"
}
