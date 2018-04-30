package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	NewCustomWindow(nil, 0).Show()

	widgets.QApplication_Exec()
}
