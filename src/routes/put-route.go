package routes

import (
	"api-generator/src/config"
	"api-generator/src/utils"
	"path/filepath"
	"strings"
)

type putRouteDataStruct struct {
	CollectionName string
	Columns        string
	ColumnsSet     string
	IdPlaceholder  int
}

func generatorPutRoute(collectionName string, collection config.Collection) string {
	columns := utils.GetCollectionFields(collection)

	data := putRouteDataStruct{
		CollectionName: collectionName,
		Columns:        strings.Join(columns, ", "),
		ColumnsSet:     utils.GenerateColumnsSet(columns),
		IdPlaceholder:  len(collection) + 1,
	}

	templatePath := filepath.Join(".", "templates", "src", "routes", "put-route.temp")
	code := utils.ParseTemplate(templatePath, data)

	return code
}
