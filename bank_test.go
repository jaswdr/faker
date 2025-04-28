package faker

import "testing"

func TestBankName(t *testing.T) {
	f := New()
	name := f.Bank().Name()

	NotExpect(t, "", name)
	found := false
	for _, bankName := range bankNames {
		if name == bankName {
			found = true
			break
		}
	}
	Expect(t, true, found)
}

func TestBankSwiftCode(t *testing.T) {
	f := New()
	swiftCode := f.Bank().SwiftCode()

	NotExpect(t, "", swiftCode)
	Expect(t, true, len(swiftCode) >= 9)

	found := false
	for _, pattern := range swiftCodes {
		if len(swiftCode) == len(pattern) && swiftCode[:8] == pattern[:8] {
			found = true
			for i := 8; i < len(swiftCode); i++ {
				if swiftCode[i] < '0' || swiftCode[i] > '9' {
					found = false
					break
				}
			}
			break
		}
	}
	Expect(t, true, found)
}

func TestBankIBAN(t *testing.T) {
	f := New()
	iban := f.Bank().IBAN()

	NotExpect(t, "", iban)

	countryCode := iban[:2]
	format, exists := countryFormats[countryCode]
	Expect(t, true, exists)

	Expect(t, len(format), len(iban))

	for i := 2; i < len(iban); i++ {
		if format[i] == '#' {
			Expect(t, true, iban[i] >= '0' && iban[i] <= '9')
		} else if format[i] == 'X' {
			Expect(t, true, iban[i] >= 'A' && iban[i] <= 'Z')
		}
	}
}
