package i18n

import (
	texttemplate "text/template"

	"github.com/nicksnyder/go-i18n/v2/i18n/template"
	"github.com/nicksnyder/go-i18n/v2/internal"
	"github.com/nicksnyder/go-i18n/v2/internal/plural"
)

// MessageTemplate is an executable template for a message.
type MessageTemplate struct {
	*Message
	PluralTemplates map[plural.Form]*internal.Template
}

// NewMessageTemplate returns a new message template.
func NewMessageTemplate(m *Message) *MessageTemplate { _ = "STUB: not implemented"; return nil }

func setPluralTemplate(pluralTemplates map[plural.Form]*internal.Template, pluralForm plural.Form, src, leftDelim, rightDelim string) {
	_ = "STUB: not implemented"
	return
}

type pluralFormNotFoundError struct {
	pluralForm plural.Form
	messageID  string
}

func (e pluralFormNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

// Execute executes the template for the plural form and template data.
// Deprecated: This method is no longer used internally by go-i18n and it probably should not have been exported to
// begin with. Its replacement is not exported. If you depend on this method for some reason and/or have
// a use case for exporting execute, please file an issue.
func (mt *MessageTemplate) Execute(pluralForm plural.Form, data interface{}, funcs texttemplate.FuncMap) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (mt *MessageTemplate) execute(pluralForm plural.Form, data interface{}, parser template.Parser) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
