package controller

import (
	"encoding/json"
	"github.com/DropKit/DropKit-Adapter/constants"
	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/DropKit/DropKit-Adapter/package/crypto/transaction"
	"github.com/DropKit/DropKit-Adapter/package/parser"
	columns "github.com/DropKit/DropKit-Adapter/package/parser/columns"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/package/utils"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
)

func checkSelectAll(sqlCommand string, columnsCanSelect []string) string {
	var idx int = -1
	for i, v := range sqlCommand {
		if v == '*' {
			idx = i
			break
		}
	}
	if idx == -1 {
		return sqlCommand
	}

	sqlSlice := []byte(sqlCommand[0:idx])
	columnsStr := strings.Join(columnsCanSelect, ",")
	sqlSlice = append(sqlSlice, []byte(columnsStr)...)
	sqlSlice = append(sqlSlice, []byte(sqlCommand[idx+1:])...)
	return string(sqlSlice)
}

func HandleDBSelection(w http.ResponseWriter, r *http.Request) {
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

		columnsCanSelect, err := services.GetColumnsRole(callerPriavteKey, callerAddress, tableName)

		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		sqlCommand = checkSelectAll(sqlCommand, columnsCanSelect)

		columnsNames, err := columns.GetSelectColumns(sqlCommand)
		if err != nil {
			services.NormalResponse(w, response.SQLResponseBadSQLStatement())
			return
		}

		columnsAuth := utils.CompareColumns(columnsCanSelect, columnsNames)

		switch columnsAuth {
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

	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
		logger.WarnAPIDatabaseSelectUnAuth(callerAddress.String())
	}
}
