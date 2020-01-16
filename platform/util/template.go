package util

import (
	"bytes"
	"fmt"
	"text/template"
)

type Template struct {
	tpl *template.Template
}

func (t *Template) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tplString string
	if err := unmarshal(&tplString); err != nil {
		return err
	}

	tpl, err := template.New("validation").Parse(tplString)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	t.tpl = tpl

	return nil
}

func (t *Template) Execute(args interface{}) (string, error) {
	var buf bytes.Buffer
	if err := t.tpl.Execute(&buf, args); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}
