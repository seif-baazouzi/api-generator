package code

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

func generatorPutRoute(collectionName string, fields config.Fields, generate bool) {
	if !generate {
		return
	}

	columns := utils.GetCollectionFields(fields)

	data := putRouteDataStruct{
		CollectionName: collectionName,
		Columns:        strings.Join(columns, ", "),
		ColumnsSet:     utils.GenerateColumnsSet(columns),
		IdPlaceholder:  len(fields) + 1,
	}

	templatePath := filepath.Join(".", "templates", "src", "routes", "put-route.temp")
	code := utils.ParseTemplate(templatePath, data)

	filePath := filepath.Join(".", "dist", "src", collectionName, "put.js")
	utils.CreateFile(filePath, code)
}
