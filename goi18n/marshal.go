package main

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func writeFile(outdir, label string, langTag language.Tag, format string, messageTemplates map[string]*i18n.MessageTemplate, sourceLanguage bool) (path string, content []byte, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func marshalValue(messageTemplates map[string]*i18n.MessageTemplate, sourceLanguage bool) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func marshal(v interface{}, format string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
