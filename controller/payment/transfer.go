package controller

import (
	"encoding/json"
	"io/ioutil"

	"net/http"

	"github.com/DropKit/DropKit-Adapter/constants"
	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/ethereum/go-ethereum/common"
)

func TransferToken(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIPaymentTransfer(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.TokenTransfer
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		services.NormalResponse(w, response.ResponseBadRequest())
		logger.WarnAPIPaymentTransfer(err)
		return
	}

	callerPriavteKey := newStatement.PrivateKey
	amount := newStatement.Amount
	to := newStatement.Account

	hash, err := services.Transfer(callerPriavteKey, common.HexToAddress(to), amount)
	if err != nil {
		if err.Error() == "-1" {
			services.NormalResponse(w, response.PaymentResponseNotEnough())
			logger.WarnAPIPaymentTransferNotEnough(newStatement)
			return
		}
		services.NormalResponse(w, response.ResponseInternalError())
		logger.WarnAPIPaymentTransfer(err)
		return
	}

	services.NormalResponse(w, response.PaymentResponseOk(hash))
	logger.InfoAPIPaymentTransfer(newStatement)
}
