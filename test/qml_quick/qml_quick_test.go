package main

import (
	"os"
	"testing"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	go func() { app.Exit(m.Run()) }()
	app.Exec()
}

func TestWindow(t *testing.T) {
	tRunner.Run(func() {
		view := quick.NewQQuickView(nil)
		view.SetSource(core.NewQUrl3("qrc:/qml/Window_test.qml", 0))
		view.Show()

		localTestRunner := NewTestRunnerFromPointer(view.RootObject().FindChild("localTestRunner", core.Qt__FindChildrenRecursively).Pointer())

		input := view.RootObject().FindChild("someLineEdit", core.Qt__FindChildrenRecursively)
		assert.Equal(t, "someInitialText", input.Property("text").ToString())

		button := view.RootObject().FindChild("someButton", core.Qt__FindChildrenRecursively)
		assert.Equal(t, "click me!", button.Property("text").ToString())

		localTestRunner.CallQml(button, "clicked")
		assert.Equal(t, "test text", input.Property("text").ToString())

		view.Hide()
	})
}
