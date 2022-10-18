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

	templatePath := filepath.Join(".", "templates", "put-route.temp")
	template, err := template.ParseFiles(templatePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse template %s\n", templatePath)
		os.Exit(1)
	}

	code, err := utils.TemplateToString(template, data)
	return code
}
