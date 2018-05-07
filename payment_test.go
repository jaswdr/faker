package faker

import (
	"strings"
	"testing"
)

func TestCreditCardType(t *testing.T) {
	p := New().Payment()
	Expect(t, true, len(p.CreditCardType()) > 0)
}

func TestCreditCardNumber(t *testing.T) {
	p := New().Payment()
	Expect(t, true, len(p.CreditCardNumber()) > 0)
}

func TestCreditCardExpirationDateString(t *testing.T) {
	p := New().Payment()

	date := p.CreditCardExpirationDateString()
	Expect(t, 5, len(date))
	Expect(t, true, strings.Contains(date, "/"))
}
