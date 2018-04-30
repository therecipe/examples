package main

import (
	"github.com/therecipe/qt/widgets"
)

type CustomWindow struct {
	widgets.QMainWindow

	_ func() `constructor:"init"`
}

func (window *CustomWindow) init() {
	window.SetObjectName("someCustomWindow")
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("Test Widgets Example")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	input := widgets.NewQLineEdit(nil)
	input.SetObjectName("someLineEdit")
	input.SetText("someInitialText")
	widget.Layout().AddWidget(input)

	button := widgets.NewQPushButton2("click me!", nil)
	button.SetObjectName("someButton")
	button.ConnectClicked(func(bool) {
		input.SetText("test text")
	})
	widget.Layout().AddWidget(button)
}
