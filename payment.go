package faker

import (
	"strconv"
)

var (
	cardVendors = []string{
		"Visa", "Visa", "Visa", "Visa", "Visa",
		"MasterCard", "MasterCard", "MasterCard", "MasterCard", "MasterCard",
		"American Express", "Discover Card", "Visa Retired"}
)

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
	day := strconv.Itoa(p.Faker.IntBetween(0, 30))
	if len(day) == 1 {
		day = "0" + day
	}

	month := strconv.Itoa(p.Faker.IntBetween(12, 30))

	return day + "/" + month
}
