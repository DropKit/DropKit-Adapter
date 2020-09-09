package user

import (
	"net/http"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/gin-gonic/gin"
)

func GenerateRandomAccount(c *gin.Context) {
	newPrivateKey, newAddress := account.GenerateWallet()

	c.JSON(http.StatusOK, userCreateResponse{0, "ok", newPrivateKey, newAddress})
	logger.InfoAPIUserCreate()
}
