package services

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
)

func MintToken(amount int64, privatekeyHex string, mintAccount common.Address) (string, error) {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "MintToken").Error(err.Error())
		return "", err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "MintToken").Error(err.Error())
		return "", err
	}

	tx, err := contractInstance.MintToken(auth, mintAccount, big.NewInt(amount))
	if err != nil {
		logger.InternalLogger.WithField("component", "MintToken").Error(err.Error())
		return "", err
	}

	return tx.Hash().String(), nil
}

func BurnToken(amount int64, privatekeyHex string, burnAccount common.Address) (string, error) {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "BurnToken").Error(err.Error())
		return "", err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "BurnToken").Error(err.Error())
		return "", err
	}

	tx, err := contractInstance.MintToken(auth, burnAccount, big.NewInt(amount))
	if err != nil {
		logger.InternalLogger.WithField("component", "BurnToken").Error(err.Error())
		return "", err
	}

	return tx.Hash().String(), nil
}

func Consume(privatekeyHex string, account common.Address, amount int64) error {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "Consume").Error(err.Error())
		return err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "Consume").Error(err.Error())
		return err
	}

	_, err = contractInstance.Consume(auth, account, big.NewInt(amount))
	if err != nil {
		logger.InternalLogger.WithField("component", "Consume").Error(err.Error())
		return err
	}

	return nil
}

func GetAccountBalance(balanceAccount common.Address) (*big.Int, error) {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "GetAccountBalance").Error(err.Error())
		return big.NewInt(0), err
	}

	amount, err := contractInstance.BalanceOf(nil, balanceAccount)
	if err != nil {
		logger.InternalLogger.WithField("component", "GetAccountBalance").Error(err.Error())
		return big.NewInt(0), err
	}

	return amount, nil
}

func Transfer(privatekeyHex string, toAccount common.Address, amount int64) (string, error) {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "Transfer").Error(err.Error())
		return "", err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "Transfer").Error(err.Error())
		return "", err
	}

	fromAccount, err := account.PrivateKeyToPublicKey(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "Transfer").Error(err.Error())
		return "", err
	}

	fromAccountBalance, err := contractInstance.BalanceOf(nil, fromAccount)
	if err != nil {
		logger.InternalLogger.WithField("component", "Transfer").Error(err.Error())
		return "", err
	}

	enough := fromAccountBalance.Cmp(big.NewInt(amount))

	switch enough {
	case -1:
		return "", errors.New("-1")

	default:
		tx, err := contractInstance.Transfer(auth, toAccount, big.NewInt(amount))
		if err != nil {
			logger.InternalLogger.WithField("component", "Transfer").Error(err.Error())
			return "", err
		}

		return tx.Hash().String(), nil
	}
}
