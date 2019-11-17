// +build !js wasm

package main

import "github.com/therecipe/qt/qml"

var engine *qml.QJSEngine

func evalJS(s string) {
	if engine == nil {
		engine = qml.NewQJSEngine()
	}
	engine.Evaluate(s, "", 0)
}
