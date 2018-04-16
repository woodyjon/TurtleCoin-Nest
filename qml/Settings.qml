import QtQuick.Window 2.2
import QtQuick 2.7
import QtQuick.Controls 2.3
import QtQuick.Controls 1.4 as OldControls
import QtQuick.Controls.Styles 1.4
import QtQuick.Dialogs 1.2
import QtQuick.Layouts 1.1

Rectangle {
    id: rectangleSettings
    anchors.fill: parent
    color: "#333333"

    Rectangle {
        id: buttonClose
        width: 33
        height: 33
        anchors.top: parent.top
        anchors.topMargin: 15
        anchors.left: parent.left
        anchors.leftMargin: 15
        color: "transparent"
        Image {
            id: image_close
            anchors.fill: parent
            fillMode: Image.PreserveAspectFit
            source: "images/close.png"
        }
        MouseArea {
            anchors.fill: parent
            onClicked: {
                QmlBridge.clickedCloseSettings()
            }
        }
    }

    Text {
        id: textSettings
        color: "#ffffff"
        text: "SETTINGS"
        anchors.horizontalCenterOffset: 0
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.top: parent.top
        anchors.topMargin: 50
        horizontalAlignment: Text.AlignHCenter
        font.bold: true
        font.family: "Arial"
        font.pixelSize: 25
    }
}