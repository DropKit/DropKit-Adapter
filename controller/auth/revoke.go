package controller

import (
	"encoding/json"
	"fmt"
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
		fmt.Println(err)
	}

	var newStatement constants.Auth
	_ = json.Unmarshal(body, &newStatement)
	logger.InfoAPIAuthorityRevoke(newStatement)
	callerPriavteKey := newStatement.PrivateKey
	revokeUser := newStatement.UserName
	revokeTable := newStatement.TableName
	callerAddress := account.PrivateKeyToPublicKey(callerPriavteKey)

	authority := services.VerifyAuthority(authorityAddr, callerPriavteKey, revokeTable, callerAddress)

	switch authority {
	case true:
		services.RevokeAuthority(authorityAddr, callerPriavteKey, revokeTable, revokeUser)
		defer r.Body.Close()
		services.ResponseWithJson(w, response.AuthResponseOk())
	case false:
		services.ResponseWithJson(w, response.AuthResponseUnauthorized())
	}
}
