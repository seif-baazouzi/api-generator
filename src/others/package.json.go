package others

import (
	"api-generator/src/utils"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func GeneratePackageJson(projectName string) {
	templatePath := filepath.Join(".", "templates", "package.json.temp")
	template, err := template.ParseFiles(templatePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse template %s\n", templatePath)
		os.Exit(1)
	}

	databaseSQL, err := utils.TemplateToString(template, nil)

	filePath := filepath.Join(".", "dist", "package.json")
	err = os.WriteFile(filePath, []byte(databaseSQL), os.ModePerm)
}
