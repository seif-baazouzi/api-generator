package utils

import "fmt"

func GenerateSqlPlaceholders(length int) []string {
	list := make([]string, length)

	for i := 0; i < length; i++ {
		list[i] = fmt.Sprintf("$%d", i+1)
	}

	return list
}
