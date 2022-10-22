package routes

import (
	"api-generator/src/config"
	"api-generator/src/utils"
	"path/filepath"
	"strings"
)

type postRouteDataStruct struct {
	CollectionName      string
	Columns             string
	ColumnsPlaceHolders string
}

func generatorPostRoute(collectionName string, collection config.Collection) string {
	data := postRouteDataStruct{
		CollectionName:      collectionName,
		Columns:             strings.Join(utils.GetCollectionFields(collection), ", "),
		ColumnsPlaceHolders: strings.Join(utils.GenerateSqlPlaceholders(len(collection)), ", "),
	}

	templatePath := filepath.Join(".", "templates", "post-route.temp")
	code := utils.ParseTemplate(templatePath, data)

	return code
}
