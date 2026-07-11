package faker

import (
	"testing"
	"time"
)

func TestUniqueIntBetween(t *testing.T) {
	u := NewWithSeedInt64(1).Unique()
	seen := make(map[int]struct{})
	for i := 0; i < 50; i++ {
		val := u.IntBetween(1, 50)
		if _, exists := seen[val]; exists {
			t.Fatalf("duplicate value %d", val)
		}
		seen[val] = struct{}{}
	}
}

func TestUniqueStringElement(t *testing.T) {
	elements := []string{"a", "b", "c", "d", "e"}
	u := New().Unique()
	seen := make(map[string]struct{})
	for i := 0; i < len(elements); i++ {
		val := u.StringElement(elements)
		if _, exists := seen[val]; exists {
			t.Fatalf("duplicate value %s", val)
		}
		seen[val] = struct{}{}
	}
}

func TestUniqueReset(t *testing.T) {
	u := NewWithSeedInt64(1).Unique()
	for i := 0; i < 5; i++ {
		u.IntBetween(1, 5)
	}
	u.Reset()
	val := u.IntBetween(1, 5)
	Expect(t, true, val >= 1 && val <= 5)
}

func TestBarcodeISBN13(t *testing.T) {
	b := New().Barcode()
	isbn := b.ISBN13()
	Expect(t, 13, len(isbn))
	Expect(t, true, IsValidISBN13(isbn))
	Expect(t, true, isbn[0:3] == "978" || isbn[0:3] == "979")
}

func TestBarcodeEAN13(t *testing.T) {
	b := New().Barcode()
	ean := b.EAN13()
	Expect(t, 13, len(ean))
	Expect(t, true, IsValidEAN13(ean))
}

func TestPersonDateOfBirth(t *testing.T) {
	p := NewWithSeedInt64(42).Person()
	dob := p.DateOfBirth(18, 65)
	now := time.Now()
	age := now.Year() - dob.Year()
	if dob.AddDate(age, 0, 0).After(now) {
		age--
	}
	Expect(t, true, age >= 18 && age <= 65)
}
