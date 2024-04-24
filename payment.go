package faker

import (
	"fmt"
	"time"
)

var cardVendors = []string{
	"Visa", "Visa", "Visa", "Visa", "Visa",
	"MasterCard", "MasterCard", "MasterCard", "MasterCard", "MasterCard",
	"American Express", "Discover Card", "Visa Retired",
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

// CreditCardExpirationDateString returns a fake credit card expiration date in string format (i.e. mm/yy) for Payment
func (p Payment) CreditCardExpirationDateString() string {
	month := p.Faker.IntBetween(1, 12)
	currentYear := time.Now().Year()
	year := p.Faker.IntBetween(currentYear, currentYear+3)

	return fmt.Sprintf("%02d/%02d", month, year%100)
}
