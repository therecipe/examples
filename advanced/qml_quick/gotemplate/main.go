package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func init() {
	ItemTemplate_QmlRegisterType2("CustomQmlTypes", 1, 0, "ItemTemplate")
	QtObjectTemplate_QmlRegisterType2("CustomQmlTypes", 1, 0, "QtObjectTemplate")
}

//same as http://doc.qt.io/qt-5/qml-qtquick-item.html
type ItemTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ string `property:"someString"`

	_ *QtObjectTemplate `property:"someNestedQtObject"`
}

func (t *ItemTemplate) init() {
	t.ConnectComponentComplete(t.componentComplete)
}

//this is the earliest possible point you can access the initial properties
func (t *ItemTemplate) componentComplete() {
	println("from Item:", t.SomeString())
	println("from Item (QtObject):", t.SomeNestedQtObject().SomeString())
	t.ComponentCompleteDefault()
}

//same as http://doc.qt.io/qt-5/qml-qtqml-qtobject.html
type QtObjectTemplate struct {
	core.QObject

	_ string `property:"someString"`

	_ func() `signal:"componentComplete,auto"`
}

//this is the earliest possible point you can access the initial properties (using a workaround in qml)
func (t *QtObjectTemplate) componentComplete() {
	println("from QtObject:", t.SomeString())
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := widgets.NewQApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetTitle("gotemplate Example")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.Show()

	app.Exec()
}
