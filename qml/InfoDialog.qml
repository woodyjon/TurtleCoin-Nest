import QtQuick.Window 2.2
import QtQuick 2.7
import QtQuick.Controls 2.3
import QtQuick.Controls.Styles 1.4
import QtQuick.Dialogs 1.2
import QtQuick.Layouts 1.1

Dialog {
    id: dialogInfo
    title: "Info"
    standardButtons: StandardButton.Ok
    width: 900

    property var addressDev: "TRTLv3jzutiQwqHL3qFwsu5EVLWesxZr1AFQ4AuMR3SD56n3rkHDkwj79eKwvaiU1nYQWGydKoXM6fXyiiGKsPDnVCNXzNdusxx"
    property var websiteChat: "http://chat.turtlecoin.lol"
    property var versionNest: ""

    Text {
        id: textNest
        text: "Nest"
        font.family: "Arial"
        font.pixelSize: 21
        font.bold: true
        anchors.horizontalCenter: parent.horizontalCenter
    }

    Text {
        id: textVersion
        text: "v. " + versionNest
        font.family: "Arial"
        font.pixelSize: 13
        anchors.horizontalCenter: textNest.horizontalCenter
        anchors.top: textNest.bottom
        anchors.topMargin: 5
    }

    Text {
        id: textDescriptionHelp
        text: "Need help?"
        font.family: "Arial"
        font.pixelSize: 13
        anchors.top: textVersion.bottom
        anchors.topMargin: 35
    }

    Text {
        id: textHelp
        anchors.top: textDescriptionHelp.bottom
        anchors.topMargin: 10
        anchors.left: textDescriptionHelp.left
        anchors.leftMargin: 0
        text: "Ask the community on " + websiteChat
        leftPadding: 5
        font.family: "Arial"
        font.pixelSize: 13
        font.bold: true
    }

    Rectangle {
        id: buttonGoToChat
        width: 25
        height: 28
        anchors.verticalCenter: textHelp.verticalCenter
        anchors.left: textHelp.right
        anchors.leftMargin: 15
        color: "transparent"
        Image {
            anchors.fill: parent
            fillMode: Image.PreserveAspectFit
            source: "images/openLink_dark_grey.svg"
        }
        MouseArea {
            anchors.fill: parent
            onClicked: {
                QmlBridge.goToWebsite(websiteChat);
            }
        }
    }

    Text {
        id: textDescriptionDonate
        text: "Donate to support Nest's dev"
        font.family: "Arial"
        font.pixelSize: 13
        anchors.top: textHelp.bottom
        anchors.topMargin: 25
        anchors.left: textHelp.left
        anchors.leftMargin: 0
    }

    TextInput {
        id: textInputAddress
        anchors.top: textDescriptionDonate.bottom
        anchors.topMargin: 10
        anchors.left: textDescriptionDonate.left
        anchors.leftMargin: 0
        readOnly: true
        text: addressDev
        leftPadding: 5
        padding: 2
        selectionColor: "#333333"
        selectedTextColor: "white"
        selectByMouse: true
        font.family: "Arial"
        font.pixelSize: 13
        font.bold: true
        horizontalAlignment: Text.AlignLeft
        verticalAlignment: Text.AlignVCenter
    }

    Rectangle {
        id: buttonCopyAddressDev
        width: 25
        height: 28
        anchors.verticalCenter: textInputAddress.verticalCenter
        anchors.left: textInputAddress.right
        anchors.leftMargin: 15
        color: "transparent"
        Image {
            anchors.fill: parent
            fillMode: Image.PreserveAspectFit
            source: "images/copy_dark_grey.svg"
        }
        MouseArea {
            anchors.fill: parent
            onClicked: {
                QmlBridge.clickedButtonCopy(addressDev);
            }
        }
    }

    function show() {
        versionNest = QmlBridge.getVersion();
        dialogInfo.open();
    }
}