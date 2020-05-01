package response

import (
	"github.com/DropKit/DropKit-Adapter/constants"
)

func PermissionResponseOk() interface{} {
	response := constants.PermissionResponse{0, "Ok"}

	return response
}

func PermissionVerifyResponse(result bool) interface{} {
	response := constants.PermissionVerifyResponse{0, "Ok", result}

	return response
}
