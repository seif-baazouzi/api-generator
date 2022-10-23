package database

import (
	"api-generator/src/config"
	"api-generator/src/utils"
	"path/filepath"
)

func GenerateDB(config config.Config) {
	templatePath := filepath.Join(".", "templates", "database", "db.temp")
	databaseSQL := utils.ParseTemplate(templatePath, config)

	dirPath := filepath.Join(".", "dist", "database")
	utils.CreateDirectory(dirPath)

	filePath := filepath.Join(".", "dist", "database", "db.sql")
	utils.CreateFile(filePath, databaseSQL)
}
