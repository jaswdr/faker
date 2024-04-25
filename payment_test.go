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
	Expect(t, true, len(p.CreditCardNumber()) > 0)
}

func TestCreditCardExpirationDateString(t *testing.T) {
	p := New().Payment()

	date := p.CreditCardExpirationDateString()
	re := regexp.MustCompile(`^(0[1-9]|1[0-2])/\d{2}$`)
	Expect(t, true, re.MatchString(date))
}
