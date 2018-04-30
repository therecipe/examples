package main

import (
	"fmt"
	"os"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func init() {
	CustomLabel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomLabel")
	BridgeTemplate_QmlRegisterType2("CustomQmlTypes", 1, 0, "BridgeTemplate")
}

type CustomLabel struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"text"`
}

func (l *CustomLabel) init() {
	CustomLabels = append(CustomLabels, l)
}

type BridgeTemplate struct {
	core.QObject

	_ func() `signal:"clicked,auto"`
}

func (b *BridgeTemplate) clicked() {
	for i, label := range CustomLabels {
		go func(i int, label *CustomLabel) {
			var tick int
			for range time.NewTicker(time.Duration((i+1)*25) * time.Millisecond).C {
				tick++
				label.SetText(fmt.Sprintf("%v %v", tick, time.Now().UTC().Format("15:04:05.0000")))
			}
		}(i, label)
	}
}

var CustomLabels []*CustomLabel

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := widgets.NewQApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetTitle("goroutine Example")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.Show()

	app.Exec()
}
