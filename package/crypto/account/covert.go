package account

import (
	"crypto/ecdsa"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/ethereum/go-ethereum/crypto"
)

func PrivateKeyToPublicKey(privatekeyHex string) (string, error) {
	privateKey, err := crypto.HexToECDSA(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "pk-converter").Warn(err.Error())
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.InternalLogger.WithField("component", "pk-converter").Warn(err.Error())
		return "", err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return fromAddress.String(), nil
}
