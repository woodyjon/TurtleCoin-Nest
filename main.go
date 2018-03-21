package main

import (
	"TurtleCoin-Nest/walletdmanager"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"time"

	"os"

	"github.com/atotto/clipboard"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"

	log "github.com/sirupsen/logrus"
)

var (
	// qmlObjects = make(map[string]*core.QObject)
	qmlBridge *QmlBridge

	transfers []walletdmanager.Transfer

	tickerRefreshWalletData *time.Ticker

	logFileFilename = "turtlecoin-nest-logs.log"
)

// QmlBridge is the bridge between qml and go
type QmlBridge struct {
	core.QObject

	_ func(data string) `signal:"displayTotalBalance"`
	_ func(data string) `signal:"displayAvailableBalance"`
	_ func(data string) `signal:"displayLockedBalance"`
	_ func(data string) `signal:"displayAddress"`
	_ func(paymentID string,
		transactionID string,
		amount string,
		confirmations string,
		time string,
		number string) `signal:"addTransactionToList"`
	_ func(text string,
		time int) `signal:"displayPopup"`
	_ func(syncing string,
		blocks string,
		peers string) `signal:"displaySyncingInfo"`
	_ func(errorMessage string) `signal:"displayErrorDialog"`
	_ func()                    `signal:"clearTransferAmount"`
	_ func()                    `signal:"clearListTransactions"`
	_ func(filename string,
		privateViewKey string,
		privateSpendKey string,
		walletAddress string) `signal:"displayPrivateKeys"`
	_ func() `signal:"displayOpenWalletScreen"`
	_ func() `signal:"displayMainWalletScreen"`
	_ func() `signal:"finishedLoadingWalletd"`
	_ func() `signal:"finishedCreatingWallet"`

	_ func(msg string)           `slot:"log"`
	_ func(transactionID string) `slot:"clickedButtonExplorer"`
	_ func(transactionID string) `slot:"clickedButtonCopyTx"`
	_ func()                     `slot:"clickedButtonCopyAddress"`
	_ func(transferAddress string,
		transferAmount string,
		transferPaymentID string) `slot:"clickedButtonSend"`
	_ func() `slot:"clickedButtonBackupWallet"`
	_ func() `slot:"clickedOpenAnotherWallet"`
	_ func(pathToWallet string,
		passwordWallet string) `slot:"clickedButtonOpen"`
	_ func(filenameWallet string,
		passwordWallet string) `slot:"clickedButtonCreate"`
	_ func(filenameWallet string,
		passwordWallet string,
		privateViewKey string,
		privateSpendKey string) `slot:"clickedButtonImport"`

	_ func(object *core.QObject) `slot:"registerToGo"`
	_ func(objectName string)    `slot:"deregisterToGo"`
}

