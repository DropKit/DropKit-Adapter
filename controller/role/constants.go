package role

type role struct {
	Columns     []string `json:"columns"`
	ColumnsName string   `json:"role_name"`
	PrivateKey  string   `json:"caller_pk"`
	TableName   string   `json:"table_name"`
}

type permissionResponse struct {
	Code    int
	Message string
}
