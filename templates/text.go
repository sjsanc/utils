package templates

import (
	"bytes"
	"text/template"
)

// RenderText processes a text template with the provided data.
func RenderText(tmpl string, data interface{}) (string, error) {
	t, err := template.New("textTemplate").Funcs(GetCustomFuncMap()).Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
