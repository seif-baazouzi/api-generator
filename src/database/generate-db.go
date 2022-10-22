package database

import (
	"api-generator/src/config"
	"api-generator/src/utils"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func GenerateDB(config config.Config) {
	templatePath := filepath.Join(".", "templates", "db.temp")
	template, err := template.ParseFiles(templatePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse template %s\n", templatePath)
		os.Exit(1)
	}

	databaseSQL, err := utils.TemplateToString(template, config)

	filePath := filepath.Join(".", "dist", "db.sql")
	utils.CreateFile(filePath, databaseSQL)
}
