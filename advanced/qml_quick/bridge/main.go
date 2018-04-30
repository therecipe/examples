package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func init() {
	ItemTemplate_QmlRegisterType2("CustomQmlTypes", 1, 0, "ItemTemplate")
}

type ItemTemplate struct {
	quick.QQuickItem

	_ string `property:"someString"`

	_ func(bool, []bool)                                                `signal:"sendBool,auto"`
	_ func(int, []int)                                                  `signal:"sendInt,auto"`
	_ func(float32)                                                     `signal:"sendFloat,auto"`
	_ func(float64, []float64)                                          `signal:"sendDouble,auto"`
	_ func(string, []string)                                            `signal:"sendString,auto"`
	_ func(error, []error)                                              `signal:"sendError,auto"`
	_ func(*core.QVariant, []*core.QVariant, map[string]*core.QVariant) `signal:"sendVariantListMap,auto"`
	_ func(*ItemTemplate)                                               `signal:"sendItemTemplate,auto"`
	_ func(*quick.QQuickItem)                                           `signal:"sendItem,auto"`
	_ func(*core.QObject, []*core.QObject)                              `signal:"sendObject,auto"`
	//...
}

func (t *ItemTemplate) sendBool(a bool, b []bool) {
	fmt.Println("sendBool:", a, b)
}

func (t *ItemTemplate) sendInt(a int, b []int) {
	fmt.Println("sendInt:", a, b)
}

func (t *ItemTemplate) sendFloat(a float32) {
	fmt.Println("sendFloat:", a)
}

func (t *ItemTemplate) sendDouble(a float64, b []float64) {
	fmt.Println("sendDouble:", a, b)
}

func (t *ItemTemplate) sendString(a string, b []string) {
	fmt.Println("sendString:", a, b)
}

func (t *ItemTemplate) sendError(a error, b []error) {
	fmt.Println("sendError:", a, b)
}

func (t *ItemTemplate) sendVariantListMap(a *core.QVariant, b []*core.QVariant, c map[string]*core.QVariant) {
	fmt.Println("sendVariantListMap:", a.ToBool(), b[0].ToDouble(false), b[1].ToString(), c)
}

func (t *ItemTemplate) sendItemTemplate(a *ItemTemplate) {
	fmt.Println("sendItemTemplate:", a.SomeString())
}

func (t *ItemTemplate) sendItem(a *quick.QQuickItem) {
	fmt.Println("sendItem:", NewItemTemplateFromPointer(a.Pointer()).SomeString())
}

func (t *ItemTemplate) sendObject(a *core.QObject, b []*core.QObject) {
	fmt.Println("sendObject:", NewItemTemplateFromPointer(a.Pointer()).SomeString(), NewItemTemplateFromPointer(b[0].Pointer()).SomeString(), NewItemTemplateFromPointer(b[1].Pointer()).SomeString())
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := widgets.NewQApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetTitle("bridge Example")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.Show()

	app.Exec()
}
