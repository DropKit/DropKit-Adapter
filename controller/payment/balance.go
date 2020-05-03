package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/DropKit/DropKit-Adapter/constants"
	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/ethereum/go-ethereum/common"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIPaymentBalance(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.TokenBalance
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIPaymentBalance(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}
	callerPriavteKey := newStatement.PrivateKey
	balanceAccount := newStatement.Account
	callerAddress, err := account.PrivateKeyToPublicKey(callerPriavteKey)
	if err != nil {
		services.NormalResponse(w, response.ResponsePKConvertError())
		return
	}

	switch balanceAccount {
	case "":
		balance, err := services.GetBalance(callerAddress)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		services.NormalResponse(w, response.PaymentBalanceResponseOk(balance.Int64()))
		logger.InfoAPIPaymentBalance(newStatement)
	default:
		result, err := services.HasDropKitAdmin(callerPriavteKey, callerAddress)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}

		switch result {
		case true:
			balance, err := services.GetBalance(common.HexToAddress(balanceAccount))
			if err != nil {
				services.NormalResponse(w, response.ResponseInternalError())
				return
			}

			services.NormalResponse(w, response.PaymentBalanceResponseOk(balance.Int64()))
			logger.InfoAPIPaymentBalance(balanceAccount)
		case false:
			services.NormalResponse(w, response.ResponseUnauthorized())
			logger.WarnAPIPaymentBurnUnAuth(callerAddress.String())
		}

	}
}
