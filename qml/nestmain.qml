import QtQuick.Window 2.2
import QtQuick 2.7
import QtQuick.Controls 2.3
import QtQuick.Controls.Styles 1.4
import QtQuick.Dialogs 1.2
import QtQuick.Layouts 1.1

ApplicationWindow {
    id: window

    property var windowWidth: 1060
    property var windowHeight: 755
    property var minWindowHeight: 500

    width: windowWidth
    height: windowHeight
    minimumWidth: windowWidth
    minimumHeight: minWindowHeight
    color: "#333333"
    title: "TurtleCoin Nest"
    visible: true

    Flickable{
        interactive: true
        boundsBehavior: Flickable.StopAtBounds
        contentWidth: windowWidth
        contentHeight: windowHeight
        width: parent.width
        height: parent.height

        OpenWallet {
            id: openWalletScreen
            anchors.fill: parent
        }

        Wallet {
            id: walletScreen
        }

        Settings {
            id: settingsScreen
        }
    }

    Dialog {
        id: dialogInfo
        title: "Info"
        standardButtons: StandardButton.Ok
        width: 800
        height: 150
        modality: Qt.WindowModal

        Text {
            id: textDialogInfo
            text: ""
            font.family: "Arial"
        }

        function show(title, msg) {
            dialogInfo.title = title
            textDialogInfo.text = msg
            dialogInfo.open()
        }

        function showError(msg) {
            dialogInfo.show("Error", msg)
        }

        Connections{
            target: QmlBridge
            onDisplayErrorDialog: {
                dialogInfo.showError(errorMessage)
            }
        }
    }

    Connections {
        target: QmlBridge

        onDisplayOpenWalletScreen: {
            openWalletScreen.clearData();
            walletScreen.hide();
        }

        onDisplayMainWalletScreen: {
            walletScreen.clearData();
            walletScreen.show();
        }

        onDisplaySettingsScreen: {
            settingsScreen.show();
        }

        onHideSettingsScreen: {
            settingsScreen.hide();
        }
    }
}
