import QtQuick.Window 2.2
import QtQuick 2.7
import QtQuick.Controls 2.3
import QtQuick.Controls 1.4 as OldControls
import QtQuick.Controls.Styles 1.4
import QtQuick.Dialogs 1.2
import QtQuick.Layouts 1.1

Rectangle {
    id: rectangleSettings
    y: parent.height // starts off "screen"
    anchors.horizontalCenter: parent.horizontalCenter
    width: parent.width
    height: parent.height
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

    OldControls.CheckBox {
        id: checkBoxUSD
        text: "Display values also in USD"
        checked: false
        anchors.top: parent.top
        anchors.topMargin: 180
        anchors.left: parent.left
        anchors.leftMargin: 180

        style: CheckBoxStyle {
            label: Text {
                color: "#ffffff"
                font.pixelSize: 17
                font.family: "Arial"
                text: control.text
                leftPadding: 15
            }
        }

        onClicked: {
            QmlBridge.choseDisplayFiat(checkBoxUSD.checked);
        }
    }

    Text {
        id: textSettingsUSDDescr
        color: "#ffffff"
        text: "(Exchange rate from cryptocompare.com.)\n(Consider cautiously as current volume and liquidity are extremely low.)"
        anchors.verticalCenterOffset: 0
        anchors.verticalCenter: checkBoxUSD.verticalCenter
        anchors.left: checkBoxUSD.right
        anchors.leftMargin: 23
        horizontalAlignment: Text.AlignLeft
        font.family: "Arial"
        font.pixelSize: 14
    }

    function show() {
        settingsScreen.state = "visible"
    }

    function hide() {
        settingsScreen.state = ""
    }

    Connections {
        target: QmlBridge

        onDisplaySettingsValues: {
            checkBoxUSD.checked = displayFiat
        }
    }

    states: State {
        name: "visible"
        PropertyChanges { target: settingsScreen; y: 0 }
    }

    transitions: [
        Transition { from: ""; PropertyAnimation { properties: "y"; duration: 200 } },
        Transition { from: "visible"; PropertyAnimation { properties: "y"; duration: 200 } }
    ]
}