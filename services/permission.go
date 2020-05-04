package services

import (
	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/utils"
	"github.com/ethereum/go-ethereum/common"
)

func AddDropKitAdmin(privatekeyHex string, adminAccount common.Address) error {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "AddDropKitAdmin").Error(err.Error())
		return err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddDropKitAdmin").Error(err.Error())
		return err
	}

	_, err = contractInstance.AddDropKitAdmin(auth, adminAccount)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddDropKitAdmin").Error(err.Error())
		return err
	}

	return nil
}

func RemoveDropKitAdmin(privatekeyHex string, adminAccount common.Address) error {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "RemoveDropKitAdmin").Error(err.Error())
		return err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "RemoveDropKitAdmin").Error(err.Error())
		return err
	}

	_, err = contractInstance.RemoveDropKitAdmin(auth, adminAccount)
	if err != nil {
		logger.InternalLogger.WithField("component", "RemoveDropKitAdmin").Error(err.Error())
		return err
	}

	return nil
}

func AddTableAdmin(privatekeyHex string, adminAddress common.Address, tableName string) error {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "AddTableAdmin").Error(err.Error())
		return err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddTableAdmin").Error(err.Error())
		return err
	}
	adminRole, maintainerRole, userRole := utils.ToRoleName(tableName)

	_, err = contractInstance.AddTableAdmin(auth, adminAddress, adminRole, maintainerRole, userRole)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddTableAdmin").Error(err.Error())
		return err
	}

	return nil
}

func RemoveTableAdmin(privatekeyHex string, adminAddress common.Address, tableName string) error {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "RemoveTableAdmin").Error(err.Error())
		return err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "RemoveTableAdmin").Error(err.Error())
		return err
	}
	adminRole, maintainerRole, userRole := utils.ToRoleName(tableName)

	_, err = contractInstance.RemoveTableAdmin(auth, adminAddress, adminRole, maintainerRole, userRole)
	if err != nil {
		logger.InternalLogger.WithField("component", "RemoveTableAdmin").Error(err.Error())
		return err
	}

	return nil
}

func AddTableMaintainer(privatekeyHex string, maintainerAddress common.Address, tableName string) error {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "AddTableMaintainer").Error(err.Error())
		return err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddTableMaintainer").Error(err.Error())
		return err
	}
	adminRole, maintainerRole, userRole := utils.ToRoleName(tableName)

	_, err = contractInstance.AddTableMaintainer(auth, maintainerAddress, adminRole, maintainerRole, userRole)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddTableMaintainer").Error(err.Error())
		return err
	}

	return nil
}

func RemoveTableMaintainer(privatekeyHex string, maintainerAddress common.Address, tableName string) error {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "RemoveTableMaintainer").Error(err.Error())
		return err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "RemoveTableMaintainer").Error(err.Error())
		return err
	}
	adminRole, maintainerRole, userRole := utils.ToRoleName(tableName)

	_, err = contractInstance.RemoveTableMaintainer(auth, maintainerAddress, adminRole, maintainerRole, userRole)
	if err != nil {
		logger.InternalLogger.WithField("component", "RemoveTableMaintainer").Error(err.Error())
		return err
	}

	return nil
}

func AddTableUser(privatekeyHex string, userAddress common.Address, tableName string) error {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "AddTableUser").Error(err.Error())
		return err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddTableUser").Error(err.Error())
		return err
	}
	_, maintainerRole, userRole := utils.ToRoleName(tableName)

	_, err = contractInstance.AddTableUser(auth, userAddress, maintainerRole, userRole)
	if err != nil {
		logger.InternalLogger.WithField("component", "AddTableUser").Error(err.Error())
		return err
	}

	return nil
}

func RemoveTableUser(privatekeyHex string, userAddress common.Address, tableName string) error {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "RemoveTableUser").Error(err.Error())
		return err
	}
	auth, err := CreateAuthInstance(privatekeyHex)
	if err != nil {
		logger.InternalLogger.WithField("component", "RemoveTableUser").Error(err.Error())
		return err
	}
	_, maintainerRole, userRole := utils.ToRoleName(tableName)

	_, err = contractInstance.RemoveTableUser(auth, userAddress, maintainerRole, userRole)
	if err != nil {
		logger.InternalLogger.WithField("component", "RemoveTableUser").Error(err.Error())
		return err
	}

	return nil
}

func HasTableAdminRole(privatekeyHex string, accountAddress common.Address, tableName string) (bool, error) {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "HasTableAdminRole").Error(err.Error())
		return false, err
	}
	adminRole, _, _ := utils.ToRoleName(tableName)

	result, err := contractInstance.HasTableAdminRole(nil, accountAddress, adminRole)
	if err != nil {
		logger.InternalLogger.WithField("component", "HasTableAdminRole").Error(err.Error())
		return false, err
	}

	return result, nil
}

func HasTableMaintainerRole(privatekeyHex string, accountAddress common.Address, tableName string) (bool, error) {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "HasTableAdminRole").Error(err.Error())
		return false, err
	}
	_, maintainerRole, _ := utils.ToRoleName(tableName)

	result, err := contractInstance.HasTableMaintainerRole(nil, accountAddress, maintainerRole)
	if err != nil {
		logger.InternalLogger.WithField("component", "HasTableAdminRole").Error(err.Error())
		return false, err
	}

	return result, nil
}

func HasTableUserRole(privatekeyHex string, accountAddress common.Address, tableName string) (bool, error) {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "HasTableAdminRole").Error(err.Error())
		return false, err
	}
	_, _, userRole := utils.ToRoleName(tableName)

	result, err := contractInstance.HasTableUserRole(nil, accountAddress, userRole)
	if err != nil {
		logger.InternalLogger.WithField("component", "HasTableAdminRole").Error(err.Error())
		return false, err
	}

	return result, nil
}

func HasDropKitAdmin(privatekeyHex string, accountAddress common.Address) (bool, error) {
	contractInstance, err := CreateDropKitInstance()
	if err != nil {
		logger.InternalLogger.WithField("component", "HasDropKitAdmin").Error(err.Error())
		return false, err
	}

	result, err := contractInstance.HasRole(nil, [32]byte{0x0000000000000000000000000000000000000000000000000000000000000000}, accountAddress)
	if err != nil {
		logger.InternalLogger.WithField("component", "HasDropKitAdmin").Error(err.Error())
		return false, err
	}

	return result, nil
}
