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

func AuthGrant(w http.ResponseWriter, r *http.Request) {
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIAuthorityGrant(err)
	}

	var newStatement constants.Auth
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIAuthorityGrant(err)
	}
	logger.InfoAPIAuthorityGrant(newStatement)

	callerPriavteKey := newStatement.PrivateKey
	grantUser := newStatement.UserName
	grantTable := newStatement.TableName
	callerAddress := account.PrivateKeyToPublicKey(callerPriavteKey)

	authority := services.VerifyAuthority(authorityAddr, callerPriavteKey, grantTable, callerAddress)

	switch authority {
	case true:
		services.GrantAuthority(authorityAddr, callerPriavteKey, grantTable, grantUser)
		defer r.Body.Close()
		services.NormalResponse(w, response.AuthResponseOk())
	case false:
		services.NormalResponse(w, response.AuthResponseUnauthorized())
	}
}
