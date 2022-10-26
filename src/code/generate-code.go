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
	copyUtils()

	for collectionName, collection := range config.Collections {
		dirPath := filepath.Join(".", "dist", "src", collectionName)
		utils.CreateDirectory(dirPath)

		generatorGetRoute(collectionName, collection.Routes.Get)
		generatorGetSingleRoute(collectionName, collection.Routes.GetSingle)
		generatorPostRoute(collectionName, collection.Fields, collection.Routes.Post)
		generatorPutRoute(collectionName, collection.Fields, collection.Routes.Put)
		generatorDeleteRoute(collectionName, collection.Routes.Delete)
	}
}
