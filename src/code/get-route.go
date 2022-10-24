package code

import (
	"api-generator/src/utils"
	"path/filepath"
)

func generatorGetRoute(collectionName string) string {
	templatePath := filepath.Join(".", "templates", "src", "routes", "get-route.temp")
	code := utils.ParseTemplate(templatePath, collectionName)

	return code
}
