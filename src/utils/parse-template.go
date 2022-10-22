package utils

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

func ParseTemplate(templatePath string, data interface{}) string {
	template, err := template.ParseFiles(templatePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse template %s\n", templatePath)
		os.Exit(1)
	}

	buffer := &bytes.Buffer{}
	err = template.Execute(buffer, data)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse template %s\n", templatePath)
		os.Exit(1)
	}

	return buffer.String()
}
