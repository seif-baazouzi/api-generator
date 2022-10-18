package utils

import (
	"fmt"
	"strings"
)

func GenerateColumnsSet(columns []string) string {
	columnsSet := make([]string, len(columns))

	for i, column := range columns {
		columnsSet[i] = fmt.Sprintf("%s = $%d", column, i+1)
	}

	return strings.Join(columnsSet, ", ")
}
