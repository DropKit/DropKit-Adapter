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

func AuthVerify(w http.ResponseWriter, r *http.Request) {
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var newStatement constants.Auth
	_ = json.Unmarshal(body, &newStatement)
	logger.InfoAPIAuthorityVerify(newStatement)
	callerPriavteKey := newStatement.PrivateKey
	checkUser := newStatement.UserName
	checkTable := newStatement.TableName
	callerAddress := account.PrivateKeyToPublicKey(callerPriavteKey)

	callerAuthority := services.VerifyAuthority(authorityAddr, callerPriavteKey, checkTable, callerAddress)

	switch callerAuthority {
	case true:
		authority := services.VerifyAuthority(authorityAddr, callerPriavteKey, checkTable, checkUser)

		switch authority {
		case true:
			defer r.Body.Close()
			services.ResponseWithJson(w, response.AuthVerifyResponse(true))
		case false:
			defer r.Body.Close()
			services.ResponseWithJson(w, response.AuthVerifyResponse(false))
		}
	case false:
		services.ResponseWithJson(w, response.AuthResponseUnauthorized())
	}
}
