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

func SQLDelete(w http.ResponseWriter, r *http.Request) {
	metaTableAddress := viper.GetString(`DROPKIT.METATABLE`)
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIDatabaseDelete(err)
	}

	var newStatement constants.SQL
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIDatabaseDelete(err)
	}
	logger.InfoAPIDatabaseDelete(newStatement)

	sqlCommand := newStatement.Statement
	callerPriavteKey := newStatement.PrivateKey
	callerAddress := account.PrivateKeyToPublicKey(callerPriavteKey)

	tableName := parser.GetTableName(sqlCommand)
	tableAddress := services.GetMetaTable(tableName, metaTableAddress)
	authority := services.VerifyAuthority(authorityAddr, callerPriavteKey, tableName, callerAddress)

	switch authority {
	case true:
		services.Exec(sqlCommand)
		aduitTransactionHash := transaction.SendRawTransaction(tableAddress, sqlCommand, 0, callerPriavteKey)
		defer r.Body.Close()
		services.NormalResponse(w, response.SQLExecResponseOk(aduitTransactionHash))
	case false:
		defer r.Body.Close()
		services.NormalResponse(w, response.SQLResponseUnauthorized())
	}
}
