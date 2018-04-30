package main

import (
	"sync"

	"github.com/therecipe/qt/core"
)

var tRunner = NewTestRunner(nil)

type testRunner struct {
	core.QObject

	_ func(f func()) `signal:"runOnMain,auto"`
}

func (t *testRunner) runOnMain(f func()) { f() }

func (t *testRunner) Run(f func()) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	t.RunOnMain(func() {
		f()
		wg.Done()
	})
	wg.Wait()
}
