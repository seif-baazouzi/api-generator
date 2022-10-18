package routes

import (
	"api-generator/src/config"
	"api-generator/src/utils"
	"fmt"
	"html/template"
	"os"
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
	template, err := template.ParseFiles(templatePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse template %s\n", templatePath)
		os.Exit(1)
	}

	code, err := utils.TemplateToString(template, data)
	return code
}
