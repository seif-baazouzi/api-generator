package code

import (
	"api-generator/src/utils"
	"path/filepath"
)

func generatorGetSingleRoute(collectionName string) {
	templatePath := filepath.Join(".", "templates", "src", "routes", "get-single-route.temp")
	code := utils.ParseTemplate(templatePath, collectionName)

	filePath := filepath.Join(".", "dist", "src", collectionName, "get-single.js")
	utils.CreateFile(filePath, code)
}
