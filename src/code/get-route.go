package code

import (
	"api-generator/src/utils"
	"path/filepath"
)

func generatorGetRoute(collectionName string) {
	templatePath := filepath.Join(".", "templates", "src", "routes", "get-route.temp")
	code := utils.ParseTemplate(templatePath, collectionName)

	filePath := filepath.Join(".", "dist", "src", collectionName, "get.js")
	utils.CreateFile(filePath, code)
}
