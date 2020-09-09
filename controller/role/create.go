package role

import (
	"net/http"
	"strconv"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/gin-gonic/gin"
)

func CreateColumnRole(c *gin.Context) {
	var newStatement role
	if err := c.ShouldBindJSON(&newStatement); err != nil {
		c.JSON(http.StatusOK, response.ResponseBadRequest())
		return
	}

	columns := newStatement.Columns
	columnsName := newStatement.ColumnsName
	tableName := newStatement.TableName
	callerPrivateKey := newStatement.PrivateKey
	callerAddress, err := account.PrivateKeyToPublicKey(callerPrivateKey)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponsePKConvertError())
		return
	}

	columnsList := []string{}

	for _, column := range columns {
		columnsList = append(columnsList, strconv.Quote(column))
	}

	result, err := services.HasTableMaintainerRole(callerPrivateKey, callerAddress, tableName)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponseInternalError())
		return
	}

	switch result {
	case true:
		err = services.InitColumnsRole(callerPrivateKey, callerAddress, tableName, columnsList, columnsName)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		c.JSON(http.StatusOK, permissionResponse{0, "Ok"})
		logger.InfoAPIRoleCreate(newStatement)
	case false:
		c.JSON(http.StatusOK, response.ResponseUnauthorized())
		logger.WarnAPIRoleCreateUnAuth(callerAddress.String())
	}
}
