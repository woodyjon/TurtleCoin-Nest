import QtQuick 2.7
import QtQuick.Controls 2.1

Item {
    id: itemListTransaction
    visible: true
    property bool newTransaction: false
    Text {
        id: transactionAmount
        color: "#cfcfcf"
        text: transactionAmountValue
        verticalAlignment: Text.AlignVCenter
        anchors.left: parent.left
        anchors.leftMargin: 35
        anchors.top: parent.top
        anchors.topMargin: 5
        font.family: "Arial"
        font.bold: true
        font.pixelSize: 17
    }
    
    Text {
        id: transactionID
        color: "#cfcfcf"
        text: "Tr. ID: " + transactionIDValue
        anchors.top: transactionAmount.bottom
        anchors.topMargin: 8
        anchors.left: transactionAmount.left
        anchors.leftMargin: 0
        verticalAlignment: Text.AlignVCenter
        font.family: "Arial"
        font.pixelSize: 11
        font.bold: false
    }
    
    Text {
        id: transactionPaymentID
        color: "#cfcfcf"
        text: "Payment ID: " + transactionPaymentIDValue
        anchors.top: transactionID.bottom
        anchors.topMargin: 8
        anchors.left: transactionID.left
        anchors.leftMargin: 0
        verticalAlignment: Text.AlignVCenter
        font.family: "Arial"
        font.pixelSize: 11
        font.bold: false
    }
    
    Text {
        id: transactionTime
        x: 196
        y: 8
        color: "#cfcfcf"
        text: transactionTimeValue
        anchors.right: transactionConfirmations.left
        anchors.rightMargin: 15
        anchors.verticalCenter: transactionAmount.verticalCenter
        horizontalAlignment: Text.AlignRight
        verticalAlignment: Text.AlignVCenter
        font.pixelSize: 12
        font.bold: false
        font.family: "Arial"
    }
    
    Text {
        id: transactionConfirmations
        x: 310
        y: 8
        color: "#cfcfcf"
        text: transactionConfirmationsValue
        anchors.right: parent.right
        anchors.rightMargin: 5
        anchors.verticalCenter: transactionTime.verticalCenter
        horizontalAlignment: Text.AlignRight
        verticalAlignment: Text.AlignVCenter
        font.pixelSize: 12
        font.bold: false
        font.family: "Arial"
    }

    Text {
        id: transactionNumber
        color: "#cfcfcf"
        text: transactionNumberValue
        anchors.top: transactionAmount.top
        anchors.topMargin: 0
        anchors.left: parent.left
        anchors.leftMargin: 5
        font.bold: false
        font.family: "Arial"
        font.pixelSize: 11
        verticalAlignment: Text.AlignVCenter
    }

    Button {
        id: buttonCopyTxID
        width: 13
        height: 14
        hoverEnabled: true
        anchors.verticalCenter: transactionID.verticalCenter
        anchors.left: transactionID.right
        anchors.leftMargin: 10
        Image {
            id: imageButtonCopyTxID
            sourceSize.height: 13
            sourceSize.width: 14
            fillMode: Image.PreserveAspectFit
            source: "images/copy_white.png"
        }
        onClicked: {
            QmlBridge.clickedButtonCopyTx(transactionIDValue);
            popupText.show("Copied!")
        }
    }

    Button {
        id: buttonExplorerTxID
        width: 13
        height: 14
        hoverEnabled: true
        anchors.verticalCenter: buttonCopyTxID.verticalCenter
        anchors.left: buttonCopyTxID.right
        anchors.leftMargin: 10
        Image {
            id: imageButtonExplorerTxID
            sourceSize.height: 13
            sourceSize.width: 14
            fillMode: Image.PreserveAspectFit
            source: "images/search_white.png"
        }
        onClicked: QmlBridge.clickedButtonExplorer(transactionIDValue)
    }
}
