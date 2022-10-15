package main

import (
	"api-generator/src/config"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	config := parseConfigFile()
	clearDistDirectory()
	fmt.Println(config)
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

func clearDistDirectory() {
	distPath := filepath.Join(".", "dist")
	os.RemoveAll(distPath)
	err := os.MkdirAll(distPath, os.ModePerm)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create dist directory\n")
		os.Exit(1)
	}
}
