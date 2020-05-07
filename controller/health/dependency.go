package controller

import (
	"net/http"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/DropKit/DropKit-Adapter/package/response"
	"github.com/DropKit/DropKit-Adapter/services"
)

func CheckDependencyServices(w http.ResponseWriter, r *http.Request) {
	service, err := services.DependencyServicesCheck()
	if err != nil {
		logger.ErrorDependencyService(service, err)
		services.NormalResponse(w, response.ResponseDependencyError())
		return
	}

	services.NormalResponse(w, response.ResponseServerOk())
	logger.InfoAPIDenpendencyOk()
}
