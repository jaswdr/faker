package faker

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func Expect(t *testing.T, expected, got interface{}, values ...interface{}) {
	t.Helper()
	if expected != got {
		t.Errorf("\nExpected: (%T) %v \nGot:\t  (%T) %v", expected, expected, got, got)
		if len(values) > 0 {
			for _, v := range values {
				t.Errorf("\n%+v", v)
			}
		}

		t.FailNow()
	}
}

func NotExpect(t *testing.T, notExpected, got interface{}, values ...interface{}) {
	t.Helper()
	if notExpected == got {
		t.Errorf("\nNot Expecting: (%T) '%v', but it was", notExpected, notExpected)
		if len(values) > 0 {
			for _, v := range values {
				t.Errorf("\n%+v", v)
			}
		}

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

func TestInt(t *testing.T) {
	f := New()
	value := f.Int()
	Expect(t, fmt.Sprintf("%T", value), "int")
}

func TestInt64(t *testing.T) {
	f := New()
	value := f.Int64()
	Expect(t, fmt.Sprintf("%T", value), "int64")
}

func TestInt32(t *testing.T) {
	f := New()
	value := f.Int32()
	Expect(t, fmt.Sprintf("%T", value), "int32")
}

func TestIntBetween(t *testing.T) {
	f := New()
	value := f.IntBetween(1, 100)
	Expect(t, fmt.Sprintf("%T", value), "int")
	Expect(t, true, value >= 1)
	Expect(t, true, value <= 100)
}

func TestIntBetweenCanGenerateFirstElementInFirst100GeneratedValues(t *testing.T) {
	f := New()
	foundZero := false
	for i := 0; i < 100; i++ {
		if f.IntBetween(0, 1) == 0 {
			foundZero = true
			break
		}
	}
	Expect(t, true, foundZero)
}

func TestIntBetweenCanGenerateLastElementInFirst100GeneratedValues(t *testing.T) {
	f := New()
	foundOne := false
	for i := 0; i < 100; i++ {
		if f.IntBetween(0, 1) == 1 {
			foundOne = true
			break
		}
	}
	Expect(t, true, foundOne)
}

func TestRandomFloat(t *testing.T) {
	f := New()
	value := f.RandomFloat(1, 1, 100)
	Expect(t, fmt.Sprintf("%T", value), "float64")
	Expect(t, true, value >= 1)
	Expect(t, true, value <= 100)
}

func TestLetter(t *testing.T) {
	f := New()
	value := f.Letter()
	Expect(t, fmt.Sprintf("%T", value), "string")
	Expect(t, 1, len(value))
}

func TestRandomLetter(t *testing.T) {
	f := New()
	value := f.RandomLetter()
	Expect(t, fmt.Sprintf("%T", value), "string")
	Expect(t, 1, len(value))
}

func TestRandomStringWithLength(t *testing.T) {
	f := New()
	length := f.IntBetween(97, 1000)
	value := f.RandomStringWithLength(length)
	Expect(t, fmt.Sprintf("%T", value), "string")
	Expect(t, length, len(value))
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

func TestBoolWithChance(t *testing.T) {
	f := New()
	tp := reflect.TypeOf(f.BoolWithChance(30))
	Expect(t, "bool", tp.String())

	Expect(t, true, f.BoolWithChance(100))
	Expect(t, false, f.BoolWithChance(0))
	Expect(t, true, f.BoolWithChance(101))
	Expect(t, false, f.BoolWithChance(-1))
}

func TestMap(t *testing.T) {
	f := New()
	mp := f.Map()
	Expect(t, true, len(mp) > 0)
}

func TestRandomStringMapKey(t *testing.T) {
	f := New()
	m := map[string]string{"k0": "v0", "k1": "v1"}
	key := f.RandomStringMapKey(m)
	Expect(t, true, key == "k0" || key == "k1")
}

func TestRandomStringMapValue(t *testing.T) {
	f := New()
	m := map[string]string{"k0": "v0", "k1": "v1"}
	key := f.RandomStringMapValue(m)
	Expect(t, true, key == "v0" || key == "v1")
}
