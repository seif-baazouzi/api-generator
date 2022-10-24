package code

import (
	"api-generator/src/utils"
	"path/filepath"
)

func generatorDeleteRoute(collectionName string) {
	templatePath := filepath.Join(".", "templates", "src", "routes", "delete-route.temp")
	code := utils.ParseTemplate(templatePath, collectionName)

	filePath := filepath.Join(".", "dist", "src", collectionName, "delete.js")
	utils.CreateFile(filePath, code)
}
