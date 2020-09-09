package response

func ResponseInternalError() interface{} {
	response := ErrorResponse{10101, "internal server error"}

	return response
}

func ResponseDependencyError() interface{} {
	response := ErrorResponse{10102, "dependency service error"}

	return response
}

func ResponseDependencyOk() interface{} {
	response := ErrorResponse{0, "Ok"}

	return response
}

func ResponseServerOk() interface{} {
	response := GeneralResponse{0, "Pong"}

	return response
}
