import QtQuick 2.10				//Item
import QtQuick.Controls 2.3		//Button

Item {
	width: 250
	height: 200

	Button {
		anchors.fill: parent
		text: ctxObject.someString
		onClicked: ctxObject.clicked()
	}
}