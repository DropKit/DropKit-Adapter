package controller

import (
	"net/http"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newPrivateKey, newAddress := account.GenerateWallet()

	services.NormalResponse(w, response.ResponseNewUser(newPrivateKey, newAddress))
	logger.InfoAPIUserCreate()
}
