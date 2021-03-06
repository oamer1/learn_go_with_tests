package propertybasedtests

import (
	"errors"
	"strings"
)

type romanNumeral struct {
	Value  uint16
	Symbol string
}
type romanNumerals []romanNumeral

// Instead you take the next highest symbol and then "subtract" by putting a symbol to the left of it.
// Not all symbols can be used as subtractors; only I (1), X (10) and C (100).
var allRomanNumerals = romanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

//IsValid checks if the number is valid in Roman numeral
func IsValid(n uint16) bool {
	if n <= 0 || n > 3999 {
		return false
	}

	return true
}

func ConvertToRoman(arabic uint16) (string, error) {

	if !IsValid(arabic) {
		return "", errors.New("number is not roman numerals, 0 <= roman <= 399")
	}
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String(), nil
}

func (r romanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func (r romanNumerals) Exists(symbols ...byte) bool {

	return r.ValueOf(symbols...) != 0
}

// later..
func ConvertToArabic(roman string) (total uint16) {

	for _, romanNum := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(romanNum...)
	}
	return
}

type windowedRoman string

// Separate roman string into its constituents components
func (w windowedRoman) Symbols() (symbols [][]byte) {

	for i := 0; i < len(w); i++ {
		notAtEnd := i+1 < len(w)
		if notAtEnd && isSubtractive(w[i]) && allRomanNumerals.Exists(w[i], w[i+1]) {
			symbols = append(symbols, []byte{w[i], w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{w[i]})

		}
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'

}
