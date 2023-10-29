package secrets

import (
	"encoding/base64"
	"fmt"

	"github.com/cometbft/cometbft/crypto/ed25519"
)

type CometKeys struct {
	Address    string `json:"address"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

func GenerateCometKeys() CometKeys {
	privKey := ed25519.GenPrivKey()
	pubKey := privKey.PubKey()
	address := pubKey.Address()
	return CometKeys{
		Address:    address.String(),
		PublicKey:  base64.StdEncoding.EncodeToString(pubKey.Bytes()),
		PrivateKey: base64.StdEncoding.EncodeToString(privKey.Bytes()),
	}
}

func ValidateCometPublicKeyAndAddressWithPrivateKey(publicKey string, address string, privateKey string) error {
	var (
		privKey           ed25519.PrivKey
		expectedPublicKey string
		expectedAddress   string
		err               error
	)

	if privKey, err = base64.StdEncoding.DecodeString(privateKey); err != nil {
		return fmt.Errorf("failed to decode privateKey %w", err)
	}

	expectedPublicKey = base64.StdEncoding.EncodeToString(privKey.PubKey().Bytes())
	expectedAddress = privKey.PubKey().Address().String()

	if publicKey != expectedPublicKey {
		return fmt.Errorf("comet public key does not match the one derived from private key, provided='%s', expected='%s'", publicKey, expectedPublicKey)
	}
	if address != expectedAddress {
		return fmt.Errorf("comet address does not match the one derived from private key, provided='%s', expected='%s'", address, expectedAddress)
	}

	return nil
}
