package parser

import (
	"reflect"
	"testing"
)

func TestGetCreateColumns(t *testing.T) {
	var tests = []struct {
		testName      string
		sqlCommand    string
		expectColumns []string
	}{
		{"Case", "CREATE TABLE t (QWERTYUIOPASDFGHJKLZXCVBNM INT PRIMARY KEY, qwertyuiopasdfghjklzxcvbnm INT);",
			[]string{"QWERTYUIOPASDFGHJKLZXCVBNM", "qwertyuiopasdfghjklzxcvbnm"}},
		{"Number", "CREATE TABLE t (num1234567890 INT PRIMARY KEY);", []string{"num1234567890"}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, _ := GetCreateColumns(tt.sqlCommand)
			if !reflect.DeepEqual(tt.expectColumns, result) {
				t.Errorf("got %v, expect %v", result, tt.expectColumns)
			}
		})
	}
}
