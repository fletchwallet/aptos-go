aptos-go
===================
This is a pure Go implementation of the API's available in Aptos: https://aptos.dev.  
This library offers *all* of the API's present in Aptos, as well as some utilities for generating and loading wallets, encrypting and verifying encrypted messages.

### Requirements
Aptos node (you need to connect to an Aptos node to use this library, or run one yourself locally.)  
Go 1.18 or newer


### Installation

```
go get github.com/c-ollins/aptos-go
```

### Usage example
```go
func _main() error {
	client := aptos.NewAptosClient(NODE_URL)
	account1, err := aptos.AccountFromRandomKey()
	if err != nil {
		return err
	}

	err = client.FundFromFaucet(FAUCET_URL, account1.Address(), 5000)
	if err != nil {
		return err
	}

	balance1, err := client.AccountBalance(account1.Address())
	if err != nil {
		return err
	}
	fmt.Printf("account1 coins: %d. Should be 5000!\n", balance1)

	account2, err := aptos.AccountFromRandomKey()
	if err != nil {
		return err
	}

	err = client.FundFromFaucet(FAUCET_URL, account2.Address(), 0)
	if err != nil {
		return err
	}

	balance2, err := client.AccountBalance(account2.Address())
	if err != nil {
		return err
	}
	fmt.Printf("account2 coins: %d. Should be 0!\n", balance2)

	tx, err := client.SendCoinsSync(account1, account2.Address(), 717)
	if err != nil {
		return err
	}

	fmt.Println("Txn successful, Hash:", tx.Hash)

	balance2, err = client.AccountBalance(account2.Address())
	if err != nil {
		return err
	}

	fmt.Printf("account2 coins: %d. Should be 717!\n", balance2)
	return nil
}
```