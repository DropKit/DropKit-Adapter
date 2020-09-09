package db

type sql struct {
	Statement  string `json:"db_statement"`
	PrivateKey string `json:"caller_pk"`
}

type sqlExecResponse struct {
	Code    int
	Message string
	Audit   string
}

type sqlQueryResponse struct {
	Code    int
	Message string
	Data    interface{}
	Audit   string
}
