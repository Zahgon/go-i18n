package main

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/nicksnyder/go-i18n/v2/internal/plural"
	"golang.org/x/text/language"
)

func usageMerge() { _ = "STUB: not implemented"; return }

type mergeCommand struct {
	messageFiles   []string
	sourceLanguage languageTag
	outdir         string
	format         string
}

func (mc *mergeCommand) name() string { _ = "STUB: not implemented"; return "" }

func (mc *mergeCommand) parse(args []string) error { _ = "STUB: not implemented"; return nil }

func (mc *mergeCommand) execute() error { _ = "STUB: not implemented"; return nil }

// Ignore error since it isn't guaranteed to exist.

type fileSystemOp struct {
	writeFiles  map[string][]byte
	deleteFiles []string
}

func merge(messageFiles map[string][]byte, sourceLanguageTag language.Tag, outdir, outputFormat string) (*fileSystemOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Non-standard languages not supported because
// we don't know if translations are complete or not.

// Check all unmerged message templates for this message id.

// Ignore empty hashes for v1 backward compatibility.

// This was translated from different content so discard.

// Merge in the translated messages.

// Non-standard languages not supported because
// we don't know if translations are complete or not.

// activeDst returns the active part of the dst and whether dst is a complete translation of src.
func activeDst(src, dst *i18n.MessageTemplate, pluralRule *plural.Rule) (active *i18n.MessageTemplate, translateMessageTemplate *i18n.MessageTemplate) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hash(t *i18n.MessageTemplate) string { _ = "STUB: not implemented"; return "" }
