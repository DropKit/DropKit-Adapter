package parser

import (
	"encoding/json"

	"github.com/DropKit/DropKit-Adapter/logger"
	pg_query "github.com/lfittl/pg_query_go"
)

type ResTarget struct {
	ResTarget struct {
		Val struct {
			ColumnRef struct {
				Fields []Fields `fields`
			} `json:"ColumnRef"`
		} `json:"val"`
	} `json:"ResTarget"`
}

type Fields struct {
	Col struct {
		Name string `json:"str"`
	} `json:"String"`
}

type SelectStatement struct {
	RawStmt struct {
		Stmt struct {
			SelectStmt struct {
				TargetList []ResTarget `json:"targetList"`
			} `json:"SelectStmt"`
		} `json:"stmt"`
	} `json:"RawStmt"`
}

func GetSelectColumns(sqlCommand string) ([]string, error) {
	SQLParserResult, err := pg_query.ParseToJSON(sqlCommand)
	if err != nil {
		logger.InternalLogger.WithField("component", "sql-parser").Warn(err.Error())
		return nil, err
	}

	SQLParserResultJSON := []byte(SQLParserResult)

	var statement []byte = SQLParserResultJSON[1:(len(SQLParserResultJSON) - 1)]

	var newStatement SelectStatement
	var columnNames []string
	json.Unmarshal(statement, &newStatement)
	for i := 0; i < len(newStatement.RawStmt.Stmt.SelectStmt.TargetList); i++ {
		statementString, _ := json.Marshal(newStatement.RawStmt.Stmt.SelectStmt.TargetList[i].ResTarget.Val.ColumnRef.Fields[0].Col.Name)
		columnNames = append(columnNames, string(statementString[1:len(statementString)-1]))
	}

	return columnNames, nil
}
