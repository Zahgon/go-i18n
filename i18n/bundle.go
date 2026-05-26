package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/internal/plural"

	"golang.org/x/text/language"
)

// UnmarshalFunc unmarshals data into v.
type UnmarshalFunc func(data []byte, v interface{}) error

// Bundle stores a set of messages and pluralization rules.
// Most applications only need a single bundle
// that is initialized early in the application's lifecycle.
// It is not goroutine safe to modify the bundle while Localizers
// are reading from it.
type Bundle struct {
	defaultLanguage  language.Tag
	unmarshalFuncs   map[string]UnmarshalFunc
	messageTemplates map[language.Tag]map[string]*MessageTemplate
	pluralRules      plural.Rules
	tags             []language.Tag
	matcher          language.Matcher
}

// artTag is the language tag used for artificial languages
// https://en.wikipedia.org/wiki/Codes_for_constructed_languages
var artTag = language.MustParse("art")

// NewBundle returns a bundle with a default language and a default set of plural rules.
func NewBundle(defaultLanguage language.Tag) *Bundle { _ = "STUB: not implemented"; return nil }

// RegisterUnmarshalFunc registers an UnmarshalFunc for format.
func (b *Bundle) RegisterUnmarshalFunc(format string, unmarshalFunc UnmarshalFunc) {
	_ = "STUB: not implemented"
	return
}

// LoadMessageFile loads the bytes from path
// and then calls ParseMessageFileBytes.
func (b *Bundle) LoadMessageFile(path string) (*MessageFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustLoadMessageFile is similar to LoadMessageFile
// except it panics if an error happens.
func (b *Bundle) MustLoadMessageFile(path string) { _ = "STUB: not implemented"; return }

// ParseMessageFileBytes parses the bytes in buf to add translations to the bundle.
//
// The format of the file is everything after the last ".".
//
// The language tag of the file is everything after the second to last "." or after the last path separator, but before the format.
func (b *Bundle) ParseMessageFileBytes(buf []byte, path string) (*MessageFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustParseMessageFileBytes is similar to ParseMessageFileBytes
// except it panics if an error happens.
func (b *Bundle) MustParseMessageFileBytes(buf []byte, path string) {
	_ = "STUB: not implemented"
	return
}

// AddMessages adds messages for a language.
// It is useful if your messages are in a format not supported by ParseMessageFileBytes.
func (b *Bundle) AddMessages(tag language.Tag, messages ...*Message) error {
	_ = "STUB: not implemented"
	return nil
}

// MustAddMessages is similar to AddMessages except it panics if an error happens.
func (b *Bundle) MustAddMessages(tag language.Tag, messages ...*Message) {
	_ = "STUB: not implemented"
	return
}

func (b *Bundle) addTag(tag language.Tag) { _ = "STUB: not implemented"; return }

// Tag already exists

// LanguageTags returns the list of language tags
// of all the translations loaded into the bundle
func (b *Bundle) LanguageTags() []language.Tag { _ = "STUB: not implemented"; return nil }

func (b *Bundle) getMessageTemplate(tag language.Tag, id string) *MessageTemplate {
	_ = "STUB: not implemented"
	return nil
}
