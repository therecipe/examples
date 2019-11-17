// +build !js wasm

package main

import (
	"unsafe"

	"github.com/therecipe/qt/qml"
)

var engine *qml.QJSEngine

func evalJSToPtr(s string) unsafe.Pointer {
	if engine == nil {
		engine = qml.NewQJSEngine()
	}
	return unsafe.Pointer(uintptr(engine.Evaluate(s, "", 0).ToVariant().ToULongLong(nil)))
}

func createGlobalFunction(fn string, f interface{}) {
	if engine == nil {
		engine = qml.NewQJSEngine()
	}
	if engine.GlobalObject().Property("console").IsUndefined() {
		engine.NewGoType("console.log", func(s string) { println(s) })
	}
	engine.NewGoType(fn, f)
}
