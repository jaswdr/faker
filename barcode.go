package faker

import (
	"strconv"
	"strings"
)

// Barcode is a faker struct for barcode identifiers.
type Barcode struct {
	Faker *Faker
}

// ISBN13 returns a valid ISBN-13 barcode string.
func (b Barcode) ISBN13() string {
	prefix := b.Faker.RandomStringElement([]string{"978", "979"})
	body := prefix + b.Faker.Numerify("#########")
	checkDigit := isbn13CheckDigit(body)
	return body + strconv.Itoa(checkDigit)
}

// EAN13 returns a valid EAN-13 barcode string.
func (b Barcode) EAN13() string {
	body := b.Faker.Numerify("############")
	checkDigit := ean13CheckDigit(body[:12])
	return body[:12] + strconv.Itoa(checkDigit)
}

func isbn13CheckDigit(digits string) int {
	sum := 0
	for i, c := range digits {
		d := int(c - '0')
		if i%2 == 0 {
			sum += d
		} else {
			sum += d * 3
		}
	}
	return (10 - (sum % 10)) % 10
}

func ean13CheckDigit(digits string) int {
	return isbn13CheckDigit(digits)
}

// IsValidISBN13 returns true if the given string is a valid ISBN-13.
func IsValidISBN13(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 13 {
		return false
	}
	for _, c := range isbn {
		if c < '0' || c > '9' {
			return false
		}
	}
	body := isbn[:12]
	check, err := strconv.Atoi(string(isbn[12]))
	if err != nil {
		return false
	}
	return isbn13CheckDigit(body) == check
}

// IsValidEAN13 returns true if the given string is a valid EAN-13.
func IsValidEAN13(ean string) bool {
	return IsValidISBN13(ean)
}
