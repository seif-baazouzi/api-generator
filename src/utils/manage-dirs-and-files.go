package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateDirectory(dirName string) string {
	path := filepath.Join(".", dirName)
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create `%s` directory\n", path)
		os.Exit(1)
	}

	return path
}

func RemoveDirectory(dirName string) string {
	path := filepath.Join(".", dirName)
	err := os.RemoveAll(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not remove `%s` directory\n", path)
		os.Exit(1)
	}

	return path
}

func ClearDirectory(dirName string) string {
	RemoveDirectory(dirName)
	path := CreateDirectory(dirName)

	return path
}

func CreateFile(filePath string, content string) {
	err := os.WriteFile(filePath, []byte(content), os.ModePerm)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create `%s` file\n", filePath)
		os.Exit(1)
	}
}
