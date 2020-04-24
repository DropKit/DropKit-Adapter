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
	"github.com/spf13/viper"
)

func AuthRevoke(w http.ResponseWriter, r *http.Request) {
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIAuthorityRevoke(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	var newStatement constants.Auth
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIAuthorityRevoke(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}
	logger.InfoAPIAuthorityRevoke(newStatement)

	callerPriavteKey := newStatement.PrivateKey
	revokeUser := newStatement.UserName
	revokeTable := newStatement.TableName
	callerAddress, err := account.PrivateKeyToPublicKey(callerPriavteKey)
	if err != nil {
		services.NormalResponse(w, response.ResponsePKConvertError())
		return
	}

	authority, _ := services.VerifyAuthority(authorityAddr, callerPriavteKey, revokeTable, callerAddress)

	switch authority {
	case true:
		services.RevokeAuthority(authorityAddr, callerPriavteKey, revokeTable, revokeUser)
		defer r.Body.Close()
		services.NormalResponse(w, response.AuthResponseOk())
	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
	}
}
