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

func GrantTableOwner(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		services.NormalResponse(w, response.ResponseBadRequest())
		logger.WarnAPIPermissionGrantAdmin(err)
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.Permission
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIPermissionGrantAdmin(err)
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
		err = services.AddTableAdmin(callerPriavteKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		err = services.GrantColumnsRole(callerPriavteKey, common.HexToAddress(grantUser), tableName, tableName+"ColumnsAdmin")
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		services.NormalResponse(w, response.PermissionResponseOk())
		logger.InfoAPIPermissionGrantAdmin(newStatement)
	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
		logger.WarnAPIPermissionGrantAdminUnAuth(callerAddress.String())
	}
}

func GrantTableMaintainer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		services.NormalResponse(w, response.ResponseBadRequest())
		logger.WarnAPIPermissionGrantMaintainer(err)
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.Permission
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIPermissionGrantMaintainer(err)
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
		err = services.AddTableMaintainer(callerPriavteKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		err = services.GrantColumnsRole(callerPriavteKey, common.HexToAddress(grantUser), tableName, columnsRole)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		services.NormalResponse(w, response.PermissionResponseOk())
		logger.InfoAPIPermissionGrantMaintainer(newStatement)

	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
		logger.WarnAPIPermissionGrantMaintainerUnAuth(callerAddress.String())
	}

}

func GrantTableViewer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		services.NormalResponse(w, response.ResponseBadRequest())
		logger.WarnAPIPermissionGrantUser(err)
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.Permission
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIPermissionGrantUser(err)
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
		err = services.AddTableUser(callerPriavteKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		err = services.GrantColumnsRole(callerPriavteKey, common.HexToAddress(grantUser), tableName, columnsRole)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		services.NormalResponse(w, response.PermissionResponseOk())
		logger.InfoAPIPermissionGrantUser(newStatement)

	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
		logger.WarnAPIPermissionGrantUserUnAuth(callerAddress.String())
	}
}
