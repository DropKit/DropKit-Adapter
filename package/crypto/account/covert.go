package account

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func PrivateKeyToPublicKey(privatekeyHex string) string {

	privateKey, err := crypto.HexToECDSA(privatekeyHex)
	if err != nil {
		print(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		print(err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return fromAddress.String()
}
