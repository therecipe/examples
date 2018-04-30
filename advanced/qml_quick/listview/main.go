package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func init() { CustomListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomListModel") }

type ListItem struct {
	firstName string
	lastName  string
}

type CustomListModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ func()                                  `signal:"remove,auto"`
	_ func(obj []*core.QVariant)              `signal:"add,auto"`
	_ func(firstName string, lastName string) `signal:"edit,auto"`

	modelData []ListItem
}

func (m *CustomListModel) init() {
	m.modelData = []ListItem{{"john", "doe"}, {"john", "bob"}}

	m.ConnectRowCount(m.rowCount)
	m.ConnectData(m.data)
}

func (m *CustomListModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *CustomListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}

	item := m.modelData[index.Row()]
	return core.NewQVariant14(fmt.Sprintf("%v %v", item.firstName, item.lastName))
}

func (m *CustomListModel) remove() {
	if len(m.modelData) == 0 {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), len(m.modelData)-1, len(m.modelData)-1)
	m.modelData = m.modelData[:len(m.modelData)-1]
	m.EndRemoveRows()
}

func (m *CustomListModel) add(item []*core.QVariant) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, ListItem{item[0].ToString(), item[1].ToString()})
	m.EndInsertRows()
}

func (m *CustomListModel) edit(firstName string, lastName string) {
	if len(m.modelData) == 0 {
		return
	}
	m.modelData[len(m.modelData)-1] = ListItem{firstName, lastName}
	m.DataChanged(m.Index(len(m.modelData)-1, 0, core.NewQModelIndex()), m.Index(len(m.modelData)-1, 0, core.NewQModelIndex()), []int{int(core.Qt__DisplayRole)})
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := widgets.NewQApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetTitle("listview Example")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.Show()

	app.Exec()
}
