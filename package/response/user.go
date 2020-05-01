package response

import (
	"github.com/DropKit/DropKit-Adapter/constants"
)

func ResponseNewUser(privateKey string, account string) interface{} {
	response := constants.UserCreateResponse{0, "ok", privateKey, account}

	return response
}
