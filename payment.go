package faker

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"unicode"
)

var cardVendors = []string{
	"Visa", "Visa", "Visa", "Visa", "Visa",
	"MasterCard", "MasterCard", "MasterCard", "MasterCard", "MasterCard",
	"American Express", "Discover Card", "Visa Retired",
}

// ibanFormats lists all IBAN formats, source: @link https://www.swift.com/swift-resource/9606/download
// n: numeric, a: alphabetic, c: alphanumeric
var ibanFormats = map[string]string{
	"AD": format("n4", "n4", "c12"),
	"AE": format("n3", "n16"),
	"AL": format("n8", "c16"),
	"AT": format("n5", "n11"),
	"AZ": format("a4", "c20"),
	"BA": format("n3", "n3", "n8", "n2"),
	"BE": format("n3", "n7", "n2"),
	"BG": format("a4", "n4", "n2", "c8"),
	"BH": format("a4", "c14"),
	"BR": format("n8", "n5", "n10", "a1", "c1"),
	"CH": format("n5", "c12"),
	"CR": format("n4", "n14"),
	"CY": format("n3", "n5", "c16"),
	"CZ": format("n4", "n6", "n10"),
	"DE": format("n8", "n10"),
	"DK": format("n4", "n9", "n1"),
	"DO": format("c4", "n20"),
	"EE": format("n2", "n2", "n11", "n1"),
	"ES": format("n4", "n4", "n1", "n1", "n10"),
	"FI": format("n6", "n7", "n1"),
	"FR": format("n5", "n5", "c11", "n2"),
	"GB": format("a4", "n6", "n8"),
	"GE": format("a2", "n16"),
	"GI": format("a4", "c15"),
	"GR": format("n3", "n4", "c16"),
	"GT": format("c4", "c20"),
	"HR": format("n7", "n10"),
	"HU": format("n3", "n4", "n1", "n15", "n1"),
	"IE": format("a4", "n6", "n8"),
	"IL": format("n3", "n3", "n13"),
	"IS": format("n4", "n2", "n6", "n10"),
	"IT": format("a1", "n5", "n5", "c12"),
	"KW": format("a4", "n22"),
	"KZ": format("n3", "c13"),
	"LB": format("n4", "c20"),
	"LI": format("n5", "c12"),
	"LT": format("n5", "n11"),
	"LU": format("n3", "c13"),
	"LV": format("a4", "c13"),
	"MC": format("n5", "n5", "c11", "n2"),
	"MD": format("c2", "c18"),
	"ME": format("n3", "n13", "n2"),
	"MK": format("n3", "c10", "n2"),
	"MR": format("n5", "n5", "n11", "n2"),
	"MT": format("a4", "n5", "c18"),
	"MU": format("a4", "n2", "n2", "n12", "n3", "a3"),
	"NL": format("a4", "n10"),
	"NO": format("n4", "n6", "n1"),
	"PK": format("a4", "c16"),
	"PL": format("n8", "n16"),
	"PS": format("a4", "c21"),
	"PT": format("n4", "n4", "n11", "n2"),
	"RO": format("a4", "c16"),
	"RS": format("n3", "n13", "n2"),
	"SA": format("n2", "c18"),
	"SE": format("n3", "n16", "n1"),
	"SI": format("n5", "n8", "n2"),
	"SK": format("n4", "n6", "n10"),
	"SM": format("a1", "n5", "n5", "c12"),
	"TN": format("n2", "n3", "n13", "n2"),
	"TR": format("n5", "n1", "c16"),
	"VG": format("a4", "n16"),
}

// format interprets the format of each section of the iban and returns a string with a specific format
// Example: format("n5", "a2", "c1") => "nnnnnaac"
func format(sections ...string) string {
	var res string
	for _, s := range sections {
		if len(s) == 0 {
			continue
		}
		if class := s[0]; unicode.IsLetter(rune(class)) {
			size, _ := strconv.Atoi(s[1:])
			res += strings.Repeat(string(class), size)
		}
	}
	return res
}

// Payment is a faker struct for Payment
type Payment struct {
	Faker *Faker
}

// CreditCardType returns a fake credit card type for Payment
func (p Payment) CreditCardType() string {
	return p.Faker.RandomStringElement(cardVendors)
}

// CreditCardNumber returns a fake credit card number for Payment
func (p Payment) CreditCardNumber() string {
	return p.Faker.Numerify("################")
}

// CreditCardExpirationDateString returns a fake credit card expiration date in string format for Payment
func (p Payment) CreditCardExpirationDateString() string {
	day := p.Faker.IntBetween(0, 30)
	month := p.Faker.IntBetween(12, 30)
	return fmt.Sprintf("%02d/%02d", day, month)
}

// Iban returns a fake IBAN for Payment
func (p Payment) Iban() string {
	return p.ibanForCountry(p.Faker.RandomStringMapKey(ibanFormats))
}

// ibanForCountry returns a fake IBAN for a specific country
// Returns an empty string if the country is not supported
func (p Payment) ibanForCountry(countryCode string) string {
	format, ok := ibanFormats[countryCode]
	if !ok {
		return ""
	}

	bban := strings.ToUpper(p.bban(format))
	checksum := ibanChecksum(countryCode + "00" + bban)

	return countryCode + checksum + bban
}

// bban generates a fake Basic Bank Account Number (BBAN) based on the format
// the provided format must be a string only containing the following characters:
// n: numeric, a: alphabetic, c: alphanumeric
// which will be replaced by a random number or letter
func (p Payment) bban(format string) string {
	format = strings.ReplaceAll(format, "n", "#")
	format = strings.ReplaceAll(format, "a", "?")
	s := "?"
	if rand.Intn(100) <= 50 {
		s = "#"
	}
	format = strings.ReplaceAll(format, "c", s)

	return p.Faker.Bothify(format)
}
