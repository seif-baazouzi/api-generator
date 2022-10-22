package utils

import (
	"api-generator/src/config"
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

func GenerateSqlPlaceholders(length int) []string {
	list := make([]string, length)

	for i := 0; i < length; i++ {
		list[i] = fmt.Sprintf("$%d", i+1)
	}

	return list
}

func GetCollectionFields(myMap config.Collection) []string {
	keys := make([]string, len(myMap))

	i := 0
	for key := range myMap {
		keys[i] = key
		i++
	}

	return keys
}
