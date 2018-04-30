package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("Hello Deploy Example")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	widget.Layout().SetAlign(core.Qt__AlignCenter)
	window.SetCentralWidget(widget)

	for _, d := range [][]string{
		{"https://github.com/therecipe/qt/wiki/Setting-the-Application-Icon#setting-the-icon-on-android", "Docs"},
		{"https://doc.qt.io/qt-5/deployment-android.html", "Qt docs"},
		{"https://developer.android.com/studio/publish/app-signing#signing-manually", "Signing your app"},
		{"https://appicon.co", "Icon tool"},
		{"https://appiconizer.com", "Icon tool #2"},
		{"https://romannurik.github.io/AndroidAssetStudio/", "Icon tool #3"},
		{"https://www.iconfinder.com/icons/52510/application_icon", "Icon credits"},
	} {
		label := widgets.NewQLabel2(fmt.Sprintf("<a href=\"%v\">%v</a>", d[0], d[1]), nil, 0)
		label.SetOpenExternalLinks(true)
		widget.Layout().AddWidget(label)
	}

	window.Show()

	app.Exec()
}
