// Command goi18n manages message files used by the i18n package.
//
//	go get -u github.com/nicksnyder/go-i18n/v2/goi18n
//	goi18n -help
//
// Use `goi18n extract` to create a message file that contains the messages defined in your Go source files.
//
//	# en.toml
//	[PersonCats]
//	description = "The number of cats a person has"
//	one = "{{.Name}} has {{.Count}} cat."
//	other = "{{.Name}} has {{.Count}} cats."
//
// Use `goi18n merge` to create message files for translation.
//
//	# translate.es.toml
//	[PersonCats]
//	description = "The number of cats a person has"
//	hash = "sha1-f937a0e05e19bfe6cd70937c980eaf1f9832f091"
//	one = "{{.Name}} has {{.Count}} cat."
//	other = "{{.Name}} has {{.Count}} cats."
//
// Use `goi18n merge` to merge translated message files with your existing message files.
//
//	# active.es.toml
//	[PersonCats]
//	description = "The number of cats a person has"
//	hash = "sha1-f937a0e05e19bfe6cd70937c980eaf1f9832f091"
//	one = "{{.Name}} tiene {{.Count}} gato."
//	other = "{{.Name}} tiene {{.Count}} gatos."
//
// Load the active messages into your bundle.
//
//	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
//	bundle.MustLoadMessageFile("active.es.toml")
package main

import (
	"os"

	"golang.org/x/text/language"
)

func mainUsage() { _ = "STUB: not implemented"; return }

type command interface {
	name() string
	parse(arguments []string) error
	execute() error
}

func main() {
	os.Exit(testableMain(os.Args[1:]))
}

func testableMain(args []string) int { _ = "STUB: not implemented"; return 0 }

type languageTag language.Tag

func (lt languageTag) String() string { _ = "STUB: not implemented"; return "" }

func (lt *languageTag) Set(value string) error { _ = "STUB: not implemented"; return nil }

func (lt languageTag) Tag() language.Tag { _ = "STUB: not implemented"; return *new(language.Tag) }
