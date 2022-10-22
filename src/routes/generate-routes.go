package routes

import (
	"api-generator/src/config"
	"api-generator/src/utils"
	t "html/template"
	"path/filepath"
	"strings"
)

func GenerateRoutes(config config.Config) {
	routesCode := strings.Join([]string{}, "\n")

	for collectionName, collection := range config.Collections {
		routesCode += generatorCollectionRoutes(collectionName, collection)
	}

	templatePath := filepath.Join(".", "templates", "main.temp")
	code := utils.ParseTemplate(templatePath, t.HTML(routesCode))

	dirPath := filepath.Join(".", "dist", "src")
	utils.CreateDirectory(dirPath)

	filePath := filepath.Join(".", "dist", "src", "server.js")
	utils.CreateFile(filePath, code)
}

func generatorCollectionRoutes(collectionName string, collection config.Collection) string {
	return strings.Join([]string{
		generatorGetRoute(collectionName),
		generatorPostRoute(collectionName, collection),
		generatorPutRoute(collectionName, collection),
		generatorDeleteRoute(collectionName),
	}, "\n")
}
