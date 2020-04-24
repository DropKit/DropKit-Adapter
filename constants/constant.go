package constants

type Auth struct {
	UserName   string `json:"user_name"`
	TableName  string `json:"table_name"`
	PrivateKey string `json:"caller_pk"`
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

type AuthResponse struct {
	Code    int
	Message string
}

type AuthVerifyResponse struct {
	Code     int
	Message  string
	Response bool
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
