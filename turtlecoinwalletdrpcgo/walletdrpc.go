// Package turtlecoinwalletdrpcgo handles the the rpc connection between your app and walletd
package turtlecoinwalletdrpcgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Transfer contains all the information about a specific transfer
type Transfer struct {
	PaymentID              string
	TxID                   string
	Timestamp              time.Time
	Amount                 float64
	Fee                    float64
	Block                  int
	Confirmations          int
	IsRecievingTransaction bool
}

var (
	rpcURL = "http://127.0.0.1:8070/json_rpc"
)

// RequestBalance provides the available and locked balances of the current wallet
// returned balances are expressed in TRTL, not in 0.01 TRTL
func RequestBalance(rpcPassword string) (availableBalance float64, lockedBalance float64, totalBalance float64, err error) {

	args := make(map[string]interface{})
	payload := rpcPayloadGetBalance(0, rpcPassword, args)

	responseMap, err := httpRequest(payload)

	if err == nil {
		availableBalance = responseMap["result"].(map[string]interface{})["availableBalance"].(float64) / 100
		lockedBalance = responseMap["result"].(map[string]interface{})["lockedAmount"].(float64) / 100
		totalBalance = availableBalance + lockedBalance

		return availableBalance, lockedBalance, totalBalance, nil
	}
	return 0, 0, 0, err
}

// RequestAddress provides the address of the current wallet
func RequestAddress(rpcPassword string) (address string, err error) {

	args := make(map[string]interface{})
	payload := rpcPayloadGetAddresses(0, rpcPassword, args)

	responseMap, err := httpRequest(payload)

	if err == nil {
		walletAddresses := responseMap["result"].(map[string]interface{})["addresses"].([]interface{})
		address = walletAddresses[0].(string)
		return address, nil
	}
	return "", err
}

// RequestListTransactions provides the list of transactions of current wallet
func RequestListTransactions(blockCount int, firstBlockIndex int, addresses []string, rpcPassword string) (transfers []Transfer, err error) {

	args := make(map[string]interface{})
	args["blockCount"] = blockCount
	args["firstBlockIndex"] = firstBlockIndex
	args["addresses"] = addresses
	payload := rpcPayloadGetTransactions(0, rpcPassword, args)

	responseMap, err := httpRequest(payload)

	if err == nil {
		blocks := responseMap["result"].(map[string]interface{})["items"].([]interface{})

		for _, block := range blocks {

			transactions := block.(map[string]interface{})["transactions"].([]interface{})

			for _, transaction := range transactions {

				mapTransaction := transaction.(map[string]interface{})

				var transfer Transfer
				transfer.PaymentID = mapTransaction["paymentId"].(string)
				transfer.TxID = mapTransaction["transactionHash"].(string)
				transfer.Timestamp = time.Unix(int64(mapTransaction["timestamp"].(float64)), 0)
				transfer.Amount = mapTransaction["amount"].(float64) / 100
				transfer.Fee = mapTransaction["fee"].(float64) / 100
				transfer.Block = int(mapTransaction["blockIndex"].(float64))
				transfer.Confirmations = blockCount - transfer.Block + 1
				transfer.IsRecievingTransaction = transfer.Amount >= 0

				transfers = append(transfers, transfer)
			}
		}
		return transfers, nil
	}
	return nil, err
}

// RequestStatus requests walletd connection and sync status
func RequestStatus(rpcPassword string) (blockCount int, knownBlockCount int, peerCount int, err error) {

	args := make(map[string]interface{})
	payload := rpcPayloadGetStatus(0, rpcPassword, args)

	responseMap, err := httpRequest(payload)

	if err == nil {
		blockCount = int(responseMap["result"].(map[string]interface{})["blockCount"].(float64))
		knownBlockCount = int(responseMap["result"].(map[string]interface{})["knownBlockCount"].(float64))
		peerCount = int(responseMap["result"].(map[string]interface{})["peerCount"].(float64))

		return blockCount, knownBlockCount, peerCount, nil
	}
	return 0, 0, 0, err
}

// SendTransaction makes a transfer with the provided information.
// parameters amount and fee are expressed in TRTL, not 0.01 TRTL
func SendTransaction(addressRecipient string, amount float64, paymentID string, fee float64, mixin int, rpcPassword string) (transactionHash string, err error) {

	amountInt := int(amount * 100) // expressed in hundredth of TRTL
	feeInt := int(fee * 100)       // expressed in hundredth of TRTL

	args := make(map[string]interface{})
	args["fee"] = feeInt
	args["paymentId"] = paymentID
	args["anonymity"] = mixin
	var transfers [1]map[string]interface{}
	transfer := make(map[string]interface{})
	transfer["amount"] = amountInt
	transfer["address"] = addressRecipient
	transfers[0] = transfer
	args["transfers"] = transfers

	payload := rpcPayloadSendTransaction(0, rpcPassword, args)

	responseMap, err := httpRequest(payload)

	if err == nil {
		responseError := responseMap["error"]
		if responseError != nil {
			return "", errors.New(responseError.(map[string]interface{})["message"].(string))
		}
		return responseMap["result"].(map[string]interface{})["transactionHash"].(string), nil
	}

	return "", err
}

// GetViewKey provides the private view key
func GetViewKey(rpcPassword string) (privateViewKey string, err error) {

	args := make(map[string]interface{})
	payload := rpcPayloadGetViewKey(0, rpcPassword, args)

	responseMap, err := httpRequest(payload)

	if err == nil {
		privateViewKey = responseMap["result"].(map[string]interface{})["viewSecretKey"].(string)
		return privateViewKey, nil
	}

	return "", err
}

// GetSpendKeys provides the private and public spend keys
func GetSpendKeys(address string, rpcPassword string) (spendSecretKey string, spendPublicKey string, err error) {

	args := make(map[string]interface{})
	args["address"] = address
	payload := rpcPayloadGetSpendKeys(0, rpcPassword, args)

	responseMap, err := httpRequest(payload)

	if err == nil {
		spendSecretKey = responseMap["result"].(map[string]interface{})["spendSecretKey"].(string)
		spendPublicKey = responseMap["result"].(map[string]interface{})["spendSecretKey"].(string)
		return spendSecretKey, spendPublicKey, nil
	}

	return "", "", err
}

// SaveWallet saves the sync info in the wallet
func SaveWallet(rpcPassword string) (err error) {

	args := make(map[string]interface{})
	payload := rpcPayloadSave(0, rpcPassword, args)

	_, err = httpRequest(payload)

	return err
}

func httpRequest(payload rpcPayload) (responseMap map[string]interface{}, err error) {

	payloadjson, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("error json marshal: ", err)
	}

	req, err := http.NewRequest("POST", rpcURL, bytes.NewBuffer(payloadjson))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("error http request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("error reading result from rpc request getSpendKey:", err)
	} else {
		var responseBodyInterface interface{}
		if err := json.Unmarshal(responseBody, &responseBodyInterface); err != nil {
			log.Fatal("JSON unmarshaling with interface failed:", err)
		} else {
			responseMap := responseBodyInterface.(map[string]interface{})

			return responseMap, nil
		}
	}

	return nil, errors.New("unknown error in function httpRequest")
}
