package permission

import (
	"net/http"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func GrantTableOwner(c *gin.Context) {
	var newStatement permission
	if err := c.ShouldBindJSON(&newStatement); err != nil {
		c.JSON(http.StatusOK, response.ResponseBadRequest())
		return
	}

	callerPrivateKey := newStatement.PrivateKey
	grantUser := newStatement.UserName
	tableName := newStatement.TableName
	callerAddress, err := account.PrivateKeyToPublicKey(callerPrivateKey)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponsePKConvertError())
		return
	}

	result, err := services.HasTableAdminRole(callerPrivateKey, callerAddress, tableName)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponseInternalError())
		return
	}

	switch result {
	case true:
		err = services.AddTableAdmin(callerPrivateKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		err = services.GrantColumnsRole(callerPrivateKey, common.HexToAddress(grantUser), tableName, tableName+"ColumnsAdmin")
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		c.JSON(http.StatusOK, permissionResponse{0, "Ok"})
		logger.InfoAPIPermissionGrantAdmin(newStatement)
	case false:
		c.JSON(http.StatusOK, response.ResponseUnauthorized())
		logger.WarnAPIPermissionGrantAdminUnAuth(callerAddress.String())
	}
}

func GrantTableMaintainer(c *gin.Context) {
	var newStatement permission
	if err := c.ShouldBindJSON(&newStatement); err != nil {
		c.JSON(http.StatusOK, response.ResponseBadRequest())
		return
	}

	callerPrivateKey := newStatement.PrivateKey
	grantUser := newStatement.UserName
	tableName := newStatement.TableName
	columnsRole := newStatement.ColumnsRole
	callerAddress, err := account.PrivateKeyToPublicKey(callerPrivateKey)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponsePKConvertError())
		return
	}

	result, err := services.HasTableAdminRole(callerPrivateKey, callerAddress, tableName)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponseInternalError())
		return
	}

	switch result {
	case true:
		err = services.AddTableMaintainer(callerPrivateKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		err = services.GrantColumnsRole(callerPrivateKey, common.HexToAddress(grantUser), tableName, columnsRole)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		c.JSON(http.StatusOK, permissionResponse{0, "Ok"})
		logger.InfoAPIPermissionGrantMaintainer(newStatement)

	case false:
		c.JSON(http.StatusOK, response.ResponseUnauthorized())
		logger.WarnAPIPermissionGrantMaintainerUnAuth(callerAddress.String())
	}

}

func GrantTableViewer(c *gin.Context) {
	var newStatement permission
	if err := c.ShouldBindJSON(&newStatement); err != nil {
		c.JSON(http.StatusOK, response.ResponseBadRequest())
		return
	}

	callerPrivateKey := newStatement.PrivateKey
	grantUser := newStatement.UserName
	tableName := newStatement.TableName
	columnsRole := newStatement.ColumnsRole
	callerAddress, err := account.PrivateKeyToPublicKey(callerPrivateKey)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponsePKConvertError())
		return
	}

	result, err := services.HasTableMaintainerRole(callerPrivateKey, callerAddress, tableName)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponseInternalError())
		return
	}

	switch result {
	case true:
		err = services.AddTableUser(callerPrivateKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		err = services.GrantColumnsRole(callerPrivateKey, common.HexToAddress(grantUser), tableName, columnsRole)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		c.JSON(http.StatusOK, permissionResponse{0, "Ok"})
		logger.InfoAPIPermissionGrantUser(newStatement)

	case false:
		c.JSON(http.StatusOK, response.ResponseUnauthorized())
		logger.WarnAPIPermissionGrantUserUnAuth(callerAddress.String())
	}
}
