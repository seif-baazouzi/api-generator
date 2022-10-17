package utils

import "api-generator/src/config"

func GetCollectionFields(myMap config.Collection) []string {
	keys := make([]string, len(myMap))

	i := 0
	for key := range myMap {
		keys[i] = key
		i++
	}

	return keys
}
