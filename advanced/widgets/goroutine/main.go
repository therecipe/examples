package main

import (
	"fmt"
	"os"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type CustomLabel struct {
	widgets.QLabel

	_ func(string) `signal:"updateTextFromGoroutine,auto(this.QLabel.setText)"` //TODO: support this.setText as well
}

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetFixedSize2(250, 200)
	window.SetWindowTitle("goroutine Example")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	labels := make([]*CustomLabel, 3)
	for i := range labels {
		label := NewCustomLabel(nil, 0)
		label.SetAlignment(core.Qt__AlignCenter)
		widget.Layout().AddWidget(label)
		labels[i] = label
	}

	button := widgets.NewQPushButton2("start!", nil)
	button.ConnectClicked(func(bool) {
		button.SetDisabled(true)
		for i, label := range labels {
			go func(i int, label *CustomLabel) {
				var tick int
				for range time.NewTicker(time.Duration((i+1)*25) * time.Millisecond).C {
					tick++
					label.UpdateTextFromGoroutine(fmt.Sprintf("%v %v", tick, time.Now().UTC().Format("15:04:05.0000")))
				}
			}(i, label)
		}
	})
	widget.Layout().AddWidget(button)

	window.Show()

	app.Exec()
}
