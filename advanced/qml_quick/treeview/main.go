package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func init() { CustomTreeModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomTreeModel") }

const (
	FirstName = int(core.Qt__UserRole) + 1<<iota
	LastName
)

type TreeItem struct {
	core.QObject

	_ func() `constructor:"init"`

	_itemData   []string
	_parentItem *TreeItem

	_childItems []*TreeItem
}

func (i *TreeItem) init() {
	i.ConnectDestroyTreeItem(i.destroyTreeItem)
}

func (i *TreeItem) initWith(data []string) *TreeItem {
	i._itemData = data
	return i
}

func (i *TreeItem) appendChild(child *TreeItem) {
	child._parentItem = i
	i._childItems = append(i._childItems, child)
}

func (i *TreeItem) child(row int) *TreeItem {
	return i._childItems[row]
}

func (i *TreeItem) childCount() int {
	return len(i._childItems)
}

func (i *TreeItem) columnCount() int {
	return len(i._itemData)
}

func (i *TreeItem) data(column int) string {
	return i._itemData[column]
}

func (i *TreeItem) row() int {
	if i._parentItem != nil {
		for index, item := range i._parentItem._childItems {
			if item.Pointer() == i.Pointer() {
				return index
			}
		}
	}
	return 0
}

func (i *TreeItem) parentItem() *TreeItem {
	return i._parentItem
}

func (i *TreeItem) destroyTreeItem() {
	for _, child := range i._childItems {
		child.DestroyTreeItem()
	}
	i.DestroyTreeItemDefault()
}

type CustomTreeModel struct {
	core.QAbstractItemModel

	_ func() `constructor:"init"`

	_ func()                                  `signal:"remove,auto"`
	_ func(item []*core.QVariant)             `signal:"add,auto"`
	_ func(firstName string, lastName string) `signal:"edit,auto"`

	rootItem *TreeItem
}

func (m *CustomTreeModel) init() {
	m.rootItem = NewTreeItem(nil).initWith([]string{"FirstName", "LastName"})
	m.rootItem.appendChild(NewTreeItem(nil).initWith([]string{"john", "doe"}))

	firstChild := NewTreeItem(nil).initWith([]string{"john", "bob"})
	secondChild := NewTreeItem(nil).initWith([]string{"jim", "bob"})
	thirdChild := NewTreeItem(nil).initWith([]string{"jimmy", "bob"})

	firstChild.appendChild(secondChild)
	secondChild.appendChild(thirdChild)
	m.rootItem.appendChild(firstChild)

	m.ConnectIndex(m.index)
	m.ConnectParent(m.parent)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *CustomTreeModel) index(row int, column int, parent *core.QModelIndex) *core.QModelIndex {
	if !m.HasIndex(row, column, parent) {
		return core.NewQModelIndex()
	}

	var parentItem *TreeItem
	if !parent.IsValid() {
		parentItem = m.rootItem
	} else {
		parentItem = NewTreeItemFromPointer(parent.InternalPointer())
	}

	childItem := parentItem.child(row).Pointer()
	if childItem != nil {
		return m.CreateIndex(row, column, childItem)
	}
	return core.NewQModelIndex()
}

func (m *CustomTreeModel) parent(index *core.QModelIndex) *core.QModelIndex {
	if !index.IsValid() {
		return core.NewQModelIndex()
	}

	item := NewTreeItemFromPointer(index.InternalPointer())
	parentItem := item.parentItem()

	if parentItem.Pointer() == m.rootItem.Pointer() {
		return core.NewQModelIndex()
	}

	return m.CreateIndex(parentItem.row(), 0, parentItem.Pointer())
}

func (m *CustomTreeModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		FirstName: core.NewQByteArray2("FirstName", -1),
		LastName:  core.NewQByteArray2("LastName", -1),
	}
}

func (m *CustomTreeModel) rowCount(parent *core.QModelIndex) int {
	if !parent.IsValid() {
		return m.rootItem.childCount()
	}
	return NewTreeItemFromPointer(parent.InternalPointer()).childCount()
}

func (m *CustomTreeModel) columnCount(parent *core.QModelIndex) int {
	if !parent.IsValid() {
		return m.rootItem.columnCount()
	}
	return NewTreeItemFromPointer(parent.InternalPointer()).columnCount()
}

func (m *CustomTreeModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	item := NewTreeItemFromPointer(index.InternalPointer())
	switch role {
	case FirstName:
		return core.NewQVariant14(item._itemData[0])
	case LastName:
		return core.NewQVariant14(item._itemData[1])
	}
	return core.NewQVariant()
}

func (m *CustomTreeModel) remove() {
	if m.rootItem.childCount() == 0 {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), len(m.rootItem._childItems)-1, len(m.rootItem._childItems)-1)
	item := m.rootItem._childItems[len(m.rootItem._childItems)-1]
	m.rootItem._childItems = m.rootItem._childItems[:len(m.rootItem._childItems)-1]
	m.EndRemoveRows()
	item.DestroyTreeItem()
}

func (m *CustomTreeModel) add(item []*core.QVariant) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.rootItem._childItems), len(m.rootItem._childItems))
	m.rootItem.appendChild(NewTreeItem(nil).initWith([]string{item[0].ToString(), item[1].ToString()}))
	m.EndInsertRows()
}

func (m *CustomTreeModel) edit(firstName string, lastName string) {
	if m.rootItem.childCount() == 0 {
		return
	}
	m.BeginResetModel()
	item := m.rootItem._childItems[len(m.rootItem._childItems)-1]
	item._itemData = []string{firstName, lastName}
	m.EndResetModel()

	//TODO:
	//ideally DataChanged should be used instead, but it doesn't seem to work ...
	//if you search for "qml treeview datachanged" online
	//it will just lead you to tons of unresolved issues
	//m.DataChanged(m.Index(item.row(), 0, core.NewQModelIndex()), m.Index(item.row(), 1, core.NewQModelIndex()), []int{FirstName, LastName})
	//feel free to send a PR, if you got it working somehow :)
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := widgets.NewQApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetTitle("treeview Example")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.Show()

	app.Exec()
}
