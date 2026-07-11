package faker

import (
	"regexp"
	"testing"
)

func TestCreditCardType(t *testing.T) {
	p := New().Payment()
	Expect(t, true, len(p.CreditCardType()) > 0)
}

func TestCreditCardNumber(t *testing.T) {
	p := New().Payment()
	number := p.CreditCardNumber()
	Expect(t, true, len(number) >= 13)
	Expect(t, true, IsValidLuhn(number))
}

func TestCreditCardNumberForType(t *testing.T) {
	p := New().Payment()

	tests := []struct {
		cardType string
		length   int
		prefix   string
	}{
		{"Visa", 16, "4"},
		{"MasterCard", 16, "5"},
		{"American Express", 15, "3"},
		{"Discover Card", 16, "6"},
		{"Visa Retired", 13, "4"},
	}

	for _, tt := range tests {
		number := p.CreditCardNumberForType(tt.cardType)
		Expect(t, tt.length, len(number))
		Expect(t, true, IsValidLuhn(number))
		Expect(t, tt.prefix, string(number[0]))
	}
}

func TestIsValidLuhn(t *testing.T) {
	Expect(t, true, IsValidLuhn("4532015112830366"))
	Expect(t, false, IsValidLuhn("4532015112830367"))
}

func TestCreditCardExpirationDateString(t *testing.T) {
	p := New().Payment()

	date := p.CreditCardExpirationDateString()
	re := regexp.MustCompile(`^(0[1-9]|1[0-2])/\d{2}$`)
	Expect(t, true, re.MatchString(date))
}
