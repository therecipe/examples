//port of: https://github.com/therecipe/examples/blob/master/basic/widgets/main.go

package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	file := core.NewQFile2(":/qml/basic.js")
	if file.Open(core.QIODevice__ReadOnly) {
		evalJS(file.ReadAll().ConstData())
		file.Close()
	}

	widgets.QApplication_Exec()
}
