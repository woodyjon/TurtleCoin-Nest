import QtQuick.Window 2.2
import QtQuick 2.7
import QtQuick.Controls 2.3
import QtQuick.Controls.Styles 1.4
import QtQuick.Dialogs 1.2
import QtQuick.Layouts 1.1

ApplicationWindow {
    id: window
    width: 1060
    height: 800
    maximumHeight: height
    maximumWidth: width
    minimumHeight: height
    minimumWidth: width
    color: "#333333"
    title: "TurtleCoin Nest"
    visible: true

    menuBar: MenuBar {
        Menu {
            id: menuWallet
            enabled: false
            title: "Wallet"
            MenuItem {
                text: "Open another wallet"
                onTriggered: {
                    walletScreen.showDialogWarningCloseWallet()
                }
            }
            MenuItem {
                text: "Backup wallet"
                onTriggered: {
                    QmlBridge.clickedButtonBackupWallet();
                }
            }
        }
    }

    OpenWallet {
        id: openWalletScreen
        anchors.fill: parent
    }

    Wallet {
        id: walletScreen
        anchors.fill: parent
        visible: false
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
            menuWallet.enabled = false;
            openWalletScreen.visible = true;
            walletScreen.visible = false;
            openWalletScreen.clearData();
        }

        onDisplayMainWalletScreen: {
            menuWallet.enabled = true;
            openWalletScreen.visible = false;
            walletScreen.visible = true;
            walletScreen.clearData();
        }
    }
}
