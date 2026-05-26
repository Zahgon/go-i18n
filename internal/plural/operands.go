package plural

// Operands is a representation of http://unicode.org/reports/tr35/tr35-numbers.html#Operands
// If there is a compact decimal exponent value C, then the N, I, V, W, F, and T values are computed after shifting the decimal point in the original by the ‘c’ value.
// So for 1.2c3, the values are the same as those of 1200: i=1200 and f=0.
// Similarly, 1.2005c3 has i=1200 and f=5 (corresponding to 1200.5).
type Operands struct {
	N float64 // absolute value of the source number (integer and decimals)
	I int64   // integer digits of n
	V int64   // number of visible fraction digits in n, with trailing zeros
	W int64   // number of visible fraction digits in n, without trailing zeros
	F int64   // visible fractional digits in n, with trailing zeros
	T int64   // visible fractional digits in n, without trailing zeros
	C int64   // compact decimal exponent value: exponent of the power of 10 used in compact decimal formatting.
}

// NEqualsAny returns true if o represents an integer equal to any of the arguments.
func (o *Operands) NEqualsAny(any ...int64) bool { _ = "STUB: not implemented"; return false }

// NModEqualsAny returns true if o represents an integer equal to any of the arguments modulo mod.
func (o *Operands) NModEqualsAny(mod int64, any ...int64) bool {
	_ = "STUB: not implemented"
	return false
}

// NInRange returns true if o represents an integer in the closed interval [from, to].
func (o *Operands) NInRange(from, to int64) bool { _ = "STUB: not implemented"; return false }

// NModInRange returns true if o represents an integer in the closed interval [from, to] modulo mod.
func (o *Operands) NModInRange(mod, from, to int64) bool { _ = "STUB: not implemented"; return false }

// NewOperands returns the operands for number.
func NewOperands(number interface{}) (*Operands, error) { _ = "STUB: not implemented"; return nil, nil }

func newOperandsInt64(i int64) *Operands { _ = "STUB: not implemented"; return nil }

func splitSignificandExponent(s string) (significand, exponent string) {
	_ = "STUB: not implemented"
	return "", ""
}

func shiftDecimalLeft(s string, n int) string { _ = "STUB: not implemented"; return "" }

func shiftDecimalRight(s string, n int) string { _ = "STUB: not implemented"; return "" }

func applyExponent(s string, exponent int) string { _ = "STUB: not implemented"; return "" }

func newOperandsString(s string) (*Operands, error) { _ = "STUB: not implemented"; return nil, nil }

// We are storing C as an int64 but only allowing
// numbers that fit into the bitsize of an int
// so C is safe to cast as a int later.
