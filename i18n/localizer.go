package i18n

import (
	texttemplate "text/template"

	"github.com/nicksnyder/go-i18n/v2/i18n/template"
	"github.com/nicksnyder/go-i18n/v2/internal/plural"
	"golang.org/x/text/language"
)

// Localizer provides Localize and MustLocalize methods that return localized messages.
// Localize and MustLocalize methods use a language.Tag matching algorithm based
// on the best possible value. This algorithm may cause an unexpected language.Tag returned
// value depending on the order of the tags stored in memory. For example, if the bundle
// used to create a Localizer instance ingested locales following this order
// ["en-US", "en-GB", "en-IE", "en"] and the locale "en" is asked, the underlying matching
// algorithm will return "en-US" thinking it is the best match possible. More information
// about the algorithm in this Github issue: https://github.com/golang/go/issues/49176.
// There is additional informations inside the Go code base:
// https://github.com/golang/text/blob/master/language/match.go#L142
type Localizer struct {
	// bundle contains the messages that can be returned by the Localizer.
	bundle *Bundle

	// tags is the list of language tags that the Localizer checks
	// in order when localizing a message.
	tags []language.Tag
}

// NewLocalizer returns a new Localizer that looks up messages
// in the bundle according to the language preferences in langs.
// It can parse Accept-Language headers as defined in http://www.ietf.org/rfc/rfc2616.txt.
func NewLocalizer(bundle *Bundle, langs ...string) *Localizer {
	_ = "STUB: not implemented"
	return nil
}

func parseTags(langs []string) []language.Tag { _ = "STUB: not implemented"; return nil }

// LocalizeConfig configures a call to the Localize method on Localizer.
type LocalizeConfig struct {
	// MessageID is the id of the message to lookup.
	// This field is ignored if DefaultMessage is set.
	MessageID string

	// TemplateData is the data passed when executing the message's template.
	// If TemplateData is nil and PluralCount is not nil, then the message template
	// will be executed with data that contains the plural count.
	TemplateData interface{}

	// PluralCount determines which plural form of the message is used.
	PluralCount interface{}

	// DefaultMessage is used if the message is not found in any message files.
	DefaultMessage *Message

	// Funcs is used to configure a template.TextParser if TemplateParser is not set.
	Funcs texttemplate.FuncMap

	// The TemplateParser to use for parsing templates.
	// If one is not set, a template.TextParser is used (configured with Funcs if it is set).
	TemplateParser template.Parser
}

var defaultTextParser = &template.TextParser{}

func (lc *LocalizeConfig) getTemplateParser() template.Parser {
	_ = "STUB: not implemented"
	return *new(template.Parser)
}

type invalidPluralCountErr struct {
	messageID   string
	pluralCount interface{}
	err         error
}

func (e *invalidPluralCountErr) Error() string { _ = "STUB: not implemented"; return "" }

// MessageNotFoundErr is returned from Localize when a message could not be found.
type MessageNotFoundErr struct {
	Tag       language.Tag
	MessageID string
}

func (e *MessageNotFoundErr) Error() string { _ = "STUB: not implemented"; return "" }

type messageIDMismatchErr struct {
	messageID        string
	defaultMessageID string
}

func (e *messageIDMismatchErr) Error() string { _ = "STUB: not implemented"; return "" }

// Localize returns a localized message.
func (l *Localizer) Localize(lc *LocalizeConfig) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// LocalizeMessage returns a localized message.
func (l *Localizer) LocalizeMessage(msg *Message) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TODO: uncomment this (and the test) when extract has been updated to extract these call sites too.
// Localize returns a localized message.
// func (l *Localizer) LocalizeMessageID(messageID string) (string, error) {
// 	return l.Localize(&LocalizeConfig{
// 		MessageID: messageID,
// 	})
// }

// LocalizeWithTag returns a localized message and the language tag.
// It may return a best effort localized message even if an error happens.
func (l *Localizer) LocalizeWithTag(lc *LocalizeConfig) (string, language.Tag, error) {
	_ = "STUB: not implemented"
	return "", *new(language.Tag), nil
}

// Attempt to fallback to "Other" pluralization in case translations are incomplete.

func (l *Localizer) getMessageTemplate(id string, defaultMessage *Message) (language.Tag, *MessageTemplate, error) {
	_ = "STUB: not implemented"
	return *new(language.Tag), nil, nil
}

// Fallback to default language in bundle.

// Fallback to default message.

func (l *Localizer) pluralForm(tag language.Tag, operands *plural.Operands) plural.Form {
	_ = "STUB: not implemented"
	return *new(plural.Form)
}

// MustLocalize is similar to Localize, except it panics if an error happens.
func (l *Localizer) MustLocalize(lc *LocalizeConfig) string { _ = "STUB: not implemented"; return "" }

// MustLocalizeMessage is similar to LocalizeMessage, except it panics if an error happens.
func (l *Localizer) MustLocalizeMessage(msg *Message) string { _ = "STUB: not implemented"; return "" }
