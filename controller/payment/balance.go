package payment

import (
	"net/http"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func GetAccountBalance(c *gin.Context) {
	var newStatement tokenBalance
	if err := c.ShouldBindJSON(&newStatement); err != nil {
		c.JSON(http.StatusOK, response.ResponseBadRequest())
		return
	}

	callerPrivateKey := newStatement.PrivateKey
	balanceAccount := newStatement.Account
	callerAddress, err := account.PrivateKeyToPublicKey(callerPrivateKey)
	if err != nil {
		c.JSON(http.StatusOK, response.ResponsePKConvertError())
		return
	}

	switch balanceAccount {
	case "":
		balance, err := services.GetAccountBalance(callerAddress)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		c.JSON(http.StatusOK, tokenBalanceResponse{0, "Ok", balance.Int64()})
		logger.InfoAPIPaymentBalance(newStatement)
	default:
		result, err := services.HasDropKitAdmin(callerPrivateKey, callerAddress)
		if err != nil {
			c.JSON(http.StatusOK, response.ResponseInternalError())
			return
		}

		switch result {
		case true:
			balance, err := services.GetAccountBalance(common.HexToAddress(balanceAccount))
			if err != nil {
				c.JSON(http.StatusOK, response.ResponseInternalError())
				return
			}

			c.JSON(http.StatusOK, tokenBalanceResponse{0, "Ok", balance.Int64()})
			logger.InfoAPIPaymentBalance(balanceAccount)
		case false:
			c.JSON(http.StatusOK, response.ResponseUnauthorized())
			logger.WarnAPIPaymentBurnUnAuth(callerAddress.String())
		}

	}
}
