package parser

import (
	"encoding/json"

	"github.com/DropKit/DropKit-Adapter/logger"
	pg_query "github.com/lfittl/pg_query_go"
)

type CreateCols struct {
	Col struct {
		Name string `json:"colname"`
	} `json:"ColumnDef"`
}

type CreateStatement struct {
	RawStmt struct {
		Stmt struct {
			CreateStmt struct {
				TableElts []CreateCols `json:"tableElts"`
			} `json:"CreateStmt"`
		} `json:"stmt"`
	} `json:"RawStmt"`
}

func GetCreateColumns(sqlCommand string) ([]string, error) {
	SQLParserResult, err := pg_query.ParseToJSON(sqlCommand)
	if err != nil {
		logger.InternalLogger.WithField("component", "sql-parser").Warn(err.Error())
		return nil, err
	}
	SQLParserResultJSON := []byte(SQLParserResult)

	var statement []byte = SQLParserResultJSON[1:(len(SQLParserResultJSON) - 1)]

	var newStatement CreateStatement
	var columnNames []string
	json.Unmarshal(statement, &newStatement)
	for i := 0; i < len(newStatement.RawStmt.Stmt.CreateStmt.TableElts); i++ {
		statementString, _ := json.Marshal(newStatement.RawStmt.Stmt.CreateStmt.TableElts[i].Col.Name)
		columnNames = append(columnNames, string(statementString[1:len(statementString)-1]))
	}

	return columnNames, nil
}
