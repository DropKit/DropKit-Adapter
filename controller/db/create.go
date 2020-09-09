package db

import (
	"net/http"
	"strings"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/DropKit/DropKit-Adapter/package/crypto/transaction"
	"github.com/DropKit/DropKit-Adapter/package/parser"
	columns "github.com/DropKit/DropKit-Adapter/package/parser/columns"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func HandleDBCreation(c *gin.Context) {
	var newStatement sql
	if err := c.ShouldBindJSON(&newStatement); err != nil {
		c.JSON(http.StatusOK, response.ResponseBadRequest())
		return
	}

	sqlCommand := newStatement.Statement
	callerPrivateKey := newStatement.PrivateKey
	callerAddress, err := account.PrivateKeyToPublicKey(callerPrivateKey)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponsePKConvertError())
		return
	}

	tableName, err := parser.GetTableName(sqlCommand)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorResponse{Code: 20202, Message: "Bad SQL statement"})
		return
	}

	columnsNames, err := columns.GetCreateColumns(sqlCommand)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponseInternalError())
		return
	}

	result, err := services.HasDropKitAdmin(callerPrivateKey, callerAddress)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponseInternalError())
		return
	}

	switch result {
	case true:
		tableAddress, err := services.AddTableMeta(tableName, callerPrivateKey)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		err = services.AddTableAdmin(callerPrivateKey, callerAddress, tableName)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		err = services.Exec(sqlCommand)
		if err != nil {
			c.JSON(http.StatusOK, response.ErrorResponseWithReason{Code: 20201, Message: "Database error", Reason: (strings.Split(err.Error(), "pq: "))[1]})
			return
		}

		err = services.AddColumnsRole(callerPrivateKey, callerAddress, tableName, columnsNames, tableName+"ColumnsAdmin")
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		err = services.Consume(callerPrivateKey, callerAddress, viper.GetInt64(`PRICE.CREATE`))
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseExceedsBalance())
			return
		}

		auditTransactionHash := transaction.SendRawTransaction(tableAddress, sqlCommand, 0, callerPrivateKey)
		c.JSON(http.StatusOK, sqlExecResponse{0, "Ok", auditTransactionHash})
		logger.InfoAPIDatabaseCreate(newStatement)

	case false:
		c.JSON(http.StatusOK, response.ResponseUnauthorized())
		logger.WarnAPIDatabaseCreateUnAuth(callerAddress.String())
	}
}
