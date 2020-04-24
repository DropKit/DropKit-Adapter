package parser

import (
	"strings"

	"github.com/DropKit/DropKit-Adapter/logger"
	pg_query "github.com/lfittl/pg_query_go"
)

func GetTableName(sqlCommand string) (string, error) {
	rawStatement, err := pg_query.ParseToJSON(sqlCommand)
	if err != nil {
		logger.InternalLogger.WithField("component", "sql-parser").Warn(err.Error())
		return "", err
	}
	rawData := (strings.Split(rawStatement, "relname\": \""))[1]
	tableName := (strings.Split(rawData, "\""))[0]

	return tableName, nil
}
