package account

import (
	"crypto/ecdsa"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateWallet() (string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		logger.InternalLogger.WithField("component", "account-creater").Error(err.Error())
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.InternalLogger.WithField("component", "account-creater").Error(err.Error())
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return hexutil.Encode(privateKeyBytes)[2:], address
}
