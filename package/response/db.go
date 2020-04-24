package response

import (
	"strings"

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

func SQLResponseDatabaseError(err error) interface{} {
	reason := (strings.Split(err.Error(), "pq: "))[1]
	response := constants.ErrorResponseWithReason{20201, "Database error", reason}

	return response
}

func SQLResponseBadSQLStatement() interface{} {
	response := constants.ErrorResponse{20202, "Bad SQL statement"}

	return response
}
