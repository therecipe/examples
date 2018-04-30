package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func init() { CustomTableModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomTableModel") }

const (
	FirstName = int(core.Qt__UserRole) + 1<<iota
	LastName
)

type TableItem struct {
	firstName string
	lastName  string
}

type CustomTableModel struct {
	core.QAbstractTableModel

	_ func() `constructor:"init"`

	_ func()                                  `signal:"remove,auto"`
	_ func(item []*core.QVariant)             `signal:"add,auto"`
	_ func(firstName string, lastName string) `signal:"edit,auto"`

	modelData []TableItem
}

func (m *CustomTableModel) init() {
	m.modelData = []TableItem{{"john", "doe"}, {"john", "bob"}}

	m.ConnectRoleNames(m.roleNames)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *CustomTableModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		FirstName: core.NewQByteArray2("FirstName", -1),
		LastName:  core.NewQByteArray2("LastName", -1),
	}
}

func (m *CustomTableModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *CustomTableModel) columnCount(*core.QModelIndex) int {
	return 2
}

func (m *CustomTableModel) data(index *core.QModelIndex, role int) *core.QVariant {
	item := m.modelData[index.Row()]
	switch role {
	case FirstName:
		return core.NewQVariant14(item.firstName)
	case LastName:
		return core.NewQVariant14(item.lastName)
	}
	return core.NewQVariant()
}

func (m *CustomTableModel) remove() {
	if len(m.modelData) == 0 {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), len(m.modelData)-1, len(m.modelData)-1)
	m.modelData = m.modelData[:len(m.modelData)-1]
	m.EndRemoveRows()
}

func (m *CustomTableModel) add(item []*core.QVariant) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, TableItem{item[0].ToString(), item[1].ToString()})
	m.EndInsertRows()
}

func (m *CustomTableModel) edit(firstName string, lastName string) {
	if len(m.modelData) == 0 {
		return
	}
	m.modelData[len(m.modelData)-1] = TableItem{firstName, lastName}
	m.DataChanged(m.Index(len(m.modelData)-1, 0, core.NewQModelIndex()), m.Index(len(m.modelData)-1, 1, core.NewQModelIndex()), []int{FirstName, LastName})
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := widgets.NewQApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetTitle("tableview Example")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.Show()

	app.Exec()
}
