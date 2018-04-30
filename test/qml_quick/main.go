package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
)

func main() {

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	widgets.NewQApplication(len(os.Args), os.Args)

	quickcontrols2.QQuickStyle_SetStyle("Material")

	view := quick.NewQQuickView(nil)
	view.SetMinimumSize(core.NewQSize2(250, 200))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetTitle("Test QML/Quick Example")
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.Show()

	widgets.QApplication_Exec()
}
