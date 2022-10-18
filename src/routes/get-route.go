package routes

import (
	"api-generator/src/utils"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func generatorGetRoute(collectionName string) string {
	templatePath := filepath.Join(".", "templates", "get-route.temp")
	template, err := template.ParseFiles(templatePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse template %s\n", templatePath)
		os.Exit(1)
	}

	code, err := utils.TemplateToString(template, collectionName)
	return code
}
