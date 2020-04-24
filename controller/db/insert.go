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

func SQLInsert(w http.ResponseWriter, r *http.Request) {
	metaTableAddress := viper.GetString(`DROPKIT.METATABLE`)
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIDatabaseInsert(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.SQL
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIDatabaseInsert(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}
	logger.InfoAPIDatabaseInsert(newStatement)

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
	tableAddress := services.GetMetaTable(tableName, metaTableAddress)
	authority, _ := services.VerifyAuthority(authorityAddr, callerPriavteKey, tableName, callerAddress)

	switch authority {
	case true:
		err = services.Exec(sqlCommand)
		if err != nil {
			services.NormalResponse(w, response.SQLResponseDatabaseError(err))
			return
		}
		aduitTransactionHash := transaction.SendRawTransaction(tableAddress, sqlCommand, 0, callerPriavteKey)
		services.NormalResponse(w, response.SQLExecResponseOk(aduitTransactionHash))
	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
	}
}
