package parser

import (
	"encoding/json"

	"github.com/DropKit/DropKit-Adapter/logger"
	pg_query "github.com/lfittl/pg_query_go"
)

type UpdateCols struct {
	Col struct {
		Name string `json:"name"`
	} `json:"ResTarget"`
}

type UpdateColsStatement struct {
	RawStmt struct {
		Stmt struct {
			UpdateStmt struct {
				Cols []UpdateCols `json:"targetList"`
			} `json:"UpdateStmt"`
		} `json:"stmt"`
	} `json:"RawStmt"`
}

func GetUpdateColumns(sqlCommand string) ([]string, error) {
	SQLParserResult, err := pg_query.ParseToJSON(sqlCommand)
	if err != nil {
		logger.InternalLogger.WithField("component", "sql-parser").Warn(err.Error())
		return nil, err
	}

	SQLParserResultJSON := []byte(SQLParserResult)

	var statement []byte = SQLParserResultJSON[1:(len(SQLParserResultJSON) - 1)]

	var newStatement UpdateColsStatement
	var columnNames []string
	json.Unmarshal(statement, &newStatement)
	for i := 0; i < len(newStatement.RawStmt.Stmt.UpdateStmt.Cols); i++ {
		statementString, _ := json.Marshal(newStatement.RawStmt.Stmt.UpdateStmt.Cols[i].Col.Name)
		columnNames = append(columnNames, string(statementString[1:len(statementString)-1]))
	}

	return columnNames, nil
}
