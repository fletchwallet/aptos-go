package aptos

import (
	"encoding/hex"
	"fmt"
	"time"
)

type UnsignedTx struct {
	Sender          string      `json:"sender"`
	SequenceNumber  uint64      `json:"sequence_number,string"`
	MaxGasAmount    uint64      `json:"max_gas_amount,string"`
	GasUnitPrice    uint64      `json:"gas_unit_price,string"`
	GasCurrencyCode string      `json:"gas_currency_code"`
	ExpirationTime  uint64      `json:"expiration_timestamp_secs,string"`
	Payload         interface{} `json:"payload"`
}

type SignedTx struct {
	*UnsignedTx
	Signature *TxSignature `json:"signature"`
}

type TxSignature struct {
	Type      string `json:"type"`
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
}

type ScriptFunctionPayload struct {
	Type          string   `json:"type"`
	Function      string   `json:"function"`
	TypeArguments []string `json:"type_arguments"`
	Arguments     []string `json:"arguments"`
}

type SigningMessage struct {
	Message string `json:"message"`
}

func (ac *AptosClient) GenerateTransaction(account *AptosAccount, payload interface{}) (*UnsignedTx, error) {
	acc, err := ac.Account(account.Address())
	if err != nil {
		return nil, err
	}

	return &UnsignedTx{
		Sender:          account.Address(),
		SequenceNumber:  acc.SequenceNumber,
		MaxGasAmount:    1000,
		GasUnitPrice:    1,
		GasCurrencyCode: "XUS",
		ExpirationTime:  uint64(time.Now().Add(10 * time.Second).Unix()),
		Payload:         payload,
	}, nil
}

func (ac *AptosClient) SignTransaction(account *AptosAccount, unsignedTx *UnsignedTx) (*SignedTx, error) {
	msg, err := ac.CreateSigningMessage(unsignedTx)
	if err != nil {
		return nil, err
	}

	msgHex, err := hex.DecodeString(msg.Message[2:])
	if err != nil {
		return nil, err
	}

	sig := &TxSignature{
		Type:      "ed25519_signature",
		PublicKey: account.PublicKey(),
		Signature: hex.EncodeToString(account.SignMessage(msgHex)),
	}

	return &SignedTx{
		UnsignedTx: unsignedTx,
		Signature:  sig,
	}, nil
}

func (ac *AptosClient) TransactionPending(txnHash string) (bool, error) {
	tx, err := ac.Transaction(txnHash)
	if err != nil {
		return false, err
	}

	return tx.Type == "pending_transaction", nil
}

func (ac *AptosClient) WaitForTransaction(txnHash string) error {
	count := 0
	for {
		time.Sleep(1 * time.Second)
		txPending, err := ac.TransactionPending(txnHash)
		if err != nil {
			return err
		}

		if !txPending {
			break
		}

		count++
		if count <= 10 {
			return fmt.Errorf("waiting for transaction %s timed out", txnHash)
		}
	}

	return nil
}
