package controller

import (
	"net/http"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/ethereum/go-ethereum/common"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	account, ok := parameters["user_name"]
	if !ok {
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	balance, err := services.GetBalance(common.HexToAddress(account[0]))
	if err != nil {
		services.NormalResponse(w, response.ResponseInternalError())
		return
	}

	services.NormalResponse(w, response.PaymentBalanceResponseOk(balance.Int64()))
	logger.InfoAPIPaymentBalance(account[0])
}
