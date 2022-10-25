package code

import (
	"api-generator/src/config"
	"api-generator/src/utils"
	"path/filepath"
	"strings"
)

type validationDataStruct struct {
	Params string
	Fields config.Collection
}

func generateValidation(config config.Config) {
	dirPath := filepath.Join(".", "dist", "src", "validation")
	utils.CreateDirectory(dirPath)

	for collectionName, collection := range config.Collections {
		data := validationDataStruct{
			Params: strings.Join(utils.GetCollectionFields(collection), ", "),
			Fields: collection,
		}

		templatePath := filepath.Join(".", "templates", "src", "validation.temp")
		code := utils.ParseTemplate(templatePath, data)

		filePath := filepath.Join(".", "dist", "src", "validation", collectionName+".js")
		utils.CreateFile(filePath, code)
	}
}
