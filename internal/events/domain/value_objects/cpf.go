package value_objects

import (
	"fmt"
	"regexp"
	"strconv"
)

var nonDigitRegex = regexp.MustCompile(`\D`)

type Cpf struct {
	value string
}

func NewCpf(raw string) (Cpf, error) {
	digits := nonDigitRegex.ReplaceAllString(raw, "")

	if len(digits) != 11 {
		return Cpf{}, fmt.Errorf("CPF must have 11 digits, got %d", len(digits))
	}

	if allSameDigit(digits) {
		return Cpf{}, fmt.Errorf("CPF must not be all same digits")
	}

	if !validateCheckDigits(digits) {
		return Cpf{}, fmt.Errorf("CPF is invalid")
	}

	return Cpf{value: digits}, nil
}

func allSameDigit(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func validateCheckDigits(digits string) bool {
	sum := 0
	for i := 0; i < 9; i++ {
		d, _ := strconv.Atoi(string(digits[i]))
		sum += d * (10 - i)
	}

	first := 11 - (sum % 11)
	if first > 9 {
		first = 0
	}

	sum = 0
	for i := 0; i < 10; i++ {
		d, _ := strconv.Atoi(string(digits[i]))
		sum += d * (11 - i)
	}

	second := 11 - (sum % 11)
	if second > 9 {
		second = 0
	}

	d9, _ := strconv.Atoi(string(digits[9]))
	d10, _ := strconv.Atoi(string(digits[10]))

	return first == d9 && second == d10
}
