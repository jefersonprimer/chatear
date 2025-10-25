package application

import (
	"bytes"
	"html/template"
	"path/filepath"
)

// TemplateParser defines the interface for parsing and rendering email templates.
type TemplateParser interface {
	ParseTemplate(templateName string, data map[string]interface{}) (string, error)
}

// HTMLTemplateParser is an implementation of TemplateParser for HTML templates.
type HTMLTemplateParser struct {
	basePath string
}

// NewHTMLTemplateParser creates a new HTMLTemplateParser.
func NewHTMLTemplateParser(basePath string) *HTMLTemplateParser {
	return &HTMLTemplateParser{basePath: basePath}
}

// ParseTemplate parses and renders an HTML template.
func (p *HTMLTemplateParser) ParseTemplate(templateName string, data map[string]interface{}) (string, error) {
	templatePath := filepath.Join(p.basePath, templateName)
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
