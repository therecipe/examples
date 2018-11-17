//port of: https://github.com/therecipe/qt/blob/master/internal/examples/widgets/line_edits/line_edits.go

// +build js,!wasm

package main

import (
	"os"

	"github.com/gopherjs/gopherjs/js"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	file := core.NewQFile2(":/qml/main.js")
	if file.Open(core.QIODevice__ReadOnly) {
		js.Global.Call("eval", file.ReadAll().ConstData())
		file.Close()
	}

	widgets.QApplication_Exec()
}
