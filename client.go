package aptos

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (ac *AptosClient) makeRequest(method, path string, result interface{}) error {
	// TODO: check path
	fullRoute := ac.nodeURL + path
	req, err := http.NewRequest(method, fullRoute, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New(string(body))
	}

	// fmt.Println(string(body))
	json.Unmarshal(body, result)
	if err != nil {
		return err
	}

	return nil
}

func (ac *AptosClient) LedgerInfo() (*LedgerInfo, error) {
	var info LedgerInfo
	err := ac.makeRequest("GET", "/", &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (ac *AptosClient) Account(address string) (*Account, error) {
	var account Account
	err := ac.makeRequest("GET", fmt.Sprint("/accounts/", address), &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (ac *AptosClient) AccountResources(address, version string) ([]AccountResource, error) {

	path := fmt.Sprint("/accounts/", address, "/resources")
	if version != "" {
		path += fmt.Sprint("?version=", version)
	}

	var accountResources []AccountResource
	err := ac.makeRequest("GET", path, &accountResources)
	if err != nil {
		return nil, err
	}

	return accountResources, nil
}

func (ac *AptosClient) AccountResourceByType(address, resourceType, version string) (*AccountResource, error) {
	path := fmt.Sprint("/accounts/", address, "/resource/", resourceType)
	if version != "" {
		path += fmt.Sprint("?version=", version)
	}

	var accountResource AccountResource
	err := ac.makeRequest("GET", path, &accountResource)
	if err != nil {
		return nil, err
	}

	return &accountResource, nil
}

//TODO: test function
func (ac *AptosClient) AccountModules(address, version string) ([]AccountModule, error) {

	path := fmt.Sprint("/accounts/", address, "/modules")
	if version != "" {
		path += fmt.Sprint("?version=", version)
	}

	var accountModules []AccountModule
	err := ac.makeRequest("GET", path, &accountModules)
	if err != nil {
		return nil, err
	}

	return accountModules, nil
}

//TODO: test function
func (ac *AptosClient) AccountModuleByID(address, moduleID, version string) (*AccountModule, error) {

	path := fmt.Sprint("/accounts/", address, "/module/", moduleID)
	if version != "" {
		path += fmt.Sprint("?version=", version)
	}

	var accountModule AccountModule
	err := ac.makeRequest("GET", path, &accountModule)
	if err != nil {
		return nil, err
	}

	return &accountModule, nil
}

func (ac *AptosClient) GetTransactions(limit, start int) ([]Transaction, error) {
	path := "/transactions"
	path += fmt.Sprint("?limit=", limit)
	path += fmt.Sprint("&start=", start)

	fmt.Println(path)

	var transactions []Transaction
	err := ac.makeRequest("GET", path, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, err
}

func (ac *AptosClient) GetAccountTransactions(address string, limit, start int) ([]Transaction, error) {
	path := fmt.Sprint("/accounts/", address, "/transactions")
	path += fmt.Sprint("?limit=", limit)
	path += fmt.Sprint("&start=", start)

	fmt.Println(path)

	var transactions []Transaction
	err := ac.makeRequest("GET", path, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, err
}

func (ac *AptosClient) GetTransaction(hashOrVersion string) (*Transaction, error) {

	path := fmt.Sprint("/transactions/", hashOrVersion)

	var transaction Transaction
	err := ac.makeRequest("GET", path, &transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
