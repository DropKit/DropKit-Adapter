package payment

import (
	"net/http"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/gin-gonic/gin"
)

func MintToken(c *gin.Context) {
	var newStatement token
	if err := c.ShouldBindJSON(&newStatement); err != nil {
		c.JSON(http.StatusOK, response.ResponseBadRequest())
		return
	}

	callerPrivateKey := newStatement.PrivateKey
	amount := newStatement.Amount
	callerAddress, err := account.PrivateKeyToPublicKey(callerPrivateKey)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponsePKConvertError())
		return
	}

	result, err := services.HasDropKitAdmin(callerPrivateKey, callerAddress)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponseInternalError())
		return
	}

	switch result {
	case true:
		hash, err := services.MintToken(amount, callerPrivateKey, callerAddress)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}
		c.JSON(http.StatusOK, tokenTransferResponse{0, "Ok", hash})
		logger.InfoAPIPaymentMint(newStatement)
	case false:
		c.JSON(http.StatusOK, response.ResponseUnauthorized())
		logger.WarnAPIPaymentMintUnAuth(callerAddress.String())
	}
}
