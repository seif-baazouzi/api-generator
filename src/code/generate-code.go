package code

import (
	"api-generator/src/config"
	"api-generator/src/utils"
	"path/filepath"
)

func GenerateCode(config config.Config) {
	dirPath := filepath.Join(".", "dist", "src")
	utils.CreateDirectory(dirPath)

	generateDB()
	generateMain(config)
	generateValidation(config)

	for collectionName, collection := range config.Collections {
		dirPath := filepath.Join(".", "dist", "src", collectionName)
		utils.CreateDirectory(dirPath)

		generatorGetRoute(collectionName)
		generatorPostRoute(collectionName, collection)
		generatorPutRoute(collectionName, collection)
		generatorDeleteRoute(collectionName)
	}
}
