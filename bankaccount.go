package faker

import (
	"strings"
)

// BAccount is a faker struct for BAccount
type BAccount struct {
	Faker *Faker
}

var accountLengths = map[string]int{
	"US": 10, // United States
	"UK": 14, // United Kingdom
	"CA": 12, // Canada
	"AU": 9,  // Australia
	"JP": 7,  // Japan
	"IN": 15, // India
	"BR": 12, // Brazil
	"CN": 19, // China
	"MX": 18, // Mexico
	"RU": 20, // Russia
	"ZA": 14, // South Africa
	"KR": 14, // South Korea
	"ID": 16, // Indonesia
	"NG": 10, // Nigeria
	"SG": 10, // Singapore
	"MY": 14, // Malaysia
	"TH": 10, // Thailand
	"":   15,
	// Add more countries and their corresponding lengths here.
}

func (f BAccount) GenerateBankAccount(countryCode string) int {
	cc := strings.ToUpper(countryCode)
	if _, exists := accountLengths[cc]; exists {
        AccountLength := accountLengths[cc]
	   return f.Faker.RandomNumber(AccountLength)
	}else{
		AccountLength := accountLengths[""]
	return f.Faker.RandomNumber(AccountLength)
	}
	
}
