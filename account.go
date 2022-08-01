package aptos

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/sha3"
)

type AptosAccount struct {
	privateKey ed25519.PrivateKey

	authKeyCached string
}

func AccountFromRandomKey() (*AptosAccount, error) {
	_, private, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	return &AptosAccount{
		privateKey: private,
	}, nil
}

func AccountFromPrivateKey(privateKey string) (*AptosAccount, error) {
	seed, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	private := ed25519.NewKeyFromSeed(seed[:32])

	return &AptosAccount{
		privateKey: private,
	}, nil
}

func (aa *AptosAccount) pubKeyBytes() []byte {
	pubKey := aa.privateKey.Public().(ed25519.PublicKey)
	return []byte(pubKey)
}

func (aa *AptosAccount) PublicKey() string {
	return fmt.Sprint("0x", hex.EncodeToString(aa.pubKeyBytes()))
}

func (aa *AptosAccount) Address() string {
	if aa.authKeyCached == "" {
		hasher := sha3.New256()

		hasher.Write(aa.pubKeyBytes())
		hasher.Write([]byte("\x00"))
		aa.authKeyCached = fmt.Sprint("0x", hex.EncodeToString(hasher.Sum(nil)))
	}

	return aa.authKeyCached
}

func (aa *AptosAccount) SignMessage(msg []byte) []byte {
	return ed25519.Sign(aa.privateKey, msg)
}

func (aa *AptosAccount) VerifyMessage(sig, msg []byte) bool {
	return ed25519.Verify(aa.pubKeyBytes(), msg, sig)
}
