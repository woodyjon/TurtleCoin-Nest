package main

import (
	"TurtleCoin-Nest/turtlecoinwalletdrpcgo"
	"TurtleCoin-Nest/walletdmanager"
	"database/sql"
	"io"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"sort"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

var (
	// qmlObjects = make(map[string]*core.QObject)
	qmlBridge               *QmlBridge
	transfers               []turtlecoinwalletdrpcgo.Transfer
	tickerRefreshWalletData *time.Ticker
	logFileFilename         = "turtlecoin-nest-logs.log"
	urlBlockExplorer        = "https://blocks.turtle.link/"
	db                      *sql.DB
	dbFilename              = "settings.db"
	useRemoteNode           = true
	stringBackupKeys        = ""
)

// QmlBridge is the bridge between qml and go
type QmlBridge struct {
	core.QObject

	// go to qml
	_ func(data string) `signal:"displayTotalBalance"`
	_ func(data string) `signal:"displayAvailableBalance"`
	_ func(data string) `signal:"displayLockedBalance"`
	_ func(address string,
		wallet string) `signal:"displayAddress"`
	_ func(paymentID string,
		transactionID string,
		amount string,
		confirmations string,
		time string,
		number string) `signal:"addTransactionToList"`
	_ func(text string, time int)                       `signal:"displayPopup"`
	_ func(syncing string, blocks string, peers string) `signal:"displaySyncingInfo"`
	_ func(errorMessage string)                         `signal:"displayErrorDialog"`
	_ func()                                            `signal:"clearTransferAmount"`
	_ func()                                            `signal:"clearListTransactions"`
	_ func(filename string,
		privateViewKey string,
		privateSpendKey string,
		walletAddress string) `signal:"displayPrivateKeys"`
	_ func()                            `signal:"displayOpenWalletScreen"`
	_ func()                            `signal:"displayMainWalletScreen"`
	_ func()                            `signal:"finishedLoadingWalletd"`
	_ func()                            `signal:"finishedCreatingWallet"`
	_ func(pathToPreviousWallet string) `signal:"displayPathToPreviousWallet"`
	_ func(walletLocation string)       `signal:"displayWalletCreationLocation"`
	_ func(useRemote bool)              `signal:"displayUseRemoteNode"`

	// qml to go
	_ func(msg string)           `slot:"log"`
	_ func(transactionID string) `slot:"clickedButtonExplorer"`
	_ func(transactionID string) `slot:"clickedButtonCopyTx"`
	_ func()                     `slot:"clickedButtonCopyAddress"`
	_ func()                     `slot:"clickedButtonCopyKeys"`
	_ func(transferAddress string,
		transferAmount string,
		transferPaymentID string) `slot:"clickedButtonSend"`
	_ func()                                             `slot:"clickedButtonBackupWallet"`
	_ func()                                             `slot:"clickedCloseWallet"`
	_ func(pathToWallet string, passwordWallet string)   `slot:"clickedButtonOpen"`
	_ func(filenameWallet string, passwordWallet string) `slot:"clickedButtonCreate"`
	_ func(filenameWallet string,
		passwordWallet string,
		privateViewKey string,
		privateSpendKey string) `slot:"clickedButtonImport"`
	_ func(remote bool) `slot:"choseRemote"`

	_ func(object *core.QObject) `slot:"registerToGo"`
	_ func(objectName string)    `slot:"deregisterToGo"`
}

func main() {

	pathToLogFile := logFileFilename
	pathToDB := dbFilename
	pathToHomeDir := ""

	if isPlatformDarwin {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		pathToHomeDir = usr.HomeDir
		pathToAppFolder := pathToHomeDir + "/Library/Application Support/TurtleCoin-Nest"
		os.Mkdir(pathToAppFolder, os.ModePerm)
		pathToLogFile = pathToAppFolder + "/" + logFileFilename
		pathToDB = pathToAppFolder + "/" + pathToDB
	}

	logFile, err := os.OpenFile(pathToLogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal("error opening log file: ", err)
	}
	defer logFile.Close()

	if isPlatformLinux {
		// log to file and console
		mw := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(mw)
	} else {
		log.SetOutput(logFile)
	}

	log.SetLevel(log.DebugLevel)

	setupDB(pathToDB)

	log.Info("Application started")

	platform := "linux"
	if isPlatformDarwin {
		platform = "darwin"
	} else if isPlatformWindows {
		platform = "windows"
	}
	walletdmanager.Setup(platform)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	// Enable high DPI scaling
	// app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Use the material style for qml
	quickcontrols2.QQuickStyle_SetStyle("material")

	engine := qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/qml/nestmain.qml", 0))

	qmlBridge = NewQmlBridge(nil)

	connectQMLToGOFunctions()

	engine.RootContext().SetContextProperty("QmlBridge", qmlBridge)

	if isPlatformDarwin {
		textLocation := "Your wallet will be saved in your home directory: " + pathToHomeDir + "/"
		qmlBridge.DisplayWalletCreationLocation(textLocation)
	}

	getAndDisplayPathWalletFromDB()
	getAndDisplayUseRemoteFromDB()

	gui.QGuiApplication_Exec()

	log.Info("Application closed")

	walletdmanager.GracefullyQuitWalletd()
}

func connectQMLToGOFunctions() {

	qmlBridge.ConnectLog(func(msg string) {
		log.Info("QML: ", msg)
	})

	qmlBridge.ConnectClickedButtonCopyAddress(func() {
		clipboard.WriteAll(walletdmanager.WalletAddress)
		qmlBridge.DisplayPopup("Copied!", 1500)
	})

	qmlBridge.ConnectClickedButtonCopyKeys(func() {
		clipboard.WriteAll(stringBackupKeys)
	})

	qmlBridge.ConnectClickedButtonCopyTx(func(transactionID string) {
		clipboard.WriteAll(transactionID)
		qmlBridge.DisplayPopup("Copied!", 1500)
	})

	qmlBridge.ConnectClickedButtonExplorer(func(transactionID string) {
		url := urlBlockExplorer + "?hash=" + transactionID + "#blockchain_transaction"
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
			recordPathWalletToDB(pathToWallet)
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

	qmlBridge.ConnectClickedCloseWallet(func() {
		closeWallet()
	})

	qmlBridge.ConnectChoseRemote(func(remote bool) {
		useRemoteNode = remote
		recordUseRemoteToDB(useRemoteNode)
	})
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

	walletAvailableBalance, walletLockedBalance, walletTotalBalance, err := walletdmanager.RequestBalance()
	if err == nil {
		qmlBridge.DisplayAvailableBalance(strconv.FormatFloat(walletAvailableBalance, 'f', -1, 64))
		qmlBridge.DisplayLockedBalance(strconv.FormatFloat(walletLockedBalance, 'f', -1, 64))
		qmlBridge.DisplayTotalBalance(strconv.FormatFloat(walletTotalBalance, 'f', -1, 64))
	}
}

func getAndDisplayAddress() {

	walletAddress, err := walletdmanager.RequestAddress()
	if err == nil {
		qmlBridge.DisplayAddress(walletAddress, walletdmanager.WalletFilename)
	}
}

func getAndDisplayConnectionInfo() {

	syncing, blockCount, knownBlockCount, peers, err := walletdmanager.RequestConnectionInfo()
	if err == nil {
		blocks := blockCount + " / " + knownBlockCount
		qmlBridge.DisplaySyncingInfo(syncing, blocks, peers)
	}
}

func getAndDisplayListTransactions() {

	newTransfers, err := walletdmanager.RequestListTransactions()
	if err == nil {
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

	err := walletdmanager.StartWalletd(pathToWallet, passwordWallet, useRemoteNode)
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

func closeWallet() {

	tickerRefreshWalletData.Stop()

	stringBackupKeys = ""

	go func() {
		walletdmanager.GracefullyQuitWalletd()
	}()

	qmlBridge.DisplayOpenWalletScreen()
}

func showWalletPrivateInfo() {

	privateViewKey, privateSpendKey, err := walletdmanager.GetPrivateViewKeyAndSpendKey()
	if err != nil {
		log.Error("Error getting view and spend key: ", err)
	} else {
		qmlBridge.DisplayPrivateKeys(walletdmanager.WalletFilename, privateViewKey, privateSpendKey, walletdmanager.WalletAddress)

		stringBackupKeys = "Wallet: " + walletdmanager.WalletFilename + "\nAddress: " + walletdmanager.WalletAddress + "\nPrivate view key: " + privateViewKey + "\nPrivate spend key: " + privateSpendKey
	}
}

func setupDB(pathToDB string) {

	var err error
	db, err = sql.Open("sqlite3", pathToDB)
	if err != nil {
		log.Fatal("error opening db file. err: ", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS pathWallet (id INTEGER PRIMARY KEY AUTOINCREMENT,path VARCHAR(64) NULL)")
	if err != nil {
		log.Fatal("error creating table pathWallet. err: ", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS remoteNode (id INTEGER PRIMARY KEY AUTOINCREMENT,useRemote BOOL NOT NULL DEFAULT '1', address VARCHAR(64),port INT)")
	if err != nil {
		log.Fatal("error creating table remoteNode. err: ", err)
	}
}

func getAndDisplayPathWalletFromDB() {

	qmlBridge.DisplayPathToPreviousWallet(getPathWalletFromDB())
}

func getPathWalletFromDB() string {

	pathToPreviousWallet := ""

	rows, err := db.Query("SELECT path FROM pathWallet ORDER BY id DESC LIMIT 1")
	if err != nil {
		log.Fatal("error reading path from pathwallet table. err: ", err)
	}

	for rows.Next() {
		path := ""
		err = rows.Scan(&path)
		if err != nil {
			log.Fatal("error reading item from pathWallet table. err: ", err)
		}
		pathToPreviousWallet = path
	}

	return pathToPreviousWallet
}

func recordPathWalletToDB(path string) {

	stmt, err := db.Prepare(`INSERT INTO pathWallet(path) VALUES(?)`)
	if err != nil {
		log.Fatal("error preparing to insert pathWallet into db. err: ", err)
	}
	_, err = stmt.Exec(path)
	if err != nil {
		log.Fatal("error inserting pathWallet into db. err: ", err)
	}
}

func getAndDisplayUseRemoteFromDB() {

	qmlBridge.DisplayUseRemoteNode(getUseRemoteFromDB())
}

func getUseRemoteFromDB() bool {

	rows, err := db.Query("SELECT useRemote FROM remoteNode ORDER BY id DESC LIMIT 1")
	if err != nil {
		log.Fatal("error reading path from remoteNode table. err: ", err)
	}

	for rows.Next() {
		useRemote := true
		err = rows.Scan(&useRemote)
		if err != nil {
			log.Fatal("error reading item from remoteNode table. err: ", err)
		}
		useRemoteNode = useRemote
	}

	return useRemoteNode
}

func recordUseRemoteToDB(useRemote bool) {

	stmt, err := db.Prepare(`INSERT INTO remoteNode(useRemote) VALUES(?)`)
	if err != nil {
		log.Fatal("error preparing to insert useRemoteNode into db. err: ", err)
	}
	_, err = stmt.Exec(useRemote)
	if err != nil {
		log.Fatal("error inserting useRemoteNode into db. err: ", err)
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
