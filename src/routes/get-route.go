package routes

import (
	"api-generator/src/utils"
	"path/filepath"
)

func generatorGetRoute(collectionName string) string {
	templatePath := filepath.Join(".", "templates", "get-route.temp")
	code := utils.ParseTemplate(templatePath, collectionName)

	return code
}
