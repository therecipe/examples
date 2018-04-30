package main

import (
	"os"
	"testing"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/testlib"
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
		window := NewCustomWindow(nil, 0)
		window.Show()

		input := widgets.NewQLineEditFromPointer(window.FindChild("someLineEdit", core.Qt__FindChildrenRecursively).Pointer())
		assert.Equal(t, "someInitialText", input.Text())

		button := widgets.NewQPushButtonFromPointer(window.FindChild("someButton", core.Qt__FindChildrenRecursively).Pointer())
		assert.Equal(t, "click me!", button.Text())

		list := testlib.NewQTestEventList()
		list.AddMouseClick(core.Qt__LeftButton, core.Qt__NoModifier, core.NewQPoint(), -1)
		list.Simulate(button)
		assert.Equal(t, "test text", input.Text())

		input.SetText("someOtherText")
		button.Click()
		assert.Equal(t, "test text", input.Text())

		window.Hide()
	})
}
