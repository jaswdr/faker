package faker

import (
	"testing"
)

func TestCurrency(t *testing.T) {
	c := New().Currency()
	currency := c.Currency()
	NotExpect(t, "", currency)
	ExpectInString(t, currency, currencies)
}

func TestCurrencyCode(t *testing.T) {
	c := New().Currency()
	code := c.Code()
	ExpectInString(t, code, currenciesCodes)
}

func TestCurrencyNumber(t *testing.T) {
	c := New().Currency()
	ExpectInInt(t, c.Number(), currenciesNumbers)
}

func TestCurrencyCountry(t *testing.T) {
	c := New().Currency()
	country := c.Country()
	NotExpect(t, "", country)
	ExpectInString(t, country, currenciesCountries)
}

func TestCurrencyAndCode(t *testing.T) {
	c := New().Currency()
	currency, code := c.CurrencyAndCode()
	NotExpect(t, "", currency)
	ExpectInString(t, currency, currencies)
	ExpectInString(t, code, currenciesCodes)
}
