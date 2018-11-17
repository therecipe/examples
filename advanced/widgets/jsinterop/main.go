// +build js,!wasm

package main

import (
	"os"
	"unsafe"

	"github.com/gopherjs/gopherjs/js"

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

	js.Global.Set("goFunc", func(title string) {
		widgets.QMessageBox_Information(nil, "OK", title, widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	button := js.Global.Call("eval", `
	   "use strict";

	   var button = widgets.NewQPushButton2("start!", 0);
	   button.ConnectClicked(function(bool) { alert("clicked the first button: " + button.Text()); goFunc("fromJS1"); });
	   button.Pointer();
	   			`)
	widget.Layout().AddWidget(widgets.NewQPushButtonFromPointer(unsafe.Pointer(button.Unsafe())))

	button2 := js.Global.Call("eval", `
	   "use strict";

	   var button = widgets.NewQPushButton2("start2!", 0);
	   button.ConnectClicked(function(bool) { alert("clicked the first button: " + button.Text()); goFunc("fromJS2"); });
	   button.Pointer();
	   			`)
	widget.Layout().AddWidget(widgets.NewQPushButtonFromPointer(unsafe.Pointer(button2.Unsafe())))

	window.Show()

	app.Exec()
}
