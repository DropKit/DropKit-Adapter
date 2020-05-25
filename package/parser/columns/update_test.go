package parser

import (
	"reflect"
	"testing"
)

func TestGetUpdateColumns(t *testing.T) {
	var tests = []struct {
		testName      string
		sqlCommand    string
		expectColumns []string
	}{
		{"DoubleQuoted", "UPDATE t set \"QWERTYUIOPASDFGHJKLZXCVBNM\" = `ABC` , \"qwertyuiopasdfghjklzxcvbnm\" = `abc` WHERE id = 1",
			[]string{"QWERTYUIOPASDFGHJKLZXCVBNM", "qwertyuiopasdfghjklzxcvbnm"}},
		{"NoQuoted", "UPDATE t set QWERTYUIOPASDFGHJKLZXCVBNM = `ABC` , qwertyuiopasdfghjklzxcvbnm = `abc` WHERE id = 1",
			[]string{"qwertyuiopasdfghjklzxcvbnm", "qwertyuiopasdfghjklzxcvbnm"}},
		{"Number", "UPDATE t set num1234567890 = -1 WHERE id = 1", []string{"num1234567890"}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, _ := GetUpdateColumns(tt.sqlCommand)
			if !reflect.DeepEqual(tt.expectColumns, result) {
				t.Errorf("got %v, expect %v", result, tt.expectColumns)
			}
		})
	}
}
