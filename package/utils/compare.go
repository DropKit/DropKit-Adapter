package utils
import (
	"github.com/wxnacy/wgo/arrays"
)

func CompareColumns(columnsPermission []string, columnsRequest []string) bool {
	for _, column := range columnsRequest {
		position := arrays.ContainsString(columnsPermission, column)
		if position == -1 {
			return false
		}
	}
	return true
}