import QtQuick 2.0			//Item
import QtQuick.Layouts 1.0	//ColumnLayout

Item {
	ColumnLayout {
		anchors.fill: parent

		Label {
			Layout.alignment: Qt.AlignCenter

			onLinkActivated: Qt.openUrlExternally(link)
			text: "<a href=\"https://github.com/therecipe/qt/wiki/Setting-the-Application-Icon\">Docs</a>"
		}

		Label {
			Layout.alignment: Qt.AlignCenter

			onLinkActivated: Qt.openUrlExternally(link)
			text: "<a href=\"https://docs.ubports.com/en/latest/appdev/index.html\">UBports Docs</a>"
		}

		Label {
			Layout.alignment: Qt.AlignCenter

			onLinkActivated: Qt.openUrlExternally(link)
			text: "<a href=\"https://www.iconfinder.com/icons/52510/application_icon\">Icon credits</a>"
		}
	}
}
