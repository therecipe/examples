import QtQuick 2.7			//Column
import QtQuick.Controls 2.1	//Button

Item {
	Column {
		anchors.centerIn: parent

		TextField {
			id: input
			objectName: "someLineEdit"
	
			anchors.horizontalCenter: parent.horizontalCenter
			text: "someInitialText"
		}

		Button {
			objectName: "someButton"

			anchors.horizontalCenter: parent.horizontalCenter
			text: "click me!"
			onClicked: input.text = "test text"
		}
	}
}