package others

import (
	"api-generator/src/utils"
	"path/filepath"
)

func GenerateDockerFiles(projectName string) {
	// db docker file
	utils.CopyFile(
		filepath.Join(".", "templates", "database", "Dockerfile"),
		filepath.Join(".", "dist", "database", "Dockerfile"),
	)

	// node docker file
	utils.CopyFile(
		filepath.Join(".", "templates", "Dockerfile"),
		filepath.Join(".", "dist", "Dockerfile"),
	)

	// docker compose
	templatePath := filepath.Join(".", "templates", "docker-compose.yml")
	databaseSQL := utils.ParseTemplate(templatePath, projectName)

	filePath := filepath.Join(".", "dist", "docker-compose.yml")
	utils.CreateFile(filePath, databaseSQL)
}
