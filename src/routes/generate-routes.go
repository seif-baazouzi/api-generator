package routes

import (
	"api-generator/src/config"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateRoutes(config config.Config) {
	routesCode := strings.Join([]string{
		`require("dotenv").config()`,
		``,
		`const express = require("express")`,
		`const app = express()`,
		`const cors = require("cors")`,
		`const { Pool } = require("pg")`,
		``,
		`app.use(cors())`,
		`app.use(require("body-parser").json())`,
		``,
		`const db = new Pool({`,
		`	user: process.env.PG_USER,`,
		`	password: process.env.PG_PASS,`,
		`	host: process.env.PG_HOST,`,
		`	database: process.env.PG_DB,`,
		`	max: 10,`,
		`})`,
		``,
		`db.on("error", (err, client) => {`,
		`	console.error("Unexpected error on idle client", err)`,
		`	process.exit(-1)`,
		`})`,
		``,
		`console.log("Connected to postgresql database")`,
		``,
	}, "\n")

	for collectionName, collection := range config.Collections {
		routesCode += generatorCollectionRoutes(collectionName, collection)
	}

	routesCode += strings.Join([]string{
		`const PORT = process.env.PORT || 3000`,
		`app.listen(PORT, () => console.log("Server started on port " + PORT))`,
		``,
	}, "\n")

	filePath := filepath.Join(".", "dist", "server.js")
	err := os.WriteFile(filePath, []byte(routesCode), os.ModePerm)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create %s file\n", filePath)
		os.Exit(1)
	}
}

func generatorCollectionRoutes(collectionName string, collection config.Collection) string {
	return strings.Join([]string{
		``,
		generatorGetRoute(collectionName),
		``,
	}, "\n")
}
