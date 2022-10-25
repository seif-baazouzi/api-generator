package code

import (
	"api-generator/src/utils"
	"path/filepath"
)

func copyUtils() {
	dirPath := filepath.Join(".", "dist", "src", "utils")
	utils.CreateDirectory(dirPath)

	utils.CopyFile(
		filepath.Join(".", "templates", "src", "utils", "try-catch.temp"),
		filepath.Join(".", "dist", "src", "utils", "try-catch.js"),
	)
}
