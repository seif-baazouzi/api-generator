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
}
