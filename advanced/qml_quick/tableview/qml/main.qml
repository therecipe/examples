import QtQuick 2.10				//Item
import QtQuick.Controls 1.4		//TableView
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout
import CustomQmlTypes 1.0		//CustomTableModel

Item {
	width: 250
	height: 200

	ColumnLayout {
		anchors.fill: parent

		TableView {
			id: tableview

			Layout.fillWidth: true
			Layout.fillHeight: true

			model: CustomTableModel{}

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
			onClicked: tableview.model.remove()
		}

		Button {
			Layout.fillWidth: true

			text: "add new item"
			onClicked: tableview.model.add(["john", "doe"])
		}

		Button {
			Layout.fillWidth: true

			text: "edit last item"
			onClicked: tableview.model.edit("bob", "omb")
		}
	}
}