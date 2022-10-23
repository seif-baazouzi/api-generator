package routes

import (
	"api-generator/src/utils"
	"path/filepath"
)

func generatorDeleteRoute(collectionName string) string {
	templatePath := filepath.Join(".", "templates", "src", "routes", "delete-route.temp")
	code := utils.ParseTemplate(templatePath, collectionName)

	return code
}
