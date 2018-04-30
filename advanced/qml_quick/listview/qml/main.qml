import QtQuick 2.10				//ListView
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout
import CustomQmlTypes 1.0		//CustomListModel

Item {
	width: 250
	height: 200

	ColumnLayout {
		anchors.fill: parent

		ListView {
			id: listview

			Layout.fillWidth: true
			Layout.fillHeight: true

			model: CustomListModel{}
			delegate: Text {
				text: display
			}
		}

		Button {
			Layout.fillWidth: true

			text: "remove last item"
			onClicked: listview.model.remove()
		}

		Button {
			Layout.fillWidth: true

			text: "add new item"
			onClicked: listview.model.add(["john", "doe"])
		}

		Button {
			Layout.fillWidth: true

			text: "edit last item"
			onClicked: listview.model.edit("bob", "omb")
		}
	}
}