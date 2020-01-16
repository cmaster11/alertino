package util

import (
	"bytes"
	"fmt"
	"text/template"
)

type IfTemplate struct {
	tpl *template.Template
}

func (t *IfTemplate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tplString string
	if err := unmarshal(&tplString); err != nil {
		return err
	}

	tplStringComplete := fmt.Sprintf(`{{ if %s }}true{{else}}false{{end}}`, tplString)
	tpl, err := template.New("validation").Parse(tplStringComplete)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	t.tpl = tpl

	return nil
}

func (t *IfTemplate) Matches(args interface{}) (bool, error) {
	var buf bytes.Buffer
	if err := t.tpl.Execute(&buf, args); err != nil {
		return false, fmt.Errorf("failed to execute template: %w", err)
	}

	result := buf.String()

	return result == "true", nil
}
