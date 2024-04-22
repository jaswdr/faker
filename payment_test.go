package faker

import (
	"regexp"
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

var ibanRegex = map[string]string{
	"AD": `^AD\d{2}\d{4}\d{4}[A-Z0-9]{12}$`,
	"AE": `^AE\d{2}\d{3}\d{16}$`,
	"AL": `^AL\d{2}\d{8}[A-Z0-9]{16}$`,
	"AT": `^AT\d{2}\d{5}\d{11}$`,
	"AZ": `^AZ\d{2}[A-Z]{4}[A-Z0-9]{20}$`,
	"BA": `^BA\d{2}\d{3}\d{3}\d{8}\d{2}$`,
	"BE": `^BE\d{2}\d{3}\d{7}\d{2}$`,
	"BG": `^BG\d{2}[A-Z]{4}\d{4}\d{2}[A-Z0-9]{8}$`,
	"BH": `^BH\d{2}[A-Z]{4}[A-Z0-9]{14}$`,
	"BR": `^BR\d{2}\d{8}\d{5}\d{10}[A-Z]{1}[A-Z0-9]{1}$`,
	"CH": `^CH\d{2}\d{5}[A-Z0-9]{12}$`,
	"CR": `^CR\d{2}\d{4}\d{14}$`,
	"CY": `^CY\d{2}\d{3}\d{5}[A-Z0-9]{16}$`,
	"CZ": `^CZ\d{2}\d{4}\d{6}\d{10}$`,
	"DE": `^DE\d{2}\d{8}\d{10}$`,
	"DK": `^DK\d{2}\d{4}\d{9}\d{1}$`,
	"DO": `^DO\d{2}[A-Z0-9]{4}\d{20}$`,
	"EE": `^EE\d{2}\d{2}\d{2}\d{11}\d{1}$`,
	"ES": `^ES\d{2}\d{4}\d{4}\d{1}\d{1}\d{10}$`,
	"FI": `^FI\d{2}\d{6}\d{7}\d{1}$`,
	"FR": `^FR\d{2}\d{5}\d{5}[A-Z0-9]{11}\d{2}$`,
	"GB": `^GB\d{2}[A-Z]{4}\d{6}\d{8}$`,
	"GE": `^GE\d{2}[A-Z]{2}\d{16}$`,
	"GI": `^GI\d{2}[A-Z]{4}[A-Z0-9]{15}$`,
	"GR": `^GR\d{2}\d{3}\d{4}[A-Z0-9]{16}$`,
	"GT": `^GT\d{2}[A-Z0-9]{4}[A-Z0-9]{20}$`,
	"HR": `^HR\d{2}\d{7}\d{10}$`,
	"HU": `^HU\d{2}\d{3}\d{4}\d{1}\d{15}\d{1}$`,
	"IE": `^IE\d{2}[A-Z]{4}\d{6}\d{8}$`,
	"IL": `^IL\d{2}\d{3}\d{3}\d{13}$`,
	"IS": `^IS\d{2}\d{4}\d{2}\d{6}\d{10}$`,
	"IT": `^IT\d{2}[A-Z]{1}\d{5}\d{5}[A-Z0-9]{12}$`,
	"KW": `^KW\d{2}[A-Z]{4}\d{22}$`,
	"KZ": `^KZ\d{2}\d{3}[A-Z0-9]{13}$`,
	"LB": `^LB\d{2}\d{4}[A-Z0-9]{20}$`,
	"LI": `^LI\d{2}\d{5}[A-Z0-9]{12}$`,
	"LT": `^LT\d{2}\d{5}\d{11}$`,
	"LU": `^LU\d{2}\d{3}[A-Z0-9]{13}$`,
	"LV": `^LV\d{2}[A-Z]{4}[A-Z0-9]{13}$`,
	"MC": `^MC\d{2}\d{5}\d{5}[A-Z0-9]{11}\d{2}$`,
	"MD": `^MD\d{2}[A-Z0-9]{2}[A-Z0-9]{18}$`,
	"ME": `^ME\d{2}\d{3}\d{13}\d{2}$`,
	"MK": `^MK\d{2}\d{3}[A-Z0-9]{10}\d{2}$`,
	"MR": `^MR\d{2}\d{5}\d{5}\d{11}\d{2}$`,
	"MT": `^MT\d{2}[A-Z]{4}\d{5}[A-Z0-9]{18}$`,
	"MU": `^MU\d{2}[A-Z]{4}\d{2}\d{2}\d{12}\d{3}[A-Z]{3}$`,
	"NL": `^NL\d{2}[A-Z]{4}\d{10}$`,
	"NO": `^NO\d{2}\d{4}\d{6}\d{1}$`,
	"PK": `^PK\d{2}[A-Z]{4}[A-Z0-9]{16}$`,
	"PL": `^PL\d{2}\d{8}\d{16}$`,
	"PS": `^PS\d{2}[A-Z]{4}[A-Z0-9]{21}$`,
	"PT": `^PT\d{2}\d{4}\d{4}\d{11}\d{2}$`,
	"RO": `^RO\d{2}[A-Z]{4}[A-Z0-9]{16}$`,
	"RS": `^RS\d{2}\d{3}\d{13}\d{2}$`,
	"SA": `^SA\d{2}\d{2}[A-Z0-9]{18}$`,
	"SE": `^SE\d{2}\d{3}\d{16}\d{1}$`,
	"SI": `^SI\d{2}\d{5}\d{8}\d{2}$`,
	"SK": `^SK\d{2}\d{4}\d{6}\d{10}$`,
	"SM": `^SM\d{2}[A-Z]{1}\d{5}\d{5}[A-Z0-9]{12}$`,
	"TN": `^TN\d{2}\d{2}\d{3}\d{13}\d{2}$`,
	"TR": `^TR\d{2}\d{5}\d{1}[A-Z0-9]{16}$`,
	"VG": `^VG\d{2}[A-Z]{4}\d{16}$`,
}

func TestIban(t *testing.T) {
	p := New().Payment()

	iban := p.Iban()
	Expect(t, true, isIbanValid(iban))
	_, ok := ibanRegex[iban[:2]]
	Expect(t, true, ok)
}

func TestIbanPerCountry(t *testing.T) {
	p := New().Payment()

	for countryCode, regex := range ibanRegex {
		iban := p.ibanForCountry(countryCode)
		Expect(t, countryCode, iban[:2])
		matched, err := regexp.MatchString(regex, iban)
		Expect(t, nil, err)
		Expect(t, true, matched)
		Expect(t, true, isIbanValid(iban))
	}

	Expect(t, "", p.ibanForCountry("unknown"))
}

func TestFormat(t *testing.T) {
	Expect(t, "nnaaaacccccc", format("n2", "a4", "c6"))
	Expect(t, "cccccccccccaan", format("c11", "a2", "n1"))
	Expect(t, "", format("8", "18", "218"))
	Expect(t, "", format("a", "bb", "cccc"))
	Expect(t, "", format("ab12", "aa3", ""))
}

func BenchmarkFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		format("n5", "n5", "c11", "n2")
	}
}

func TestBban(t *testing.T) {
	p := New().Payment()
	bban := p.bban("nnnaaaaccaannccc")
	matched, err := regexp.MatchString(`^\d{3}[a-z]{4}[a-z0-9]{2}[a-z]{2}\d{2}[a-z0-9]{3}$`, bban)
	Expect(t, nil, err)
	Expect(t, true, matched)
}

func BenchmarkBban(b *testing.B) {
	p := New().Payment()
	for i := 0; i < b.N; i++ {
		_ = p.bban("nnnnnaaaaacccccccccccnn")
	}
}
