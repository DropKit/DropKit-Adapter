package response

import (
	"github.com/DropKit/DropKit-Adapter/constants"
)

func PaymentResponseOk(transacionHash string) interface{} {
	response := constants.TokenTransferResponse{0, "Ok", transacionHash}

	return response
}

func PaymentBalanceResponseOk(amount int64) interface{} {
	response := constants.TokenBalanceResponse{0, "Ok", amount}

	return response
}

func PaymentResponseNotEnough() interface{} {
	response := constants.TokenTransferFailResponse{20401, "not enough balance"}

	return response
}
