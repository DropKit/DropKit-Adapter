package payment

import (
	"net/http"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func TransferToken(c *gin.Context) {
	var newStatement tokenTransfer
	if err := c.ShouldBindJSON(&newStatement); err != nil {
		c.JSON(http.StatusOK, response.ResponseBadRequest())
		return
	}

	callerPrivateKey := newStatement.PrivateKey
	amount := newStatement.Amount
	to := newStatement.Account

	hash, err := services.Transfer(callerPrivateKey, common.HexToAddress(to), amount)
	if err != nil {
		if err.Error() == "-1" {
			c.JSON(http.StatusOK, tokenTransferFailResponse{20401, "not enough balance"})
			logger.WarnAPIPaymentTransferNotEnough(newStatement)
			return
		}
		c.JSON(http.StatusOK, response.ResponseInternalError())
		logger.WarnAPIPaymentTransfer(err)
		return
	}

	c.JSON(http.StatusOK, tokenTransferResponse{0, "Ok", hash})
	logger.InfoAPIPaymentTransfer(newStatement)
}
