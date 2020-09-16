package health

import (
	"net/http"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/gin-gonic/gin"
)

// PerformHealthCheck Response pong for checking server status
func PerformHealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "pong")
	logger.InfoAPIPing()
}
