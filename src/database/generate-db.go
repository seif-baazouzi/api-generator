package database

import (
	"api-generator/src/config"
	"fmt"
	"os"
	"path/filepath"
)

func GenerateDB(config config.Config) {
	databaseCode := "CREATE DATABASE api;\n"
	databaseCode += "\\c api;\n"

	for collectionName, collection := range config.Collections {
		databaseCode += createCollection(collectionName, collection)
	}

	filePath := filepath.Join(".", "dist", "db.sql")
	err := os.WriteFile(filePath, []byte(databaseCode), os.ModePerm)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create %s file\n", filePath)
		os.Exit(1)
	}
}

func createCollection(collectionName string, collection config.Collection) string {
	collectionCode := ""

	collectionCode += fmt.Sprintf("CREATE TABLE %s (\n", collectionName)
	collectionCode += "    id UUID NOT NULL,\n"

	for fieldName, field := range collection {
		collectionCode += createField(fieldName, field)
	}

	collectionCode += "    PRIMARY KEY (id)\n"
	collectionCode += ");\n"

	return collectionCode
}

func createField(fieldName string, field config.Field) string {
	var isNull string
	if field.Required {
		isNull = "NOT NULL"
	} else {
		isNull = "NULL"
	}

	return fmt.Sprintf("    %s %s %s,\n", fieldName, field.Type, isNull)
}
