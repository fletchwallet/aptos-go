package aptos

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type AptosClient struct {
	nodeURL string
}

func NewAptosClient(nodeURL string) *AptosClient {
	return &AptosClient{
		nodeURL: nodeURL,
	}
}

func (ac *AptosClient) AccountBalance(address string) (int64, error) {
	res, err := ac.AccountResourceByType(address, "0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>", "")
	if err != nil {
		return -1, err
	}

	v := res.Data["coin"].(map[string]interface{})

	return strconv.ParseInt(v["value"].(string), 10, 0)
}

func (ac *AptosClient) SendCoins(account *AptosAccount, destAddr string, amount int64) (*Transaction, error) {
	payload := &ScriptFunctionPayload{
		Type:          "script_function_payload",
		Function:      "0x1::coin::transfer",
		TypeArguments: []string{"0x1::aptos_coin::AptosCoin"},
		Arguments:     []string{destAddr, strconv.FormatInt(amount, 10)},
	}

	unsignedTx, err := ac.GenerateTransaction(account, payload)
	if err != nil {
		return nil, err
	}

	signedTx, err := ac.SignTransaction(account, unsignedTx)
	if err != nil {
		return nil, err
	}

	return ac.SubmitTransaction(signedTx)
}

func (ac *AptosClient) SendCoinsSync(account *AptosAccount, destAddr string, amount int64) (*Transaction, error) {
	tx, err := ac.SendCoins(account, destAddr, amount)
	if err != nil {
		return nil, err
	}

	err = ac.WaitForTransaction(tx.Hash)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

//TODO: decimal places
func (ac *AptosClient) FundFromFaucet(faucetURL, address string, amount int64) error {
	url := fmt.Sprintf("%s/mint?amount=%d&address=%s", faucetURL, amount, address)
	resp, err := http.DefaultClient.Post(url, "application/json", nil)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(string(body))
	}

	var txnHashes []string
	err = json.Unmarshal(body, &txnHashes)
	if err != nil {
		return err
	}

	return ac.WaitForTransaction(txnHashes[0])
}
