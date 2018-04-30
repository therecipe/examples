package main

import (
	"sync"

	"github.com/therecipe/qt/core"
)

func init() { testRunner_QmlRegisterType2("GoTest", 1, 0, "TestRunner") }

var tRunner = NewTestRunner(nil)

type testRunner struct {
	core.QObject

	_ func(f func())                                             `signal:"runOnMain,auto"`
	_ func(obj *core.QObject, sig string)                        `signal:"callQml"`
	_ func(obj *core.QObject, sig string, args []*core.QVariant) `signal:"callQmlWithArgs"`
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
