package services

import (
	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/utils"
	"github.com/ethereum/go-ethereum/common"
)

func InitColumnsRole(privatekeyHex string, userAddress common.Address, tableName string, columns []string, columnRoleName string) error {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "AddColumnsRole").Error(err.Error())
		return err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddColumnsRole").Error(err.Error())
		return err
	}
	_, maintainerRole, _ := utils.ToRoleName(tableName)

	_, err = contractInstance.AddColumnsRole(auth, userAddress, maintainerRole, columnRoleName, columns)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddColumnsRole").Error(err.Error())
		return err
	}

	return nil
}

func AddColumnsRole(privatekeyHex string, userAddress common.Address, tableName string, columns []string, columnRoleName string) error {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "AddColumnsRole").Error(err.Error())
		return err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddColumnsRole").Error(err.Error())
		return err
	}
	_, maintainerRole, _ := utils.ToRoleName(tableName)

	_, err = contractInstance.AddColumnsRole(auth, userAddress, maintainerRole, columnRoleName, columns)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddColumnsRole").Error(err.Error())
		return err
	}

	err = SetUserColumnsRole(privatekeyHex, userAddress, tableName, columnRoleName)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddColumnsRole").Error(err.Error())
		return err
	}

	return nil
}

func GrantColumnsRole(privatekeyHex string, userAddress common.Address, tableName string, columnRoleName string) error {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "GrantColumnsRole").Error(err.Error())
		return err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "GrantColumnsRole").Error(err.Error())
		return err
	}
	_, maintainerRole, _ := utils.ToRoleName(tableName)

	_, err = contractInstance.GrantColumnsRole(auth, userAddress, maintainerRole, columnRoleName)
	if err != nil {
		logger.InternalLogger.WithField("component", "GrantColumnsRole").Error(err.Error())
		return err
	}

	err = SetUserColumnsRole(privatekeyHex, userAddress, tableName, columnRoleName)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddColumnsRole").Error(err.Error())
		return err
	}

	return nil
}

func RevokeColumnsRole(privatekeyHex string, userAddress common.Address, tableName string, columnRoleName string) error {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "RevokeColumnsRole").Error(err.Error())
		return err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "RevokeColumnsRole").Error(err.Error())
		return err
	}
	_, maintainerRole, _ := utils.ToRoleName(tableName)

	_, err = contractInstance.RevokeColumnsRole(auth, userAddress, maintainerRole, columnRoleName)
	if err != nil {
		logger.InternalLogger.WithField("component", "RevokeColumnsRole").Error(err.Error())
		return err
	}

	err = SetUserColumnsRole(privatekeyHex, userAddress, tableName, "")
	if err != nil {
		logger.InternalLogger.WithField("component", "AddColumnsRole").Error(err.Error())
		return err
	}

	return nil
}

func GetColumnsRole(privatekeyHex string, userAddress common.Address, tableName string) ([]string, error) {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "GetColumnsRole").Error(err.Error())
		return nil, err
	}
	_, _, userRole := utils.ToRoleName(tableName)

	// columnRoleName, err := contractInstance.GetUserColumnsRole(nil, userAddress, tableName)
	// if err != nil {
	// 	logger.InternalLogger.WithField("component", "AddColumnsRole").Error(err.Error())
	// 	return nil, err
	// }

	columnRoleName, err := GetUserColumnsRole(userAddress, tableName)

	columns, err := contractInstance.GetColumnsRole(nil, userAddress, userRole, columnRoleName)
	if err != nil {
		logger.InternalLogger.WithField("component", "GetColumnsRole").Error(err.Error())
		return nil, err
	}

	return columns, nil
}

func SetUserColumnsRole(privatekeyHex string, userAddress common.Address, tableName string, columnRoleName string) error {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "SetUserColumnsRole").Error(err.Error())
		return err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "SetUserColumnsRole").Error(err.Error())
		return err
	}

	_, maintainerRole, _ := utils.ToRoleName(tableName)

	_, err = contractInstance.SetUserColumnsRole(auth, userAddress, tableName, maintainerRole, columnRoleName)
	if err != nil {
		logger.InternalLogger.WithField("component", "SetUserColumnsRole").Error(err.Error())
		return err
	}

	return nil
}

func GetUserColumnsRole(userAddress common.Address, tableName string) (string, error) {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "GetColumnsRole").Error(err.Error())
		return "", err
	}

	columnRoleName, err := contractInstance.GetUserColumnsRole(nil, userAddress, tableName)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddColumnsRole").Error(err.Error())
		return "", err
	}

	return columnRoleName, nil
}
