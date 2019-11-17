package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("jsinterop Example")
	window.SetFixedSize2(250, 200)
	window.Move2(100, 100)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	createGlobalFunction("goFunc", func(title string) {
		widgets.QMessageBox_Information(nil, "OK", title, widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	button := evalJSToPtr(`(function(){
	   var button = widgets.NewQPushButton2("start!");
	   button.ConnectClicked(function(bool) { console.log("clicked the first button: " + button.Text()); goFunc("fromJS1"); });
	   return button.Pointer();
	   })()`)
	widget.Layout().AddWidget(widgets.NewQPushButtonFromPointer(button))

	button2 := evalJSToPtr(`(function(){
	   var button = widgets.NewQPushButton2("start2!");
	   button.ConnectClicked(function(bool) { console.log("clicked the second button: " + button.Text()); goFunc("fromJS2"); });
	   return button.Pointer();
	   })()`)
	widget.Layout().AddWidget(widgets.NewQPushButtonFromPointer(button2))

	window.Show()

	app.Exec()
}
