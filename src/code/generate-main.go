package code

import (
	"api-generator/src/config"
	"api-generator/src/utils"
	"path/filepath"
)

func generateMain(config config.Config) {
	templatePath := filepath.Join(".", "templates", "src", "main.temp")
	code := utils.ParseTemplate(templatePath, config)

	filePath := filepath.Join(".", "dist", "src", "server.js")
	utils.CreateFile(filePath, code)
}
