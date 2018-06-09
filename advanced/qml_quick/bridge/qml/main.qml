import QtQuick 2.0				//needed for js
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout
import CustomQmlTypes 1.0		//ItemTemplate

ItemTemplate {
	id: root

	width: 250
	height: 450

	someString: "ItemTemplateString"

	ColumnLayout {
		anchors.fill: parent

		Button {
			Layout.fillWidth: true

			text: "sendBool"
			onClicked: {
				text = "look into the console"
				root.sendBool(true, [true, true])
			}
		}

		Button {
			Layout.fillWidth: true

			text: "sendInt"
			onClicked: {
				text = "look into the console"
				root.sendInt(123, [456, 789])
			}
		}

		Button {
			Layout.fillWidth: true

			text: "sendFloat"
			onClicked: {
				text = "look into the console"
				root.sendFloat(1.23)
			}
		}

		Button {
			Layout.fillWidth: true

			text: "sendDouble"
			onClicked: {
				text = "look into the console"
				root.sendDouble(1.23, [4.56, 7.89])
			}
		}

		Button {
			Layout.fillWidth: true

			text: "sendString"
			onClicked: {
				text = "look into the console"
				root.sendString("hello", ["hello", "world"])
			}
		}

		Button {
			Layout.fillWidth: true

			text: "sendError"
			onClicked: {
				text = "look into the console"
				root.sendError("error1", ["error2", "error3"])
			}
		}

		Button {
			Layout.fillWidth: true

			text: "sendVariantListMap"
			onClicked: {
				text = "look into the console"
				root.sendVariantListMap(true, [1.23, "hello"],  {"A": true, "B": 1.23, "C": "hello", "D": root, "E": [root, root]})
			}
		}

		Button {
			Layout.fillWidth: true

			text: "sendItemTemplate"
			onClicked: {
				text = "look into the console"
				root.sendItemTemplate(root)
			}
		}

		Button {
			Layout.fillWidth: true

			text: "sendItem"
			onClicked: {
				text = "look into the console"
				root.sendItem(root)
			}
		}

		Button {
			Layout.fillWidth: true

			text: "sendObject"
			onClicked: {
				text = "look into the console"
				root.sendObject(root, [root, root])
			}
		}
	}
}