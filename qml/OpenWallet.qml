import QtQuick.Window 2.2
import QtQuick 2.7
import QtQuick.Controls 2.3
import QtQuick.Controls.Styles 1.4
import QtQuick.Dialogs 1.2
import QtQuick.Layouts 1.1

Rectangle {
    id: rectangleOpenWallet
    anchors.fill: parent
    color: "transparent"

    Text {
        id: textOpenWalletDescr
        color: "#ffffff"
        text: qsTr("Welcome to your TurtleCoin wallet. Choose one of the options below. If you are new to TurtleCoin, create a new wallet.")
        anchors.right: parent.right
        anchors.rightMargin: 15
        anchors.left: parent.left
        anchors.leftMargin: 40
        anchors.top: parent.top
        anchors.topMargin: 40
        font.pixelSize: 14
        verticalAlignment: Text.AlignBottom
        font.family: "Arial"
        font.bold: true
        horizontalAlignment: Text.AlignLeft
    }

    // section open existing wallet
    Rectangle {
        id: rectangleOpenExistingWallet
        color: "transparent"
        anchors.left: parent.left
        anchors.leftMargin: 10
        anchors.right: parent.right
        anchors.rightMargin: 10
        anchors.top: textOpenWalletDescr.bottom
        anchors.topMargin: 50
        height: 120
        radius: 7

        Text {
            id: textOpenExistingWalletDescr
            color: "#ffffff"
            text: qsTr("Open an existing wallet")
            horizontalAlignment: Text.AlignLeft
            verticalAlignment: Text.AlignBottom
            font.bold: true
            font.pixelSize: 16
            font.family: "Arial"
            anchors.top: parent.top
            anchors.topMargin: 10
            anchors.left: parent.left
            anchors.leftMargin: 40
        }

        Text {
            id: textExistingWalletPathDescr
            color: "#ffffff"
            text: qsTr("Path to wallet file")
            horizontalAlignment: Text.AlignLeft
            verticalAlignment: Text.AlignBottom
            font.pixelSize: 14
            font.family: "Arial"
            anchors.top: textOpenExistingWalletDescr.bottom
            anchors.topMargin: 20
            anchors.left: parent.left
            anchors.leftMargin: 80
        }

        Rectangle {
            id: rectangleTextInputExistingWalletPath
            color: "#555555"
            height: 25
            width: 360
            anchors.top: textExistingWalletPathDescr.bottom
            anchors.topMargin: 15
            anchors.left: textExistingWalletPathDescr.left
            anchors.leftMargin: 0
            radius: 3

            TextInput {
                id: textInputExistingWalletPath
                readOnly: true
                anchors.fill: parent
                color: "#cfcfcf"
                text: qsTr("")
                rightPadding: 5
                leftPadding: 5
                selectionColor: "#eeeeee"
                selectedTextColor: "#999999"
                selectByMouse: true
                clip: true
                font.family: "Arial"
                horizontalAlignment: Text.AlignLeft
                font.pixelSize: 14
                verticalAlignment: Text.AlignVCenter

                onTextChanged: {
                    buttonOpenWallet.enabled = textInputExistingWalletPath.text != "" && textInputExistingWalletPassword.text != ""
                }
            }

            MouseArea {
                id: mouseAreaExistingWalletPath
                anchors.fill: parent
                onClicked: {
                    dialogChooseWalletFile.show();
                }
            }
        }

        Text {
            id: textExistingWalletPasswordDescr
            color: "#ffffff"
            text: qsTr("Password")
            horizontalAlignment: Text.AlignLeft
            verticalAlignment: Text.AlignBottom
            font.pixelSize: 14
            font.family: "Arial"
            anchors.top: textExistingWalletPathDescr.top
            anchors.topMargin: 0
            anchors.left: parent.left
            anchors.leftMargin: 520
        }

        Rectangle {
            id: rectangleTextInputExistingWalletPassword
            color: "#555555"
            height: 25
            width: 200
            anchors.top: rectangleTextInputExistingWalletPath.top
            anchors.topMargin: 0
            anchors.left: textExistingWalletPasswordDescr.left
            anchors.leftMargin: 0
            radius: 3

            TextInput {
                id: textInputExistingWalletPassword
                echoMode: TextInput.Password
                anchors.fill: parent
                color: "#cfcfcf"
                text: qsTr("")
                rightPadding: 5
                leftPadding: 5
                selectionColor: "#eeeeee"
                selectedTextColor: "#999999"
                selectByMouse: true
                clip: true
                font.family: "Arial"
                horizontalAlignment: Text.AlignLeft
                font.pixelSize: 14
                verticalAlignment: Text.AlignVCenter

                onTextChanged: {
                    buttonOpenWallet.enabled = textInputExistingWalletPath.text != "" && textInputExistingWalletPassword.text != ""
                }
            }

        }

        Button {
            id: buttonOpenWallet
            text: qsTr("OPEN")
            anchors.right: parent.right
            anchors.rightMargin: 60
            anchors.bottom: parent.bottom
            anchors.bottomMargin: 15
            enabled: false

            contentItem: Text {
                text: buttonOpenWallet.text
                font.bold: true
                font.pointSize: 20
                font.family: "Arial"
                opacity: enabled ? 1.0 : 0.3
                color: buttonOpenWallet.down ? "#dddddd" : "#ffffff"
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
                elide: Text.ElideRight
            }

            background: Rectangle {
                implicitWidth: 140
                implicitHeight: 60
                opacity: enabled ? 1 : 0.3
                radius: 6
                color: buttonOpenWallet.down ? "#383838" : "#444444"
            }

            onClicked: {
                busyIndicator.running = true
                QmlBridge.clickedButtonOpen(textInputExistingWalletPath.text, textInputExistingWalletPassword.text);
                textInputExistingWalletPassword.text = ""
            }
        }
    }

    // separator horizontal 1
    Rectangle {
        height: 2
        color: "#858585"
        anchors.top: rectangleOpenExistingWallet.bottom
        anchors.topMargin: 10
        anchors.left: rectangleOpenExistingWallet.left
        anchors.leftMargin: 30
        anchors.right: rectangleOpenExistingWallet.right
        anchors.rightMargin: 30
    }

    // section create new wallet
    Rectangle {
        id: rectangleCreateWallet
        color: "transparent"
        anchors.left: rectangleOpenExistingWallet.left
        anchors.leftMargin: 0
        anchors.right: rectangleOpenExistingWallet.right
        anchors.rightMargin: 0
        anchors.top: rectangleOpenExistingWallet.bottom
        anchors.topMargin: 20
        height: 160
        radius: 7

        Text {
            id: textCreateWalletDescr
            color: "#ffffff"
            text: qsTr("Create a new wallet")
            horizontalAlignment: Text.AlignLeft
            verticalAlignment: Text.AlignBottom
            font.bold: true
            font.pixelSize: 16
            font.family: "Arial"
            anchors.top: parent.top
            anchors.topMargin: 10
            anchors.left: parent.left
            anchors.leftMargin: 40
        }

        Text {
            id: textCreateWalletFilenameDescr
            color: "#ffffff"
            text: qsTr("Choose a filename for your new wallet file")
            horizontalAlignment: Text.AlignLeft
            verticalAlignment: Text.AlignBottom
            font.pixelSize: 14
            font.family: "Arial"
            anchors.top: textCreateWalletDescr.bottom
            anchors.topMargin: 20
            anchors.left: parent.left
            anchors.leftMargin: 80
        }

        Rectangle {
            id: rectangleTextInputCreateWalletFilename
            color: "#555555"
            height: 25
            width: 200
            anchors.top: textCreateWalletFilenameDescr.bottom
            anchors.topMargin: 15
            anchors.left: textCreateWalletFilenameDescr.left
            anchors.leftMargin: 0
            radius: 3

            TextInput {
                id: textInputCreateWalletFilename
                anchors.fill: parent
                color: "#cfcfcf"
                text: qsTr("myFirstTRTLWallet")
                rightPadding: 5
                leftPadding: 5
                selectionColor: "#eeeeee"
                selectedTextColor: "#999999"
                selectByMouse: true
                clip: true
                font.family: "Arial"
                horizontalAlignment: Text.AlignLeft
                font.pixelSize: 14
                verticalAlignment: Text.AlignVCenter

                onTextChanged: {
                    buttonCreateWallet.enabled = textInputCreateWalletFilename.text != "" && textInputCreateWalletPassword.text != ""
                }
            }
        }

        Text {
            id: textCreateWalletFilename
            color: "#999999"
            text: textInputCreateWalletFilename.text + ".wallet"
            anchors.verticalCenter: rectangleTextInputCreateWalletFilename.verticalCenter
            anchors.left: rectangleTextInputCreateWalletFilename.right
            anchors.leftMargin: 20
            font.pixelSize: 13
            verticalAlignment: Text.AlignBottom
            font.family: "Arial"
            font.bold: false
            horizontalAlignment: Text.AlignLeft
        }

        Text {
            id: textCreateWalletExtensionDescr
            color: "#999999"
            text: qsTr("Do not include any extension, a \".wallet\" will be added automatically. \nAvoid spaces and most special characters in the filename.")
            anchors.top: rectangleTextInputCreateWalletFilename.bottom
            anchors.topMargin: 12
            anchors.left: rectangleTextInputCreateWalletFilename.left
            anchors.leftMargin: 0
            font.pixelSize: 13
            verticalAlignment: Text.AlignBottom
            font.family: "Arial"
            font.bold: false
            horizontalAlignment: Text.AlignLeft
        }

        Text {
            id: textCreateWalletPasswordDescr
            color: "#ffffff"
            text: qsTr("Choose a strong password")
            horizontalAlignment: Text.AlignLeft
            verticalAlignment: Text.AlignBottom
            font.pixelSize: 14
            font.family: "Arial"
            anchors.top: textCreateWalletFilenameDescr.top
            anchors.topMargin: 0
            anchors.left: parent.left
            anchors.leftMargin: 520
        }

        Rectangle {
            id: rectangleTextInputCreateWalletPassword
            color: "#555555"
            height: 25
            width: 200
            anchors.top: rectangleTextInputCreateWalletFilename.top
            anchors.topMargin: 0
            anchors.left: textCreateWalletPasswordDescr.left
            anchors.leftMargin: 0
            radius: 3

            TextInput {
                id: textInputCreateWalletPassword
                echoMode: TextInput.Password
                anchors.fill: parent
                color: "#cfcfcf"
                text: qsTr("")
                rightPadding: 5
                leftPadding: 5
                selectionColor: "#eeeeee"
                selectedTextColor: "#999999"
                selectByMouse: true
                clip: true
                font.family: "Arial"
                horizontalAlignment: Text.AlignLeft
                font.pixelSize: 14
                verticalAlignment: Text.AlignVCenter

                onTextChanged: {
                    buttonCreateWallet.enabled = textInputCreateWalletFilename.text != "" && textInputCreateWalletPassword.text != ""
                }
            }
        }

        Button {
            id: buttonCreateWallet
            text: qsTr("CREATE")
            anchors.right: parent.right
            anchors.rightMargin: 60
            anchors.bottom: parent.bottom
            anchors.bottomMargin: 15
            enabled: false

            contentItem: Text {
                text: buttonCreateWallet.text
                font.bold: true
                font.pointSize: 20
                font.family: "Arial"
                opacity: enabled ? 1.0 : 0.3
                color: buttonOpenWallet.down ? "#dddddd" : "#ffffff"
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
                elide: Text.ElideRight
            }

            background: Rectangle {
                implicitWidth: 140
                implicitHeight: 60
                opacity: enabled ? 1 : 0.3
                radius: 6
                color: buttonCreateWallet.down ? "#383838" : "#444444"
            }

            onClicked: {
                busyIndicator.running = true
                QmlBridge.clickedButtonCreate(textCreateWalletFilename.text, textInputCreateWalletPassword.text);
                textInputCreateWalletPassword.text = ""
            }
        }
    }

    // separator horizontal 2
    Rectangle {
        height: 2
        color: "#858585"
        anchors.top: rectangleCreateWallet.bottom
        anchors.topMargin: 10
        anchors.left: rectangleOpenExistingWallet.left
        anchors.leftMargin: 30
        anchors.right: rectangleOpenExistingWallet.right
        anchors.rightMargin: 30
    }

    // section import wallet from keys
    Rectangle {
        id: rectangleImportWalletFromKeys
        color: "transparent"
        anchors.left: rectangleOpenExistingWallet.left
        anchors.leftMargin: 0
        anchors.right: rectangleOpenExistingWallet.right
        anchors.rightMargin: 0
        anchors.top: rectangleCreateWallet.bottom
        anchors.topMargin: 20
        anchors.bottom: parent.bottom
        anchors.bottomMargin: 10
        radius: 7

        Text {
            id: textImportWalletDescr
            color: "#ffffff"
            text: qsTr("Import wallet from keys")
            horizontalAlignment: Text.AlignLeft
            verticalAlignment: Text.AlignBottom
            font.bold: true
            font.pixelSize: 16
            font.family: "Arial"
            anchors.top: parent.top
            anchors.topMargin: 10
            anchors.left: parent.left
            anchors.leftMargin: 40
        }

        Text {
            id: textImportWalletFilenameDescr
            color: "#ffffff"
            text: qsTr("Choose a filename for your new wallet file")
            horizontalAlignment: Text.AlignLeft
            verticalAlignment: Text.AlignBottom
            font.pixelSize: 14
            font.family: "Arial"
            anchors.top: textImportWalletDescr.bottom
            anchors.topMargin: 20
            anchors.left: parent.left
            anchors.leftMargin: 80
        }

        Rectangle {
            id: rectangleTextInputImportWalletFilename
            color: "#555555"
            height: 25
            width: 200
            anchors.top: textImportWalletFilenameDescr.bottom
            anchors.topMargin: 15
            anchors.left: textImportWalletFilenameDescr.left
            anchors.leftMargin: 0
            radius: 3

            TextInput {
                id: textInputImportWalletFilename
                anchors.fill: parent
                color: "#cfcfcf"
                text: qsTr("myTRTLWallet")
                rightPadding: 5
                leftPadding: 5
                selectionColor: "#eeeeee"
                selectedTextColor: "#999999"
                selectByMouse: true
                clip: true
                font.family: "Arial"
                horizontalAlignment: Text.AlignLeft
                font.pixelSize: 14
                verticalAlignment: Text.AlignVCenter

                onTextChanged: {
                    rectangleImportWalletFromKeys.checkEnableButton()
                }
            }
        }

        Text {
            id: textImportWalletFilename
            color: "#999999"
            text: textInputImportWalletFilename.text + ".wallet"
            anchors.verticalCenter: rectangleTextInputImportWalletFilename.verticalCenter
            anchors.left: rectangleTextInputImportWalletFilename.right
            anchors.leftMargin: 20
            font.pixelSize: 13
            verticalAlignment: Text.AlignBottom
            font.family: "Arial"
            font.bold: false
            horizontalAlignment: Text.AlignLeft
        }

        Text {
            id: textImportWalletExtensionDescr
            color: "#999999"
            text: qsTr("Do not include any extension, a \".wallet\" will be added automatically. \nAvoid spaces and most special characters in the filename.")
            anchors.top: rectangleTextInputImportWalletFilename.bottom
            anchors.topMargin: 12
            anchors.left: rectangleTextInputImportWalletFilename.left
            anchors.leftMargin: 0
            font.pixelSize: 13
            verticalAlignment: Text.AlignBottom
            font.family: "Arial"
            font.bold: false
            horizontalAlignment: Text.AlignLeft
        }

        Text {
            id: textImportWalletPasswordDescr
            color: "#ffffff"
            text: qsTr("Choose a strong password")
            horizontalAlignment: Text.AlignLeft
            verticalAlignment: Text.AlignBottom
            font.pixelSize: 14
            font.family: "Arial"
            anchors.top: textImportWalletFilenameDescr.top
            anchors.topMargin: 0
            anchors.left: parent.left
            anchors.leftMargin: 520
        }

        Rectangle {
            id: rectangleTextInputImportWalletPassword
            color: "#555555"
            height: 25
            width: 200
            anchors.top: rectangleTextInputImportWalletFilename.top
            anchors.topMargin: 0
            anchors.left: textImportWalletPasswordDescr.left
            anchors.leftMargin: 0
            radius: 3

            TextInput {
                id: textInputImportWalletPassword
                echoMode: TextInput.Password
                anchors.fill: parent
                color: "#cfcfcf"
                text: qsTr("")
                rightPadding: 5
                leftPadding: 5
                selectionColor: "#eeeeee"
                selectedTextColor: "#999999"
                selectByMouse: true
                clip: true
                font.family: "Arial"
                horizontalAlignment: Text.AlignLeft
                font.pixelSize: 14
                verticalAlignment: Text.AlignVCenter

                onTextChanged: {
                    rectangleImportWalletFromKeys.checkEnableButton()
                }
            }
        }

        Text {
            id: textImportWalletPrivateViewKeyDescr
            color: "#ffffff"
            text: qsTr("Private view key")
            horizontalAlignment: Text.AlignLeft
            verticalAlignment: Text.AlignBottom
            font.pixelSize: 14
            font.family: "Arial"
            anchors.top: textImportWalletExtensionDescr.bottom
            anchors.topMargin: 20
            anchors.left: textImportWalletExtensionDescr.left
            anchors.leftMargin: 0
        }

        Rectangle {
            id: rectangleTextInputImportWalletPrivateViewKey
            color: "#555555"
            height: 25
            anchors.top: textImportWalletPrivateViewKeyDescr.bottom
            anchors.topMargin: 15
            anchors.left: textImportWalletPrivateViewKeyDescr.left
            anchors.leftMargin: 0
            anchors.right: rectangleTextInputImportWalletPassword.right
            anchors.rightMargin: 0
            radius: 3

            TextInput {
                id: textInputImportWalletPrivateViewKey
                anchors.fill: parent
                color: "#cfcfcf"
                text: qsTr("")
                rightPadding: 5
                leftPadding: 5
                selectionColor: "#eeeeee"
                selectedTextColor: "#999999"
                selectByMouse: true
                clip: true
                font.family: "Arial"
                horizontalAlignment: Text.AlignLeft
                font.pixelSize: 14
                verticalAlignment: Text.AlignVCenter

                onTextChanged: {
                    rectangleImportWalletFromKeys.checkEnableButton()
                }
            }
        }

        Text {
            id: textImportWalletPrivateSpendKeyDescr
            color: "#ffffff"
            text: qsTr("Private spend key")
            horizontalAlignment: Text.AlignLeft
            verticalAlignment: Text.AlignBottom
            font.pixelSize: 14
            font.family: "Arial"
            anchors.top: rectangleTextInputImportWalletPrivateViewKey.bottom
            anchors.topMargin: 20
            anchors.left: rectangleTextInputImportWalletPrivateViewKey.left
            anchors.leftMargin: 0
        }

        Rectangle {
            id: rectangleTextInputImportWalletPrivateSpendKey
            color: "#555555"
            height: 25
            anchors.top: textImportWalletPrivateSpendKeyDescr.bottom
            anchors.topMargin: 15
            anchors.left: textImportWalletPrivateSpendKeyDescr.left
            anchors.leftMargin: 0
            anchors.right: rectangleTextInputImportWalletPrivateViewKey.right
            anchors.rightMargin: 0
            radius: 3

            TextInput {
                id: textInputImportWalletPrivateSpendKey
                anchors.fill: parent
                color: "#cfcfcf"
                text: qsTr("")
                rightPadding: 5
                leftPadding: 5
                selectionColor: "#eeeeee"
                selectedTextColor: "#999999"
                selectByMouse: true
                clip: true
                font.family: "Arial"
                horizontalAlignment: Text.AlignLeft
                font.pixelSize: 14
                verticalAlignment: Text.AlignVCenter

                onTextChanged: {
                    rectangleImportWalletFromKeys.checkEnableButton()
                }
            }
        }

        Button {
            id: buttonImportWallet
            text: qsTr("IMPORT")
            anchors.right: parent.right
            anchors.rightMargin: 60
            anchors.bottom: parent.bottom
            anchors.bottomMargin: 15
            enabled: false

            contentItem: Text {
                text: buttonImportWallet.text
                font.bold: true
                font.pointSize: 20
                font.family: "Arial"
                opacity: enabled ? 1.0 : 0.3
                color: buttonImportWallet.down ? "#dddddd" : "#ffffff"
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
                elide: Text.ElideRight
            }

            background: Rectangle {
                implicitWidth: 140
                implicitHeight: 60
                opacity: enabled ? 1 : 0.3
                radius: 6
                color: buttonImportWallet.down ? "#383838" : "#444444"
            }

            onClicked: {
                busyIndicator.running = true
                QmlBridge.clickedButtonImport(textImportWalletFilename.text, textInputImportWalletPassword.text, textInputImportWalletPrivateViewKey.text, textInputImportWalletPrivateSpendKey.text);
                textInputImportWalletPassword.text = ""
            }
        }

        function checkEnableButton() {

            buttonImportWallet.enabled = textInputImportWalletFilename.text != "" && textInputImportWalletPassword.text != "" && textInputImportWalletPrivateViewKey.text != "" && textInputImportWalletPrivateSpendKey.text != ""

        }
    }

    FileDialog {
        id: dialogChooseWalletFile
        title: "Please select your wallet file"
        visible: false
        onAccepted: {
            textInputExistingWalletPath.text = dialogChooseWalletFile.fileUrl
        }
        
        function show() {
            dialogChooseWalletFile.open()
        }
    }

    BusyIndicator {
        id: busyIndicator
        anchors.centerIn: parent
        running: false
    }

    Connections {
        target: QmlBridge

        onFinishedLoadingWalletd:
        {
            busyIndicator.running = false
        }

        onFinishedCreatingWallet:
        {
            busyIndicator.running = false
        }

        onDisplayPathToPreviousWallet:
        {
            textInputExistingWalletPath.text = pathToPreviousWallet
        }

    }

    function clearData() {

        textInputImportWalletPrivateViewKey.text = ""
        textInputImportWalletPrivateSpendKey.text = ""

    }

}
