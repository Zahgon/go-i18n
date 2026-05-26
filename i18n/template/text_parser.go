package template

import (
	"text/template"
)

// TextParser is a Parser that uses text/template.
type TextParser struct {
	LeftDelim  string
	RightDelim string
	Funcs      template.FuncMap
	Option     string
}

func (te *TextParser) Cacheable() bool { _ = "STUB: not implemented"; return false }

func (te *TextParser) Parse(src, leftDelim, rightDelim string) (ParsedTemplate, error) {
	_ = "STUB: not implemented"
	return *new(ParsedTemplate), nil
}

// Fast path to avoid parsing a template that has no actions.

type parsedTextTemplate struct {
	tmpl *template.Template
}

func (t *parsedTextTemplate) Execute(data any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
