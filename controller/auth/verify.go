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

func AuthVerify(w http.ResponseWriter, r *http.Request) {
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WarnAPIAuthorityVerify(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}

	var newStatement constants.Auth
	err = json.Unmarshal(body, &newStatement)
	if err != nil {
		logger.WarnAPIAuthorityVerify(err)
		services.NormalResponse(w, response.ResponseBadRequest())
		return
	}
	logger.InfoAPIAuthorityVerify(newStatement)

	callerPriavteKey := newStatement.PrivateKey
	checkUser := newStatement.UserName
	checkTable := newStatement.TableName
	callerAddress, err := account.PrivateKeyToPublicKey(callerPriavteKey)
	if err != nil {
		services.NormalResponse(w, response.ResponsePKConvertError())
		return
	}

	callerAuthority, _ := services.VerifyAuthority(authorityAddr, callerPriavteKey, checkTable, callerAddress)

	switch callerAuthority {
	case true:
		authority, _ := services.VerifyAuthority(authorityAddr, callerPriavteKey, checkTable, checkUser)

		switch authority {
		case true:
			defer r.Body.Close()
			services.NormalResponse(w, response.AuthVerifyResponse(true))
		case false:
			defer r.Body.Close()
			services.NormalResponse(w, response.AuthVerifyResponse(false))
		}
	case false:
		services.NormalResponse(w, response.ResponseUnauthorized())
	}
}
