package i18n

import (
	"errors"

	"golang.org/x/text/language"
)

// MessageFile represents a parsed message file.
type MessageFile struct {
	Path     string
	Tag      language.Tag
	Format   string
	Messages []*Message
}

// ParseMessageFileBytes returns the messages parsed from file.
func ParseMessageFileBytes(buf []byte, path string, unmarshalFuncs map[string]UnmarshalFunc) (*MessageFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const nestedSeparator = "."

var errInvalidTranslationFile = errors.New("invalid translation file, expected key-values, got a single value")

// recGetMessages looks for translation messages inside "raw" parameter,
// scanning nested maps using recursion.
func recGetMessages(raw interface{}, isMapMessage, isInitialCall bool) ([]*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// recursively scan map items

// recursively scan map items

// Backward compatibility for v1 file format.

// recursively scan slice items

func addChildMessages(id string, data interface{}, messages []*Message) ([]*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// start with innermost key

// update ID with each nested key on the way

func parsePath(path string) (langTag, format string) { _ = "STUB: not implemented"; return "", "" }
