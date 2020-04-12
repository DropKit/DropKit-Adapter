package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	"github.com/DropKit/DropKit-Adapter/constants"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/spf13/viper"
)

func AuthGrant(w http.ResponseWriter, r *http.Request) {
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var newStatement constants.Auth_Statement
	_ = json.Unmarshal(body, &newStatement)
	callerPriavteKey := newStatement.PrivateKey
	grantUser := newStatement.UserName
	grantTable := newStatement.TableName
	callerAddress := account.PrivateKeyToPublicKey(callerPriavteKey)

	authority := services.VerifyAuthority(authorityAddr, callerPriavteKey, grantTable, callerAddress)

	switch authority {
	case true:
		services.GrantAuthority(authorityAddr, callerPriavteKey, grantTable, grantUser)

		defer r.Body.Close()
		response := constants.Auth_Response{"200"}
		services.ResponseWithJson(w, http.StatusOK, response)
	case false:
		response := constants.Auth_Response{"401"}
		services.ResponseWithJson(w, http.StatusUnauthorized, response)
	}
}

func AuthRevoke(w http.ResponseWriter, r *http.Request) {
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var newStatement constants.Auth_Statement
	_ = json.Unmarshal(body, &newStatement)
	callerPriavteKey := newStatement.PrivateKey
	revokeUser := newStatement.UserName
	revokeTable := newStatement.TableName
	callerAddress := account.PrivateKeyToPublicKey(callerPriavteKey)

	authority := services.VerifyAuthority(authorityAddr, callerPriavteKey, revokeTable, callerAddress)

	switch authority {
	case true:
		services.RevokeAuthority(authorityAddr, callerPriavteKey, revokeTable, revokeUser)

		defer r.Body.Close()
		response := constants.Auth_Response{"200"}
		services.ResponseWithJson(w, http.StatusOK, response)
	case false:
		response := constants.Auth_Response{"401"}
		services.ResponseWithJson(w, http.StatusUnauthorized, response)
	}

}

func AuthVerify(w http.ResponseWriter, r *http.Request) {
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var newStatement constants.Auth_Statement
	_ = json.Unmarshal(body, &newStatement)
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
			response := constants.Auth_Verify_Response{"200", true}
			services.ResponseWithJson(w, http.StatusOK, response)
		case false:
			defer r.Body.Close()
			response := constants.Auth_Verify_Response{"200", false}
			services.ResponseWithJson(w, http.StatusOK, response)
		}
	case false:
		response := constants.Auth_Response{"401"}
		services.ResponseWithJson(w, http.StatusUnauthorized, response)
	}
}
