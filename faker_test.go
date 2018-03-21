package faker

import (
	"fmt"
	"math/rand"
	"testing"
)

func Expect(t *testing.T, expected, got interface{}) {
	t.Helper()
	if expected != got {
		t.Errorf("\nExpected: (%T) %v \nGot:\t  (%T) %v", expected, expected, got, got)
		t.FailNow()
	}
}

func TestNew(t *testing.T) {
	f := New()
	Expect(t, fmt.Sprintf("%T", f), "faker.Faker")
}

func TestNewWithSeed(t *testing.T) {
	seed := rand.NewSource(0)
	f := NewWithSeed(seed)
	Expect(t, fmt.Sprintf("%T", f), "faker.Faker")
}

func TestRandomDigit(t *testing.T) {
	f := New()
	value := f.RandomDigit()
	Expect(t, fmt.Sprintf("%T", value), "int")
	Expect(t, true, value >= 0)
	Expect(t, true, value < 10)
}

func TestRandomDigitNotNull(t *testing.T) {
	f := New()
	value := f.RandomDigitNotNull()
	Expect(t, fmt.Sprintf("%T", value), "int")
	Expect(t, true, value > 0)
	Expect(t, true, value <= 9)
}

func TestRandomNumber(t *testing.T) {
	f := New()
	value := f.RandomNumber(4)
	Expect(t, fmt.Sprintf("%T", value), "int")
	Expect(t, true, value >= 1000)
	Expect(t, true, value <= 9999)
}

func TestNumberBetween(t *testing.T) {
	f := New()
	value := f.NumberBetween(1, 100)
	Expect(t, fmt.Sprintf("%T", value), "int")
	Expect(t, true, value >= 1)
	Expect(t, true, value <= 100)
}

func TestRandomFloat(t *testing.T) {
	f := New()
	value := f.RandomFloat(1, 1, 100)
	Expect(t, fmt.Sprintf("%T", value), "float64")
	Expect(t, true, value >= 1)
	Expect(t, true, value <= 100)
}

func TestRandomLetter(t *testing.T) {
	f := New()
	value := f.RandomLetter()
	Expect(t, fmt.Sprintf("%T", value), "string")
	Expect(t, 1, len(value))
}
