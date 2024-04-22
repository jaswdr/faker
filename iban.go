package faker

import (
	"fmt"
	"strconv"
	"strings"
)

func ibanChecksum(iban string) string {
	iban = strings.ToUpper(strings.ReplaceAll(iban, " ", ""))

	// Move first 4 characters to the end, and set checksum to 00
	iban = iban[4:] + iban[:2] + "00"

	// Replace letters with their respective numbers
	var numericIBAN string
	for _, char := range iban {
		if char >= 'A' && char <= 'Z' {
			numericIBAN += strconv.Itoa(int(char - 'A' + 10))
		} else {
			numericIBAN += string(char)
		}
	}

	// Perform modulo 97 operation on the numeric IBAN
	remainder := 0
	for _, char := range numericIBAN {
		digit := int(char - '0')
		remainder = (remainder*10 + digit) % 97
	}

	// Calculate checksum
	checksum := 98 - remainder
	if checksum < 10 {
		return fmt.Sprintf("0%d", checksum)
	}
	return fmt.Sprintf("%d", checksum)
}

func isIbanValid(iban string) bool {
	return ibanChecksum(iban) == iban[2:4]
}
