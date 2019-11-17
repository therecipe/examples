// +build !js wasm

package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

var engine *qml.QJSEngine

func evalJS(s string) {
	if engine == nil {
		engine = qml.NewQJSEngine()
	}
	if engine.GlobalObject().Property("setInterval").IsUndefined() {
		engine.NewGoType("setInterval", func(f func(), msec int) {
			t := core.NewQTimer(nil)
			t.ConnectTimeout(f)
			t.Start(msec)
		})
	}
	engine.Evaluate(s, "", 0)
}
