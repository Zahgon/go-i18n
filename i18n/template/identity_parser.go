package template

// IdentityParser is an Parser that does no parsing and returns template string unchanged.
type IdentityParser struct{}

func (IdentityParser) Cacheable() bool {
	_ = "STUB: not implemented"
	// Caching is not necessary because Parse is cheap.
	return false
}

func (IdentityParser) Parse(src, leftDelim, rightDelim string) (ParsedTemplate, error) {
	_ = "STUB: not implemented"
	return *new(ParsedTemplate), nil
}

type identityParsedTemplate struct {
	src string
}

func (t *identityParsedTemplate) Execute(data any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
