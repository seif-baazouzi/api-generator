package utils

import (
	"bytes"
	"html/template"
)

func TemplateToString(t *template.Template, data interface{}) (string, error) {
	buffer := &bytes.Buffer{}
	err := t.Execute(buffer, data)

	return buffer.String(), err
}
