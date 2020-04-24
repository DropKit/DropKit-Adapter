package account

import (
	"crypto/ecdsa"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/ethereum/go-ethereum/crypto"
)

func PrivateKeyToPublicKey(privatekeyHex string) string {
	privateKey, err := crypto.HexToECDSA(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.InternalLogger.WithField("component", "internal").Error(err.Error())
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return fromAddress.String()
}
