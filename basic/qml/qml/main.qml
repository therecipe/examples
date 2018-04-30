import QtQuick 2.7			//ApplicationWindow
import QtQuick.Controls 2.1	//Dialog

ApplicationWindow {
	id: window

	visible: true
	title: "Hello QML Example"
	minimumWidth: 250
	minimumHeight: 200

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