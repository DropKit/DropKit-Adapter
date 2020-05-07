package controller

import (
	"net/http"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
)

func PerformHealthCheck(w http.ResponseWriter, r *http.Request) {
	services.NormalResponse(w, response.ResponseServerOk())
	logger.InfoAPIPing()
}
