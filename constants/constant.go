package constants

type Auth_Statement struct {
	UserName   string `json:"user_name"`
	TableName  string `json:"table_name"`
	PrivateKey string `json:"caller_pk"`
}

type DB_Statement struct {
	Statement  string `json:"db_statement"`
	PrivateKey string `json:"caller_pk"`
}

type Exec_Response struct {
	Result string
	Hash   string
}

type Query_Response struct {
	Result   string
	Response interface{}
	Hash     string
}

type Auth_Response struct {
	Result string
}

type Auth_Verify_Response struct {
	Result   string
	Response bool
}
