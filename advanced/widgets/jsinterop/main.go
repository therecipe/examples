package main

import (
	"os"
	"unsafe"

	"github.com/therecipe/qt/widgets"

	"github.com/gopherjs/gopherjs/js"
)

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetFixedSize2(250, 200)
	window.SetWindowTitle("jsinterop Example")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	js.Global.Set("goFunc", func(title string) {
		widgets.QMessageBox_Information(nil, "OK", title, widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	button := js.Global.Call("eval", `
			var button = Module.NewQPushButton2("start!", 0);
			var func = function(bool) { alert("clicked the button: " + button.__internal_object__.QAbstractButton.Text()); goFunc("fromJS"); };
		
			button.__internal_object__.QAbstractButton.ConnectClicked(func);

			button.__internal_object__.Pointer();
		`)

	widget.Layout().AddWidget(widgets.NewQPushButtonFromPointer(unsafe.Pointer(button.Unsafe())))

	window.Show()

	app.Exec()
}