func main() {

	currentDirectory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("error finding current directory. Error: ", err)
	}
	pathToLogFile := filepath.Dir(filepath.Dir(filepath.Dir(currentDirectory)))
	pathToLogFile = pathToLogFile + "/" + logFileFilename

	logFile, err := os.OpenFile(pathToLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("error opening log file: ", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	log.SetLevel(log.DebugLevel)

	log.Info("Application started")

	gui.NewQGuiApplication(len(os.Args), os.Args)

	// Enable high DPI scaling
	// app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Use the material style for qml
	quickcontrols2.QQuickStyle_SetStyle("material")

	engine := qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/qml/nestmain.qml", 0))

	qmlBridge = NewQmlBridge(nil)

	qmlBridge.ConnectLog(func(msg string) {

		log.Info("QML: ", msg)

	})

	qmlBridge.ConnectClickedButtonCopyAddress(func() {

		clipboard.WriteAll(walletdmanager.WalletAddress)

		qmlBridge.DisplayPopup("Copied!", 1500)

	})

	qmlBridge.ConnectClickedButtonCopyTx(func(transactionID string) {

		clipboard.WriteAll(transactionID)

		qmlBridge.DisplayPopup("Copied!", 1500)

	})

	qmlBridge.ConnectClickedButtonExplorer(func(transactionID string) {

		url := "https://turtle-coin.com/?hash=" + transactionID + "#blockchain_transaction"
		successOpenBrowser := openBrowser(url)

		if !successOpenBrowser {
			log.Error("failure opening browser, url: " + url)
		}

	})

	qmlBridge.ConnectClickedButtonSend(func(transferAddress string, transferAmount string, transferPaymentID string) {

		transfer(transferAddress, transferAmount, transferPaymentID)

	})

	qmlBridge.ConnectClickedButtonBackupWallet(func() {

		showWalletPrivateInfo()

	})

	qmlBridge.ConnectClickedButtonOpen(func(pathToWallet string, passwordWallet string) {

		go func() {

			startWalletWithWalletInfo(pathToWallet, passwordWallet)

		}()

	})

	qmlBridge.ConnectClickedButtonCreate(func(filenameWallet string, passwordWallet string) {

		go func() {

			createWalletWithWalletInfo(filenameWallet, passwordWallet)

		}()

	})

	qmlBridge.ConnectClickedButtonImport(func(filenameWallet string, passwordWallet string, privateViewKey string, privateSpendKey string) {

		go func() {

			importWalletWithWalletInfo(filenameWallet, passwordWallet, privateViewKey, privateSpendKey)

		}()

	})

	qmlBridge.ConnectClickedOpenAnotherWallet(func() {

		openAnotherWallet()

	})

	engine.RootContext().SetContextProperty("QmlBridge", qmlBridge)

	gui.QGuiApplication_Exec()

	log.Info("Application closed")

	walletdmanager.StopWalletd()

}

func startDisplayWalletInfo() {

	getAndDisplayBalances()

	getAndDisplayAddress()

	getAndDisplayListTransactions()

	getAndDisplayConnectionInfo()

	go func() {

		tickerRefreshWalletData = time.NewTicker(time.Second * 15)

		for range tickerRefreshWalletData.C {

			getAndDisplayBalances()

			getAndDisplayListTransactions()

			getAndDisplayConnectionInfo()

		}

	}()

}

func getAndDisplayBalances() {

	walletAvailableBalance, walletLockedBalance, walletTotalBalance := walletdmanager.RequestBalance()

	qmlBridge.DisplayAvailableBalance(strconv.FormatFloat(walletAvailableBalance, 'f', -1, 64))
	qmlBridge.DisplayLockedBalance(strconv.FormatFloat(walletLockedBalance, 'f', -1, 64))
	qmlBridge.DisplayTotalBalance(strconv.FormatFloat(walletTotalBalance, 'f', -1, 64))

}

func getAndDisplayAddress() {

	walletAddress := walletdmanager.RequestAddress()

	qmlBridge.DisplayAddress(walletAddress)

}

func getAndDisplayConnectionInfo() {

	syncing, blockCount, knownBlockCount, peers, err := walletdmanager.RequestConnectionInfo()

	if err == nil {

		blocks := blockCount + " / " + knownBlockCount

		qmlBridge.DisplaySyncingInfo(syncing, blocks, peers)

	}

}

func getAndDisplayListTransactions() {

	newTransfers := walletdmanager.RequestListTransactions()

	if len(newTransfers) != len(transfers) {

		transfers = newTransfers

		// sort starting by the most recent transaction
		sort.Slice(transfers, func(i, j int) bool { return transfers[i].Timestamp.After(transfers[j].Timestamp) })

		transactionNumber := len(transfers)

		qmlBridge.ClearListTransactions()

		for _, transfer := range transfers {

			amount := transfer.Amount
			amountString := ""
			if amount >= 0 {
				amountString += "+ "
				amountString += strconv.FormatFloat(amount, 'f', -1, 64)
			} else {
				amountString += "- "
				amountString += strconv.FormatFloat(-amount, 'f', -1, 64)
			}
			amountString += " TRTL (fee: " + strconv.FormatFloat(transfer.Fee, 'f', 2, 64) + ")"

			confirmationsString := "(" + strconv.Itoa(transfer.Confirmations) + " conf.)"

			timeString := transfer.Timestamp.Format("2006-01-02 15:04:05")

			transactionNumberString := strconv.Itoa(transactionNumber) + ")"
			transactionNumber--

			qmlBridge.AddTransactionToList(transfer.PaymentID, transfer.TxID, amountString, confirmationsString, timeString, transactionNumberString)

		}

	}

}

func transfer(transferAddress string, transferAmount string, transferPaymentID string) bool {

	log.Info("SEND: ", transferAddress, transferAmount, transferPaymentID)

	transactionID, err := walletdmanager.SendTransaction(transferAddress, transferAmount, transferPaymentID)

	if err != nil {

		log.Warn("error transfer: ", err)

		qmlBridge.DisplayErrorDialog(err.Error())

		return false

	}

	log.Info("succes transfer: ", transactionID)

	getAndDisplayBalances()

	qmlBridge.ClearTransferAmount()

	qmlBridge.DisplayPopup("TRTLs sent successfully", 4000)

	return true

}

func startWalletWithWalletInfo(pathToWallet string, passwordWallet string) bool {

	err := walletdmanager.StartWalletd(pathToWallet, passwordWallet)

	if err != nil {

		log.Warn("error starting walletd with provided wallet info. error: ", err)

		qmlBridge.FinishedLoadingWalletd()

		qmlBridge.DisplayErrorDialog(err.Error())

		return false

	}

	log.Info("success starting walletd")

	qmlBridge.FinishedLoadingWalletd()

	startDisplayWalletInfo()

	qmlBridge.DisplayMainWalletScreen()

	return true

}

func createWalletWithWalletInfo(filenameWallet string, passwordWallet string) bool {

	err := walletdmanager.CreateWallet(filenameWallet, passwordWallet, "", "")

	if err != nil {

		log.Warn("error creating wallet. error: ", err)

		qmlBridge.FinishedCreatingWallet()

		qmlBridge.DisplayErrorDialog(err.Error())

		return false

	}

	log.Info("success creating wallet")

	startWalletWithWalletInfo(filenameWallet, passwordWallet)

	showWalletPrivateInfo()

	return true

}

func importWalletWithWalletInfo(filenameWallet string, passwordWallet string, privateViewKey string, privateSpendKey string) bool {

	err := walletdmanager.CreateWallet(filenameWallet, passwordWallet, privateViewKey, privateSpendKey)

	if err != nil {

		log.Warn("error importing wallet. error: ", err)

		qmlBridge.FinishedCreatingWallet()

		qmlBridge.DisplayErrorDialog(err.Error())

		return false

	}

	log.Info("success importing wallet")

	startWalletWithWalletInfo(filenameWallet, passwordWallet)

	return true

}

func openAnotherWallet() {

	tickerRefreshWalletData.Stop()

	walletdmanager.StopWalletd()

	qmlBridge.DisplayOpenWalletScreen()

}

func showWalletPrivateInfo() {

	privateViewKey, privateSpendKey, err := walletdmanager.GetPrivateViewKeyAndSpendKey()

	if err != nil {

		log.Error("Error getting view and spend key: ", err)

	} else {

		qmlBridge.DisplayPrivateKeys(walletdmanager.WalletFilename, privateViewKey, privateSpendKey, walletdmanager.WalletAddress)

	}

}

func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
