package internal

import (
	"sync"

	"github.com/nicksnyder/go-i18n/v2/i18n/template"
)

// Template stores the template for a string and a cached version of the parsed template if they are cacheable.
type Template struct {
	Src        string
	LeftDelim  string
	RightDelim string

	parseOnce      sync.Once
	parsedTemplate template.ParsedTemplate
	parseError     error
}

func (t *Template) Execute(parser template.Parser, data interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
