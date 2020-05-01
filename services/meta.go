package services

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
)

func AddTableMeta(tableName string, privatekeyHex string) (string, error) {
	_, tableAddress := account.GenerateWallet()

	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "AddTableMeta").Error(err.Error())
		return "", err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddTableMeta").Error(err.Error())
		return "", err
	}

	_, err = contractInstance.AddTableMeta(auth, tableName, common.HexToAddress(tableAddress))
	if err != nil {
		logger.InternalLogger.WithField("component", "AddTableMeta").Error(err.Error())
		return "", err
	}

	return tableAddress, nil
}

func GetTableMeta(tableName string) (string, error) {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "GetTableMeta").Error(err.Error())
		return "", err
	}

	tableAddress, err := contractInstance.GetTableMeta(nil, tableName)
	if err != nil {
		logger.InternalLogger.WithField("component", "GetTableMeta").Error(err.Error())
		return "", err
	}

	return tableAddress.String(), err
}
