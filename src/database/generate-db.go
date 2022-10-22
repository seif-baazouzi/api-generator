package database

import (
	"api-generator/src/config"
	"api-generator/src/utils"
	"path/filepath"
)

func GenerateDB(config config.Config) {
	templatePath := filepath.Join(".", "templates", "db.temp")
	databaseSQL := utils.ParseTemplate(templatePath, config)

	filePath := filepath.Join(".", "dist", "db.sql")
	utils.CreateFile(filePath, databaseSQL)
}
