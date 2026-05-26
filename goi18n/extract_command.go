package main

import (
	"go/ast"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func usageExtract() { _ = "STUB: not implemented"; return }

type extractCommand struct {
	paths          []string
	sourceLanguage languageTag
	outdir         string
	format         string
}

func (ec *extractCommand) name() string { _ = "STUB: not implemented"; return "" }

func (ec *extractCommand) parse(args []string) error { _ = "STUB: not implemented"; return nil }

func (ec *extractCommand) execute() error { _ = "STUB: not implemented"; return nil }

// Don't extract from test files.

type duplicateMessageIDErr struct {
	messageID string
}

func (e *duplicateMessageIDErr) Error() string { _ = "STUB: not implemented"; return "" }

// extractMessages extracts messages from the bytes of a Go source file.
func extractMessages(buf []byte) ([]*i18n.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newExtractor(file *ast.File) *extractor { _ = "STUB: not implemented"; return nil }

type extractor struct {
	i18nPackageName string
	messages        []*i18n.Message
}

func (e *extractor) Visit(node ast.Node) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}

func (e *extractor) extractMessages(node ast.Node) { _ = "STUB: not implemented"; return }

func (e *extractor) isMessageType(expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

func unwrapSelectorExpr(e ast.Expr) *ast.SelectorExpr { _ = "STUB: not implemented"; return nil }

func (e *extractor) extractMessage(cl *ast.CompositeLit) { _ = "STUB: not implemented"; return }

func extractStringLiteral(expr ast.Expr) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func i18nPackageName(file *ast.File) string { _ = "STUB: not implemented"; return "" }
