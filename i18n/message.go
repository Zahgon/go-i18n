package i18n

// Message is a string that can be localized.
type Message struct {
	// ID uniquely identifies the message.
	ID string

	// Hash uniquely identifies the content of the message
	// that this message was translated from.
	Hash string

	// Description describes the message to give additional
	// context to translators that may be relevant for translation.
	Description string

	// LeftDelim is the left Go template delimiter.
	LeftDelim string

	// RightDelim is the right Go template delimiter.
	RightDelim string

	// Zero is the content of the message for the CLDR plural form "zero".
	Zero string

	// One is the content of the message for the CLDR plural form "one".
	One string

	// Two is the content of the message for the CLDR plural form "two".
	Two string

	// Few is the content of the message for the CLDR plural form "few".
	Few string

	// Many is the content of the message for the CLDR plural form "many".
	Many string

	// Other is the content of the message for the CLDR plural form "other".
	Other string
}

// NewMessage parses data and returns a new message.
func NewMessage(data interface{}) (*Message, error) { _ = "STUB: not implemented"; return nil, nil }

// MustNewMessage is similar to NewMessage except it panics if an error happens.
func MustNewMessage(data interface{}) *Message { _ = "STUB: not implemented"; return nil }

// unmarshalInterface unmarshals a message from data.
func (m *Message) unmarshalInterface(v interface{}) error { _ = "STUB: not implemented"; return nil }

type keyTypeErr struct {
	key interface{}
}

func (err *keyTypeErr) Error() string { _ = "STUB: not implemented"; return "" }

type valueTypeErr struct {
	value interface{}
}

func (err *valueTypeErr) Error() string { _ = "STUB: not implemented"; return "" }

func stringMap(v interface{}) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func stringSubmap(k string, v interface{}, strdata map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

var reservedKeys = map[string]struct{}{
	"id":          {},
	"description": {},
	"hash":        {},
	"leftdelim":   {},
	"rightdelim":  {},
	"zero":        {},
	"one":         {},
	"two":         {},
	"few":         {},
	"many":        {},
	"other":       {},
	"translation": {},
}

func isReserved(key string, val any) bool { _ = "STUB: not implemented"; return false }

// isMessage returns true if v contains only message keys and false if it contains no message keys.
// It returns an error if v contains both message and non-message keys.
// - {"message": {"description": "world"}} is a message
// - {"error": {"description": "world", "foo": "bar"}} is an error
// - {"notmessage": {"description": {"hello": "world"}}} is not a message
// - {"notmessage": {"foo": "bar"}} is not a message
func isMessage(v interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

type mixedKeysError struct {
	reservedKeys   []string
	unreservedKeys []string
}

func (e *mixedKeysError) Error() string { _ = "STUB: not implemented"; return "" }
