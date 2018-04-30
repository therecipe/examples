import QtQuick 2.10				//Component
import QtQuick.Controls 2.3		//Label
import CustomQmlTypes 1.0		//ItemTemplate and QtObjectTemplate

ItemTemplate {
	id: root

	someString: "hello world"

	someNestedQtObject: QtObjectTemplate {
		someString: root.someString
		Component.onCompleted: componentComplete()
	}

	width: 250
	height: 200

	Label {
		anchors.fill: parent
		text: "look into the console"
	}
}