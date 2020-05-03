package constants

type Permission struct {
	UserName   string `json:"user_name"`
	TableName  string `json:"table_name"`
	PrivateKey string `json:"caller_pk"`
}

type PermissionResponse struct {
	Code    int
	Message string
}

type PermissionVerifyResponse struct {
	Code     int
	Message  string
	Response bool
}

type SQL struct {
	Statement  string `json:"db_statement"`
	PrivateKey string `json:"caller_pk"`
}

type SQLExecResponse struct {
	Code    int
	Message string
	Audit   string
}

type SQLQueryResponse struct {
	Code    int
	Message string
	Data    interface{}
	Audit   string
}

type Token struct {
	Amount     int64  `json:"amount"`
	PrivateKey string `json:"caller_pk"`
}

type TokenTransfer struct {
	Account    string `json:"user_name"`
	Amount     int64  `json:"amount"`
	PrivateKey string `json:"caller_pk"`
}

type TokenTransferResponse struct {
	Code    int
	Message string
	Hash    string
}

type TokenTransferFailResponse struct {
	Code    int
	Message string
}

type TokenBalance struct {
	PrivateKey string `json:"caller_pk"`
	Account    string `json:"user_name"`
}

type TokenBalanceResponse struct {
	Code    int
	Message string
	Amount  int64
}

type ErrorResponse struct {
	Code    int
	Message string
}

type ErrorResponseWithReason struct {
	Code    int
	Message string
	Reason  string
}

type UserCreateResponse struct {
	Code       int
	Message    string
	PrivateKey string
	Account    string
}

type GeneralResponse struct {
	Code    int
	Message string
}
