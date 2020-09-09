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

func VerifyTableOwner(c *gin.Context) {
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
		authority, err := services.HasTableAdminRole(callerPrivateKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		switch authority {
		case true:
			c.JSON(http.StatusOK, permissionVerifyResponse{0, "Ok", true})
			logger.InfoAPIPermissionVerifyAdmin(newStatement)

		case false:
			c.JSON(http.StatusOK, permissionVerifyResponse{0, "Ok", false})
			logger.InfoAPIPermissionVerifyAdmin(newStatement)
		}
	case false:
		c.JSON(http.StatusOK, response.ResponseUnauthorized())
		logger.WarnAPIPermissionVerifyAdminUnAuth(callerAddress.String())
	}

}

func VerifyTableMaintainer(c *gin.Context) {
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
		authority, err := services.HasTableMaintainerRole(callerPrivateKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		switch authority {
		case true:
			c.JSON(http.StatusOK, permissionVerifyResponse{0, "Ok", true})
			logger.InfoAPIPermissionVerifyMaintainer(newStatement)

		case false:
			c.JSON(http.StatusOK, permissionVerifyResponse{0, "Ok", false})
			logger.InfoAPIPermissionVerifyMaintainer(newStatement)
		}
	case false:
		c.JSON(http.StatusOK, response.ResponseUnauthorized())
		logger.WarnAPIPermissionRevokeMaintainerUnAuth(callerAddress.String())
	}

}

func TableViewer(c *gin.Context) {
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

	result, err := services.HasTableMaintainerRole(callerPrivateKey, callerAddress, tableName)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponseInternalError())
		return
	}

	switch result {
	case true:
		authority, err := services.HasTableUserRole(callerPrivateKey, common.HexToAddress(grantUser), tableName)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		switch authority {
		case true:
			c.JSON(http.StatusOK, permissionVerifyResponse{0, "Ok", true})
			logger.InfoAPIPermissionVerifyUser(newStatement)

		case false:
			c.JSON(http.StatusOK, permissionVerifyResponse{0, "Ok", false})
			logger.InfoAPIPermissionVerifyUser(newStatement)
		}
	case false:
		c.JSON(http.StatusOK, response.ResponseUnauthorized())
		logger.WarnAPIPermissionVerifyUserUnAuth(callerAddress.String())
	}

}
