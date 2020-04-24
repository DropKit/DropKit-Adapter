package response

import (
	"github.com/DropKit/DropKit-Adapter/constants"
)

func ResponseUnauthorized() interface{} {
	response := constants.ErrorResponse{20101, "Unauthorized account"}

	return response
}

func ResponseBadRequest() interface{} {
	response := constants.ErrorResponse{20102, "Bad input parameter"}

	return response
}

func ResponsePKConvertError() interface{} {
	response := constants.ErrorResponse{20103, "caller_pk format is not correct"}

	return response
}
