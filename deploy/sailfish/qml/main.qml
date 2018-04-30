import QtQuick 2.0
import QtQuick.Layouts 1.0	//ColumnLayout
import Sailfish.Silica 1.0	//ApplicationWindow

ApplicationWindow {
	allowedOrientations: Orientation.Portrait
	initialPage: Page {
		PageHeader {
			title: "Hello Deploy Example"
		}

		ColumnLayout {
			anchors.fill: parent

			Label {
				Layout.alignment: Qt.AlignCenter

				onLinkActivated: Qt.openUrlExternally(link)
				text: "<a href=\"https://github.com/therecipe/qt/wiki/Setting-the-Application-Icon#setting-the-icon-on-sailfishos\">Docs</a>"
			}

			Label {
				Layout.alignment: Qt.AlignCenter

				onLinkActivated: Qt.openUrlExternally(link)
				text: "<a href=\"https://sailfishos.org/develop/docs/silica/sailfish-silica-all.html\">Silica Reference</a>"
			}

			Label {
				Layout.alignment: Qt.AlignCenter

				onLinkActivated: Qt.openUrlExternally(link)
				text: "<a href=\"https://www.iconfinder.com/icons/52510/application_icon\">Icon credits</a>"
			}
		}
	}
}
