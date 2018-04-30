package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type TableItem struct {
	firstName string
	lastName  string
}

type CustomTableModel struct {
	core.QAbstractTableModel

	_ func() `constructor:"init"`

	_ func()                                  `signal:"remove,auto"`
	_ func(item TableItem)                    `signal:"add,auto"`
	_ func(firstName string, lastName string) `signal:"edit,auto"`

	modelData []TableItem
}

func (m *CustomTableModel) init() {
	m.modelData = []TableItem{{"john", "doe"}, {"john", "bob"}}

	m.ConnectHeaderData(m.headerData)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *CustomTableModel) headerData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) || orientation == core.Qt__Vertical {
		return m.HeaderDataDefault(section, orientation, role)
	}

	switch section {
	case 0:
		return core.NewQVariant14("FirstName")
	case 1:
		return core.NewQVariant14("LastName")
	}
	return core.NewQVariant()
}

func (m *CustomTableModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *CustomTableModel) columnCount(*core.QModelIndex) int {
	return 2
}

func (m *CustomTableModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}

	item := m.modelData[index.Row()]
	switch m.HeaderData(index.Column(), core.Qt__Horizontal, role).ToString() {
	case "FirstName":
		return core.NewQVariant14(item.firstName)
	case "LastName":
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

func (m *CustomTableModel) add(item TableItem) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, item)
	m.EndInsertRows()
}

func (m *CustomTableModel) edit(firstName string, lastName string) {
	if len(m.modelData) == 0 {
		return
	}
	m.modelData[len(m.modelData)-1] = TableItem{firstName, lastName}
	m.DataChanged(m.Index(len(m.modelData)-1, 0, core.NewQModelIndex()), m.Index(len(m.modelData)-1, 1, core.NewQModelIndex()), []int{int(core.Qt__DisplayRole)})
}

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("tableview Example")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	tableview := widgets.NewQTableView(nil)
	model := NewCustomTableModel(nil)
	tableview.SetModel(model)
	widget.Layout().AddWidget(tableview)

	remove := widgets.NewQPushButton2("remove last item", nil)
	remove.ConnectClicked(func(bool) {
		model.Remove()
	})
	widget.Layout().AddWidget(remove)

	add := widgets.NewQPushButton2("add new item", nil)
	add.ConnectClicked(func(bool) {
		model.Add(TableItem{"john", "doe"})
	})
	widget.Layout().AddWidget(add)

	edit := widgets.NewQPushButton2("edit last item", nil)
	edit.ConnectClicked(func(bool) {
		model.Edit("bob", "omb")
	})
	widget.Layout().AddWidget(edit)

	window.Show()

	app.Exec()
}
