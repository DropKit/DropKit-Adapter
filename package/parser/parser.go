package parser

import (
	"strings"

	pg_query "github.com/lfittl/pg_query_go"
)

func GetTableName(sqlCommand string) string {
	rawStatement, _ := pg_query.ParseToJSON(sqlCommand)
	rawData := (strings.Split(rawStatement, "relname\": \""))[1]
	tableName := (strings.Split(rawData, "\""))[0]

	return tableName
}
