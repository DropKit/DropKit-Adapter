package account

import (
	"crypto/ecdsa"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func PrivateKeyToPublicKey(privatekeyHex string) (common.Address, error) {
	privateKey, err := crypto.HexToECDSA(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "pk-converter").Warn(err.Error())
		return common.HexToAddress(""), err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.InternalLogger.WithField("component", "pk-converter").Warn(err.Error())
		return common.HexToAddress(""), err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return fromAddress, nil
}
