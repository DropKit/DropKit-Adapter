package parser

import (
	"reflect"
	"testing"
)

func TestGetInsertColumns(t *testing.T) {
	var tests = []struct {
		testName      string
		sqlCommand    string
		expectColumns []string
	}{
		{"DoubleQuoted", "INSERT INTO t (\"QWERTYUIOPASDFGHJKLZXCVBNM\", \"qwertyuiopasdfghjklzxcvbnm\") VALUES(`aaa`,`bbb`)",
			[]string{"QWERTYUIOPASDFGHJKLZXCVBNM", "qwertyuiopasdfghjklzxcvbnm"}},
		{"NoQuoted", "INSERT INTO t (QWERTYUIOPASDFGHJKLZXCVBNM, qwertyuiopasdfghjklzxcvbnm) VALUES(`aaa`,`bbb`)",
			[]string{"qwertyuiopasdfghjklzxcvbnm", "qwertyuiopasdfghjklzxcvbnm"}},
		{"Number", "INSERT INTO t (num1234567890) VALUES (1)", []string{"num1234567890"}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, _ := GetInsertColumns(tt.sqlCommand)
			if !reflect.DeepEqual(tt.expectColumns, result) {
				t.Errorf("got %v, expect %v", result, tt.expectColumns)
			}
		})
	}
}
