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

func RevokeAdmin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIPermissionRevokeAdmin(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.Permission
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIPermissionRevokeAdmin(err)
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
		err = services.RemoveTableAdmin(callerPriavteKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		err = services.RevokeColumnsRole(callerPriavteKey, common.HexToAddress(grantUser), tableName, tableName+"ColumnsAdmin")
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		services.NormalResponse(w, response.PermissionResponseOk())
		logger.InfoAPIPermissionRevokeAdmin(newStatement)

	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
		logger.WarnAPIPermissionRevokeAdminUnAuth(callerAddress.String())
	}

}

func RevokeMaintainer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIPermissionRevokeMaintainer(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.Permission
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIPermissionRevokeMaintainer(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	callerPriavteKey := newStatement.PrivateKey
	grantUser := newStatement.UserName
	tableName := newStatement.TableName
	columnsRole := newStatement.ColumnsRole
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
		err = services.RemoveTableMaintainer(callerPriavteKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		err = services.RevokeColumnsRole(callerPriavteKey, common.HexToAddress(grantUser), tableName, columnsRole)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		services.NormalResponse(w, response.PermissionResponseOk())
		logger.InfoAPIPermissionRevokeMaintainer(newStatement)

	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
		logger.WarnAPIPermissionRevokeMaintainerUnAuth(callerAddress.String())
	}

}

func RevokeUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIPermissionRevokeUser(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.Permission
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIPermissionRevokeUser(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	callerPriavteKey := newStatement.PrivateKey
	grantUser := newStatement.UserName
	tableName := newStatement.TableName
	columnsRole := newStatement.ColumnsRole
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
		err = services.RemoveTableUser(callerPriavteKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		err = services.RevokeColumnsRole(callerPriavteKey, common.HexToAddress(grantUser), tableName, columnsRole)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		services.NormalResponse(w, response.PermissionResponseOk())
		logger.InfoAPIPermissionRevokeUser(newStatement)

	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
		logger.WarnAPIPermissionRevokeUserUnAuth(callerAddress.String())
	}

}
