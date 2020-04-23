package response

import (
	"github.com/DropKit/DropKit-Adapter/constants"
)

func SQLExecResponseOk(aduitTransactionHash string) interface{} {
	response := constants.SQLExecResponse{0, "Ok", aduitTransactionHash}

	return response
}

func SQLQueryResponseOk(metadata interface{}, aduitTransactionHash string) interface{} {
	response := constants.SQLQueryResponse{0, "Ok", metadata, aduitTransactionHash}

	return response
}

func SQLResponseUnauthorized() interface{} {
	response := constants.ErrorResponse{20101, "Unauthorized"}

	return response
}
