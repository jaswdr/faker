package faker

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var cardVendors = []string{
	"Visa", "Visa", "Visa", "Visa", "Visa",
	"MasterCard", "MasterCard", "MasterCard", "MasterCard", "MasterCard",
	"American Express", "Discover Card", "Visa Retired",
}

type cardBrandConfig struct {
	prefixes []string
	length   int
}

var cardBrandConfigs = map[string]cardBrandConfig{
	"Visa": {
		prefixes: []string{"4"},
		length:   16,
	},
	"MasterCard": {
		prefixes: []string{"51", "52", "53", "54", "55"},
		length:   16,
	},
	"American Express": {
		prefixes: []string{"34", "37"},
		length:   15,
	},
	"Discover Card": {
		prefixes: []string{"6011", "65"},
		length:   16,
	},
	"Visa Retired": {
		prefixes: []string{"4"},
		length:   13,
	},
}

// Payment is a faker struct for Payment
type Payment struct {
	Faker *Faker
}

// CreditCardType returns a fake credit card type for Payment
func (p Payment) CreditCardType() string {
	return p.Faker.RandomStringElement(cardVendors)
}

// CreditCardNumber returns a Luhn-valid fake credit card number for a random card type.
func (p Payment) CreditCardNumber() string {
	return p.CreditCardNumberForType(p.CreditCardType())
}

// CreditCardNumberForType returns a Luhn-valid fake credit card number for the given card type.
func (p Payment) CreditCardNumberForType(cardType string) string {
	config, ok := cardBrandConfigs[cardType]
	if !ok {
		config = cardBrandConfigs["Visa"]
	}

	prefix := p.Faker.RandomStringElement(config.prefixes)
	return generateLuhnNumber(*p.Faker, prefix, config.length)
}

// generateLuhnNumber builds a Luhn-valid number with the given prefix and total length.
func generateLuhnNumber(f Faker, prefix string, length int) string {
	if length <= len(prefix) {
		return prefix
	}

	remaining := length - len(prefix) - 1
	partial := prefix + f.Numerify(strings.Repeat("#", remaining))
	checkDigit := luhnCheckDigit(partial)
	return partial + strconv.Itoa(checkDigit)
}

// luhnCheckDigit computes the Luhn check digit for a partial card number (without check digit).
func luhnCheckDigit(partial string) int {
	sum := 0
	for i := 0; i < len(partial); i++ {
		d := int(partial[len(partial)-1-i] - '0')
		if i%2 == 0 {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}
	return (10 - (sum % 10)) % 10
}

// IsValidLuhn returns true if the given numeric string passes the Luhn checksum.
func IsValidLuhn(number string) bool {
	sum := 0
	for i := 0; i < len(number); i++ {
		if number[len(number)-1-i] < '0' || number[len(number)-1-i] > '9' {
			return false
		}
		d := int(number[len(number)-1-i] - '0')
		if i%2 == 1 {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}
	return sum%10 == 0
}

// CreditCardExpirationDateString returns a fake credit card expiration date in string format (i.e. mm/yy) for Payment
func (p Payment) CreditCardExpirationDateString() string {
	month := p.Faker.IntBetween(1, 12)
	currentYear := time.Now().Year()
	year := p.Faker.IntBetween(currentYear, currentYear+3)

	return fmt.Sprintf("%02d/%02d", month, year%100)
}
