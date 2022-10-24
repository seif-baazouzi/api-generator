package code

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
	IdPlaceholder       int
}

func generatorPostRoute(collectionName string, collection config.Collection) {
	data := postRouteDataStruct{
		CollectionName:      collectionName,
		Columns:             strings.Join(utils.GetCollectionFields(collection), ", "),
		ColumnsPlaceHolders: strings.Join(utils.GenerateSqlPlaceholders(len(collection)), ", "),
		IdPlaceholder:       len(collection) + 1,
	}

	templatePath := filepath.Join(".", "templates", "src", "routes", "post-route.temp")
	code := utils.ParseTemplate(templatePath, data)

	filePath := filepath.Join(".", "dist", "src", collectionName, "post.js")
	utils.CreateFile(filePath, code)
}
