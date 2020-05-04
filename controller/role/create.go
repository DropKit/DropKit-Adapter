package controller

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"net/http"

	"github.com/DropKit/DropKit-Adapter/constants"
	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
)

func CreateColumnRole(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIRoleCreate(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.Role
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIDatabaseCreate(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	columns := newStatement.Columns
	columnsName := newStatement.ColumnsName
	tableName := newStatement.TableName
	callerPriavteKey := newStatement.PrivateKey
	callerAddress, err := account.PrivateKeyToPublicKey(callerPriavteKey)
	if err != nil {
		services.NormalResponse(w, response.ResponsePKConvertError())
		return
	}

	columnsList := []string{}

	for _, column := range columns {
		columnsList = append(columnsList, strconv.Quote(column))
	}

	result, err := services.HasTableMaintainerRole(callerPriavteKey, callerAddress, tableName)
	if err != nil {
		services.NormalResponse(w, response.ResponseInternalError())
		return
	}

	switch result {
	case true:
		err = services.InitColumnsRole(callerPriavteKey, callerAddress, tableName, columnsList, columnsName)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		services.NormalResponse(w, response.PermissionResponseOk())
		logger.InfoAPIRoleCreate(newStatement)
	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
		logger.WarnAPIRoleCreateUnAuth(callerAddress.String())
	}
}
