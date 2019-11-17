// +build js,!wasm

package main

import (
	"unsafe"

	"github.com/gopherjs/gopherjs/js"
)

func evalJSToPtr(s string) unsafe.Pointer {
	return unsafe.Pointer(js.Global.Call("eval", s).Unsafe())
}

func createGlobalFunction(fn string, f interface{}) {
	js.Global.Set(fn, f)
}
