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
)

func BurnToken(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIPaymentBurn(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	if body != nil {
		defer r.Body.Close()
	}

	var newStatement constants.Token
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIPaymentBurn(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	callerPriavteKey := newStatement.PrivateKey
	amount := newStatement.Amount
	callerAddress, err := account.PrivateKeyToPublicKey(callerPriavteKey)
	if err != nil {
		services.NormalResponse(w, response.ResponsePKConvertError())
		return
	}

	result, err := services.HasDropKitAdmin(callerPriavteKey, callerAddress)
	if err != nil {
		services.NormalResponse(w, response.ResponseInternalError())
		return
	}

	switch result {
	case true:
		hash, err := services.MintToken(amount, callerPriavteKey, callerAddress)
		if err != nil {
			services.NormalResponse(w, response.ResponseInternalError())
			return
		}
		services.NormalResponse(w, response.PaymentResponseOk(hash))
		logger.InfoAPIPaymentBurn(newStatement)
	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
		logger.WarnAPIPaymentBurnUnAuth(callerAddress.String())
	}
}
