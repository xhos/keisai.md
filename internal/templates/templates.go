package templates

import (
	"bytes"
	"fmt"
	"github.com/xhos/keisai.md/internal/parser"
	"html/template"
	"os"
	"path/filepath"
)

func loadTemplate(templateName string) (*template.Template, error) {
	templatePath := filepath.Join("internal", "templates", templateName, "index.html")

	content, err := os.ReadFile(templatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read template file %s: %w", templatePath, err)
	}

	return template.New(templateName).Parse(string(content))
}

func renderTemplate(tmp *template.Template, data parser.PageData) (*bytes.Buffer, error) {
	var buf bytes.Buffer

	if err := tmp.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return &buf, nil
}

func WrapInTemplate(data parser.PageData, templateName string) (*bytes.Buffer, error) {
	tmpl, err := loadTemplate(templateName)
	if err != nil {
		return nil, fmt.Errorf("failed to load template %s: %w", templateName, err)
	}

	return renderTemplate(tmpl, data)
}
