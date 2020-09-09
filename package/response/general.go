package response

func ResponseUnauthorized() interface{} {
	response := ErrorResponse{20101, "Unauthorized account"}

	return response
}

func ResponseBadRequest() interface{} {
	response := ErrorResponse{20102, "Bad input parameter"}

	return response
}

func ResponsePKConvertError() interface{} {
	response := ErrorResponse{20103, "caller_pk format invalid"}

	return response
}

func ResponseExceedsBalance() interface{} {
	response := ErrorResponse{20103, "Insufficient balance"}

	return response
}
