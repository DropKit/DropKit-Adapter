package response

type ErrorResponse struct {
	Code    int
	Message string
}

type ErrorResponseWithReason struct {
	Code    int
	Message string
	Reason  string
}

type GeneralResponse struct {
	Code    int
	Message string
}
