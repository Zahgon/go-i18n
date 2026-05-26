package main

import (
	"encoding/xml"
	"regexp"
)

// SupplementalData is the top level struct of plural.xml
type SupplementalData struct {
	XMLName      xml.Name      `xml:"supplementalData"`
	PluralGroups []PluralGroup `xml:"plurals>pluralRules"`
}

// PluralGroup is a group of locales with the same plural rules.
type PluralGroup struct {
	Locales     string       `xml:"locales,attr"`
	PluralRules []PluralRule `xml:"pluralRule"`
}

// Name returns a unique name for this plural group.
func (pg *PluralGroup) Name() string { _ = "STUB: not implemented"; return "" }

// SplitLocales returns all the locales in the PluralGroup as a slice.
func (pg *PluralGroup) SplitLocales() []string { _ = "STUB: not implemented"; return nil }

// PluralRule is the rule for a single plural form.
type PluralRule struct {
	Count string `xml:"count,attr"`
	Rule  string `xml:",innerxml"`
}

// CountTitle returns the title case of the PluralRule's count.
func (pr *PluralRule) CountTitle() string { _ = "STUB: not implemented"; return "" }

// Condition returns the condition where the PluralRule applies.
func (pr *PluralRule) Condition() string { _ = "STUB: not implemented"; return "" }

// Examples returns the integer and decimal examples for the PluralRule.
func (pr *PluralRule) Examples() (integers []string, decimals []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IntegerExamples returns the integer examples for the PluralRule.
func (pr *PluralRule) IntegerExamples() []string { _ = "STUB: not implemented"; return nil }

// DecimalExamples returns the decimal examples for the PluralRule.
func (pr *PluralRule) DecimalExamples() []string { _ = "STUB: not implemented"; return nil }

var relationRegexp = regexp.MustCompile(`([niftvwce])(?:\s*%\s*([0-9]+))?\s*(!=|=)(.*)`)

// GoCondition converts the XML condition to valid Go code.
func (pr *PluralRule) GoCondition() string { _ = "STUB: not implemented"; return "" }

// E is a deprecated symbol for C
// https://unicode.org/reports/tr35/tr35-numbers.html#Plural_Operand_Meanings
