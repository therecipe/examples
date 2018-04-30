package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/sailfish"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := sailfish.SailfishApp_Application(len(os.Args), os.Args)

	view := sailfish.SailfishApp_CreateView()
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	app.Exec()
}
