package code

import (
	"api-generator/src/utils"
	"path/filepath"
)

func generateDB() {
	utils.CopyFile(
		filepath.Join(".", "templates", "src", "db.temp"),
		filepath.Join(".", "dist", "src", "db.js"),
	)
}
