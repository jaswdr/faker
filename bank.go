package faker

import (
	"strings"
)

type Bank struct {
	Faker *Faker
}

var (
	bankNames      = []string{"JP Morgan Chase", "Bank of America", "Wells Fargo", "Citigroup", "Goldman Sachs", "Morgan Stanley", "Barclays", "Deutsche Bank", "UBS", "Credit Suisse", "HSBC", "RBS", "NatWest", "Lloyds Bank", "Barclays", "Deutsche Bank", "UBS", "Credit Suisse", "HSBC", "RBS", "NatWest", "Lloyds Bank"}
	swiftCodes     = []string{"CHASUS33###", "BOFAUS6S###", "WFBIUS6S###", "CITIUS33###", "GSBCUS33###", "MSBCUS33###", "BARCUS33###", "DBKCUS33###", "UBSUS33###", "CSUS33###", "HSBCUS33###", "RBSUS33###", "NWBKUS33###", "LLOYDGB21###", "BARCUS33###", "DBKCUS33###", "UBSUS33###", "CSUS33###", "HSBCUS33###", "RBSUS33###", "NWBKUS33###", "LLOYDGB21###"}
	countryFormats = map[string]string{
		"GB": "GB##NWBK############",
		"DE": "DE##########",
		"FR": "FR##############",
		"IT": "IT##X###########",
		"ES": "ES####################",
	}
)

// Name returns a random bank name
func (b Bank) Name() string {
	return b.Faker.RandomStringElement(bankNames)
}

// SwiftCode returns a random swift code
func (b Bank) SwiftCode() string {
	return b.Faker.Numerify(b.Faker.RandomStringElement(swiftCodes))
}

// IBAN returns a random IBAN (International Bank Account Number)
func (b Bank) IBAN() string {
	countries := make([]string, 0, len(countryFormats))
	for k := range countryFormats {
		countries = append(countries, k)
	}

	country := countries[b.Faker.Generator.Intn(len(countries))]
	format := countryFormats[country]
	iban := b.Faker.Numerify(format)
	iban = strings.Replace(iban, "X", string(rune('A'+b.Faker.Generator.Intn(26))), -1)
	return iban
}
