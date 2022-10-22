package others

import (
	"api-generator/src/utils"
	"path/filepath"
)

func GeneratePackageJson(projectName string) {
	templatePath := filepath.Join(".", "templates", "package.json.temp")
	content := utils.ParseTemplate(templatePath, projectName)

	filePath := filepath.Join(".", "dist", "package.json")
	utils.CreateFile(filePath, content)
}
