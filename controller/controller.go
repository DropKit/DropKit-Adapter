package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	"github.com/DropKit/DropKit-Adapter/constants"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/DropKit/DropKit-Adapter/package/crypto/transaction"
	"github.com/DropKit/DropKit-Adapter/package/parser"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/spf13/viper"
)

func SQLCreate(w http.ResponseWriter, r *http.Request) {
	metaTableAddress := viper.GetString(`DROPKIT.METATABLE`)
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var newStatement constants.DB_Statement
	_ = json.Unmarshal(body, &newStatement)
	sqlCommand := newStatement.Statement
	callerPriavteKey := newStatement.PrivateKey
	callerAddress := account.PrivateKeyToPublicKey(callerPriavteKey)

	tableName := parser.GetTableName(sqlCommand)
	_, tableAddress := account.GenerateWallet()

	services.AddMetaTable(tableName, tableAddress, metaTableAddress, callerPriavteKey)
	services.GrantAuthority(authorityAddr, callerPriavteKey, tableName, callerAddress)
	services.Exec(sqlCommand)

	aduitTransactionHash := transaction.SendRawTransaction(tableAddress, sqlCommand, 0, callerPriavteKey)

	defer r.Body.Close()
	response := constants.Exec_Response{"200", aduitTransactionHash}

	services.ResponseWithJson(w, http.StatusOK, response)
}

func SQLInsert(w http.ResponseWriter, r *http.Request) {
	metaTableAddress := viper.GetString(`DROPKIT.METATABLE`)
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var newStatement constants.DB_Statement
	_ = json.Unmarshal(body, &newStatement)
	sqlCommand := newStatement.Statement
	callerPriavteKey := newStatement.PrivateKey
	callerAddress := account.PrivateKeyToPublicKey(callerPriavteKey)

	tableName := parser.GetTableName(sqlCommand)
	tableAddress := services.GetMetaTable(tableName, metaTableAddress)
	authority := services.VerifyAuthority(authorityAddr, callerPriavteKey, tableName, callerAddress)

	switch authority {
	case true:
		services.Exec(sqlCommand)
		aduitTransactionHash := transaction.SendRawTransaction(tableAddress, sqlCommand, 0, callerPriavteKey)
		defer r.Body.Close()
		response := constants.Exec_Response{"200", aduitTransactionHash}
		services.ResponseWithJson(w, http.StatusOK, response)
	case false:
		defer r.Body.Close()
		response := constants.Exec_Response{"401", "NULL"}
		services.ResponseWithJson(w, http.StatusUnauthorized, response)
	}
}

func SQLSelect(w http.ResponseWriter, r *http.Request) {
	metaTableAddress := viper.GetString(`DROPKIT.METATABLE`)
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var newStatement constants.DB_Statement
	_ = json.Unmarshal(body, &newStatement)
	sqlCommand := newStatement.Statement
	callerPriavteKey := newStatement.PrivateKey
	callerAddress := account.PrivateKeyToPublicKey(callerPriavteKey)

	tableName := parser.GetTableName(sqlCommand)
	tableAddress := services.GetMetaTable(tableName, metaTableAddress)
	authority := services.VerifyAuthority(authorityAddr, callerPriavteKey, tableName, callerAddress)

	switch authority {
	case true:
		metadata, _ := services.Query(sqlCommand)
		aduitTransactionHash := transaction.SendRawTransaction(tableAddress, sqlCommand, 0, callerPriavteKey)
		defer r.Body.Close()
		response := constants.Query_Response{"200", metadata, aduitTransactionHash}
		services.ResponseWithJson(w, http.StatusOK, response)
	case false:
		defer r.Body.Close()
		response := constants.Query_Response{"401", "NULL", "NULL"}
		services.ResponseWithJson(w, http.StatusUnauthorized, response)
	}
}

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

	services.GrantAuthority(authorityAddr, callerPriavteKey, grantTable, grantUser)

	defer r.Body.Close()
	response := constants.Auth_Response{"200"}
	services.ResponseWithJson(w, http.StatusOK, response)
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

	services.RevokeAuthority(authorityAddr, callerPriavteKey, revokeTable, revokeUser)

	defer r.Body.Close()
	response := constants.Auth_Response{"200"}
	services.ResponseWithJson(w, http.StatusOK, response)
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

	authority := services.VerifyAuthority(authorityAddr, callerPriavteKey, checkTable, checkUser)

	switch authority {
	case true:
		defer r.Body.Close()
		response := constants.Auth_Response{"200"}
		services.ResponseWithJson(w, http.StatusOK, response)
	case false:
		defer r.Body.Close()
		response := constants.Auth_Response{"401"}
		services.ResponseWithJson(w, http.StatusUnauthorized, response)
	}
}
