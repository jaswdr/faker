package faker

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func Expect(t *testing.T, expected, got interface{}) {
	t.Helper()
	if expected != got {
		t.Errorf("\nExpected: (%T) %v \nGot:\t  (%T) %v", expected, expected, got, got)
		t.FailNow()
	}
}

func F(t *testing.T) Faker {
	return NewWithSeed(rand.NewSource(0))
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

func TestRandomDigitNot(t *testing.T) {
	f := New()
	value := f.RandomDigitNot(1)
	Expect(t, fmt.Sprintf("%T", value), "int")
	Expect(t, true, value != 1)
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

func TestIntBetween(t *testing.T) {
	f := New()
	value := f.IntBetween(1, 100)
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

func TestRandomIntElement(t *testing.T) {
	f := New()
	elements := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	element := f.RandomIntElement(elements)
	found := false
	for _, i := range elements {
		if i == element {
			found = true
		}
	}
	Expect(t, true, found)
}

func TestShuffleString(t *testing.T) {
	f := New()
	orig := "foo bar"
	returned := f.ShuffleString("foo bar")
	Expect(t, len(orig), len(returned))
	for _, s := range strings.Split(returned, "") {
		Expect(t, true, strings.Contains(orig, s))
	}
}

func TestNumerify(t *testing.T) {
	f := New()
	value := f.Numerify("Hello ##?#")
	Expect(t, 10, len(value))
	Expect(t, true, strings.Contains(value, "Hello"))
	Expect(t, true, strings.Contains(value, "?"))
	Expect(t, false, strings.Contains(value, "#"))
}

func TestLexify(t *testing.T) {
	f := New()
	value := f.Lexify("Hello ??#?")
	Expect(t, 10, len(value))
	Expect(t, true, strings.Contains(value, "Hello"))
	Expect(t, true, strings.Contains(value, "#"))
	Expect(t, false, strings.Contains(value, "?"))
}

func TestBothify(t *testing.T) {
	f := New()
	value := f.Bothify("Hello ??#?")
	Expect(t, 10, len(value))
	Expect(t, true, strings.Contains(value, "Hello"))
	Expect(t, false, strings.Contains(value, "#"))
	Expect(t, false, strings.Contains(value, "?"))
}

func TestAsciify(t *testing.T) {
	f := New()
	value := f.Asciify("Hello ??#?****")
	Expect(t, 14, len(value))
	Expect(t, true, strings.Contains(value, "Hello"))
	Expect(t, true, strings.Contains(value, "#"))
	Expect(t, true, strings.Contains(value, "?"))
	Expect(t, false, strings.Contains(value, "*"))
}

func TestBool(t *testing.T) {
	f := New()
	tp := reflect.TypeOf(f.Bool())
	Expect(t, "bool", tp.String())
}
