package controller

import (
	"encoding/json"
	"io/ioutil"

	"net/http"

	"github.com/DropKit/DropKit-Adapter/constants"
	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/ethereum/go-ethereum/common"
)

func VerifyTableOwner(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIPermissionVerifyAdmin(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.Permission
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIPermissionVerifyAdmin(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	callerPriavteKey := newStatement.PrivateKey
	grantUser := newStatement.UserName
	tableName := newStatement.TableName
	callerAddress, err := account.PrivateKeyToPublicKey(callerPriavteKey)
	if err != nil {
		services.NormalResponse(w, response.ResponsePKConvertError())
		return
	}

	result, err := services.HasTableAdminRole(callerPriavteKey, callerAddress, tableName)
	if err != nil {
		services.NormalResponse(w, response.ResponseInternalError())
		return
	}

	switch result {
	case true:
		authority, err := services.HasTableAdminRole(callerPriavteKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		switch authority {
		case true:
			services.NormalResponse(w, response.PermissionVerifyResponse(true))
			logger.InfoAPIPermissionVerifyAdmin(newStatement)

		case false:
			services.NormalResponse(w, response.PermissionVerifyResponse(false))
			logger.InfoAPIPermissionVerifyAdmin(newStatement)
		}
	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
		logger.WarnAPIPermissionVerifyAdminUnAuth(callerAddress.String())
	}

}

func VerifyTableMaintainer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIPermissionVerifyMaintainer(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.Permission
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIPermissionVerifyMaintainer(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	callerPriavteKey := newStatement.PrivateKey
	grantUser := newStatement.UserName
	tableName := newStatement.TableName
	callerAddress, err := account.PrivateKeyToPublicKey(callerPriavteKey)
	if err != nil {
		services.NormalResponse(w, response.ResponsePKConvertError())
		return
	}

	result, err := services.HasTableAdminRole(callerPriavteKey, callerAddress, tableName)
	if err != nil {
		services.NormalResponse(w, response.ResponseInternalError())
		return
	}

	switch result {
	case true:
		authority, err := services.HasTableMaintainerRole(callerPriavteKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		switch authority {
		case true:
			services.NormalResponse(w, response.PermissionVerifyResponse(true))
			logger.InfoAPIPermissionVerifyMaintainer(newStatement)

		case false:
			services.NormalResponse(w, response.PermissionVerifyResponse(false))
			logger.InfoAPIPermissionVerifyMaintainer(newStatement)
		}
	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
		logger.WarnAPIPermissionRevokeMaintainerUnAuth(callerAddress.String())
	}

}

func TableViewer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIPermissionVerifyUser(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.Permission
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIPermissionVerifyUser(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	callerPriavteKey := newStatement.PrivateKey
	grantUser := newStatement.UserName
	tableName := newStatement.TableName
	callerAddress, err := account.PrivateKeyToPublicKey(callerPriavteKey)
	if err != nil {
		services.NormalResponse(w, response.ResponsePKConvertError())
		return
	}

	result, err := services.HasTableMaintainerRole(callerPriavteKey, callerAddress, tableName)
	if err != nil {
		services.NormalResponse(w, response.ResponseInternalError())
		return
	}

	switch result {
	case true:
		authority, err := services.HasTableUserRole(callerPriavteKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		switch authority {
		case true:
			services.NormalResponse(w, response.PermissionVerifyResponse(true))
			logger.InfoAPIPermissionVerifyUser(newStatement)

		case false:
			services.NormalResponse(w, response.PermissionVerifyResponse(false))
			logger.InfoAPIPermissionVerifyUser(newStatement)
		}
	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
		logger.WarnAPIPermissionVerifyUserUnAuth(callerAddress.String())
	}

}
