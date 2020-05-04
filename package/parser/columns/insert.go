package parser

import (
	"encoding/json"

	"github.com/DropKit/DropKit-Adapter/logger"
	pg_query "github.com/lfittl/pg_query_go"
)

type InsertCols struct {
	Col struct {
		Name     string `json:"name"`
		Location int    `json:"location"`
	} `json:"ResTarget"`
}

type InsertStatement struct {
	RawStmt struct {
		Stmt struct {
			InsertStmt struct {
				Cols []InsertCols `json:"cols"`
			} `json:"InsertStmt"`
		} `json:"stmt"`
	} `json:"RawStmt"`
}

func GetInsertColumns(sqlCommand string) ([]string, error) {
	SQLParserResult, err := pg_query.ParseToJSON(sqlCommand)
	if err != nil {
		logger.InternalLogger.WithField("component", "sql-parser").Warn(err.Error())
		return nil, err
	}
	SQLParserResultJSON := []byte(SQLParserResult)

	var statement []byte = SQLParserResultJSON[1:(len(SQLParserResultJSON) - 1)]

	var newStatement InsertStatement
	var columnNames []string
	json.Unmarshal(statement, &newStatement)
	for i := 0; i < len(newStatement.RawStmt.Stmt.InsertStmt.Cols); i++ {
		statementString, _ := json.Marshal(newStatement.RawStmt.Stmt.InsertStmt.Cols[i].Col.Name)
		columnNames = append(columnNames, string(statementString))
	}

	return columnNames, nil
}
