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

func RevokeTableOwner(c *gin.Context) {
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
		err = services.RemoveTableAdmin(callerPrivateKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		err = services.RevokeColumnsRole(callerPrivateKey, common.HexToAddress(grantUser), tableName, tableName+"ColumnsAdmin")
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		c.JSON(http.StatusOK, permissionResponse{0, "Ok"})
		logger.InfoAPIPermissionRevokeAdmin(newStatement)

	case false:
		c.JSON(http.StatusOK, response.ResponseUnauthorized())
		logger.WarnAPIPermissionRevokeAdminUnAuth(callerAddress.String())
	}

}

func RevokeTableMaintainer(c *gin.Context) {
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
		err = services.RemoveTableMaintainer(callerPrivateKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		err = services.RevokeColumnsRole(callerPrivateKey, common.HexToAddress(grantUser), tableName, columnsRole)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		c.JSON(http.StatusOK, permissionResponse{0, "Ok"})
		logger.InfoAPIPermissionRevokeMaintainer(newStatement)

	case false:
		c.JSON(http.StatusOK, response.ResponseUnauthorized())
		logger.WarnAPIPermissionRevokeMaintainerUnAuth(callerAddress.String())
	}

}

func RevokeTableViewer(c *gin.Context) {
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
		err = services.RemoveTableUser(callerPrivateKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		err = services.RevokeColumnsRole(callerPrivateKey, common.HexToAddress(grantUser), tableName, columnsRole)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		c.JSON(http.StatusOK, permissionResponse{0, "Ok"})
		logger.InfoAPIPermissionRevokeUser(newStatement)

	case false:
		c.JSON(http.StatusOK, response.ResponseUnauthorized())
		logger.WarnAPIPermissionRevokeUserUnAuth(callerAddress.String())
	}

}
