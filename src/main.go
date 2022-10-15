package main

import (
	"api-generator/src/config"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		usage()
	}

	configPath := args[0]
	fmt.Println(config.Parse(configPath))
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: ./main [config-file]\n")
	os.Exit(1)
}
