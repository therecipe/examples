import QtQuick 2.7			//Item
import QtQuick.Controls 2.1	//Dialog

Item {
	id: window

	Column {
		anchors.centerIn: parent

		TextField {
			id: input
	
			anchors.horizontalCenter: parent.horizontalCenter
			placeholderText: "Write something ..."
		}

		Button {
			anchors.horizontalCenter: parent.horizontalCenter
			text: "and click me!"
			onClicked: dialog.open()
		}
	}

	Dialog {
		id: dialog

		x: (window.width - width) * 0.5
		y: (window.height - height) * 0.5

		contentWidth: window.width * 0.5
		contentHeight: window.height * 0.25
		standardButtons: Dialog.Ok

		contentItem: Label {
			text: input.text
		}
	}
}