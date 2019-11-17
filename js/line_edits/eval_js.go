// +build js,!wasm

package main

import "github.com/gopherjs/gopherjs/js"

func evalJS(s string) {
	js.Global.Call("eval", s)
}
