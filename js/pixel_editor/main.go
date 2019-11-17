//port of: https://github.com/therecipe/qt/blob/master/internal/examples/widgets/pixel_editor/pixel_editor.go

package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	file := core.NewQFile2(":/qml/pixel_editor.js")
	if file.Open(core.QIODevice__ReadOnly) {
		evalJS(file.ReadAll().ConstData())
		file.Close()
	}

	widgets.QApplication_Exec()
}
