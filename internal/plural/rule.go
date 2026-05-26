package plural

// Rule defines the CLDR plural rules for a language.
// http://www.unicode.org/cldr/charts/latest/supplemental/language_plural_rules.html
// http://unicode.org/reports/tr35/tr35-numbers.html#Operands
type Rule struct {
	PluralForms    map[Form]struct{}
	PluralFormFunc func(*Operands) Form
}

func addPluralRules(rules Rules, ids []string, ps *Rule) { _ = "STUB: not implemented"; return }

func newPluralFormSet(pluralForms ...Form) map[Form]struct{} { _ = "STUB: not implemented"; return nil }

func intInRange(i, from, to int64) bool { _ = "STUB: not implemented"; return false }

func intEqualsAny(i int64, any ...int64) bool { _ = "STUB: not implemented"; return false }
