package parser

import (
	"reflect"
	"testing"
)

func TestGetSelectColumns(t *testing.T) {
	var tests = []struct {
		testName      string
		sqlCommand    string
		expectColumns []string
	}{
		{"Asterisk", "SELECT * FROM t WHERE id = 1", []string{""}},
		{"DoubleQuoted", "SELECT \"QWERTYUIOPASDFGHJKLZXCVBNM\", \"qwertyuiopasdfghjklzxcvbnm\" FROM t WHERE id = 1",
			[]string{"QWERTYUIOPASDFGHJKLZXCVBNM", "qwertyuiopasdfghjklzxcvbnm"}},
		{"NoQuoted", "SELECT QWERTYUIOPASDFGHJKLZXCVBNM, qwertyuiopasdfghjklzxcvbnm FROM t WHERE id = 1",
			[]string{"qwertyuiopasdfghjklzxcvbnm", "qwertyuiopasdfghjklzxcvbnm"}},
		{"Number", "SELECT num1234567890 FROM t WHERE id = 1", []string{"num1234567890"}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, _ := GetSelectColumns(tt.sqlCommand)
			if !reflect.DeepEqual(tt.expectColumns, result) {
				t.Errorf("got %v, expect %v", result, tt.expectColumns)
			}
		})
	}
}
