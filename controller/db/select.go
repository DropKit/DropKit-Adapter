package controller

import (
	"encoding/json"
	"io/ioutil"

	"net/http"

	"github.com/DropKit/DropKit-Adapter/constants"
	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/DropKit/DropKit-Adapter/package/crypto/transaction"
	"github.com/DropKit/DropKit-Adapter/package/parser"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/spf13/viper"
)

func SQLSelect(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIDatabaseUpdate(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.SQL
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIDatabaseUpdate(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	sqlCommand := newStatement.Statement
	callerPriavteKey := newStatement.PrivateKey
	callerAddress, err := account.PrivateKeyToPublicKey(callerPriavteKey)
	if err != nil {
		services.NormalResponse(w, response.ResponsePKConvertError())
		return
	}

	tableName, err := parser.GetTableName(sqlCommand)
	if err != nil {
		services.NormalResponse(w, response.SQLResponseBadSQLStatement())
		return
	}

	result, err := services.HasTableUserRole(callerPriavteKey, callerAddress, tableName)
	if err != nil {
		services.NormalResponse(w, response.ResponseInternalError())
		return
	}

	switch result {
	case true:
		tableAddress, err := services.GetTableMeta(tableName)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		metadata, err := services.Query(sqlCommand)
		if err != nil {
			services.NormalResponse(w, response.SQLResponseDatabaseError(err))
			return
		}

		err = services.Consume(callerPriavteKey, callerAddress, viper.GetInt64(`PRICE.SELECT`))
		if err != nil {
			services.NormalResponse(w, response.ResponseExceedsBalance())
			return
		}

		aduitTransactionHash := transaction.SendRawTransaction(tableAddress, sqlCommand, 0, callerPriavteKey)
		services.NormalResponse(w, response.SQLQueryResponseOk(metadata, aduitTransactionHash))
		logger.InfoAPIDatabaseSelect(newStatement)

	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
		logger.WarnAPIDatabaseSelectUnAuth(callerAddress.String())
	}
}
