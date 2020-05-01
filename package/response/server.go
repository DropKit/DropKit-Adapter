package response

import (
	"github.com/DropKit/DropKit-Adapter/constants"
)

func ResponseInternalError() interface{} {
	response := constants.ErrorResponse{10101, "internal server error"}

	return response
}

func ResponseDependencyError() interface{} {
	response := constants.ErrorResponse{10102, "dependency service error"}

	return response
}

func ResponseDependencyOk() interface{} {
	response := constants.ErrorResponse{0, "Ok"}

	return response
}

func ResponseServerOk() interface{} {
	response := constants.GeneralResponse{0, "Pong"}

	return response
}
