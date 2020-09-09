package permission

type permission struct {
	UserName    string `json:"user_name"`
	TableName   string `json:"table_name"`
	PrivateKey  string `json:"caller_pk"`
	ColumnsRole string `json:"column_role"`
}

type permissionResponse struct {
	Code    int
	Message string
}

type permissionVerifyResponse struct {
	Code     int
	Message  string
	Response bool
}
