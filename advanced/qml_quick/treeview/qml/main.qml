import QtQml.Models 2.2			//needed for js
import QtQuick 2.10				//Item
import QtQuick.Controls 1.4		//TreeView
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout
import CustomQmlTypes 1.0		//CustomTableModel

Item {
	width: 250
	height: 200

	ColumnLayout {
		anchors.fill: parent

		TreeView {
			id: treeview

			Layout.fillWidth: true
			Layout.fillHeight: true

			model: CustomTreeModel{}

			TableViewColumn {
				role: "FirstName"
				title: role
			}

			TableViewColumn {
				role: "LastName"
				title: role
			}
		}

		Button {
			Layout.fillWidth: true

			text: "remove last item"
			onClicked: treeview.model.remove()
		}

		Button {
			Layout.fillWidth: true

			text: "add new item"
			onClicked: treeview.model.add(["john", "doe"])
		}

		Button {
			Layout.fillWidth: true

			text: "edit last item"
			onClicked: treeview.model.edit("bob", "omb")
		}
	}
}