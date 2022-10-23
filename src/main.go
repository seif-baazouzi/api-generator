package main

import (
	"api-generator/src/config"
	"api-generator/src/database"
	"api-generator/src/others"
	"api-generator/src/routes"
	"api-generator/src/utils"
	"flag"
	"fmt"
	"os"
)

func main() {
	config := parseConfigFile()

	utils.ClearDirectory("dist")
	database.GenerateDB(config)
	routes.GenerateRoutes(config)
	others.GeneratePackageJson(config.ProjectName)
	others.GenerateDockerFiles(config.ProjectName)

	fmt.Println("[+] Done!")
}

func parseConfigFile() config.Config {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: ./main [config-file]\n")
		os.Exit(1)
	}

	configPath := args[0]
	config, err := config.Parse(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not read config file %s\n", configPath)
		os.Exit(1)
	}

	return config
}
