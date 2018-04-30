package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

type CtxObject struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"someString"`

	_ func() `signal:"clicked,auto"`
}

func (t *CtxObject) init() {
	t.SetSomeString("click me!\nand look into the console")
}

func (t *CtxObject) clicked() {
	println("clicked qml button")
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := widgets.NewQApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetTitle("ctxproperty Example")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.RootContext().SetContextProperty("ctxObject", NewCtxObject(nil))
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.Show()

	app.Exec()
}
