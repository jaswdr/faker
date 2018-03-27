package faker

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Faker struct {
	Generator *rand.Rand
}

func (f Faker) RandomDigit() int {
	return f.Generator.Int() % 10
}

func (f Faker) RandomDigitNot(ignore ...int) int {
	inSlice := func(el int, list []int) bool {
		for i := range list {
			if i == el {
				return true
			}
		}

		return false
	}

	for {
		current := f.RandomDigit()
		if inSlice(current, ignore) {
			return current
		}
	}
}

func (f Faker) RandomDigitNotNull() int {
	return f.Generator.Int()%8 + 1
}

func (f Faker) RandomNumber(size int) int {
	if size == 1 {
		return f.RandomDigit()
	}

	var min int = int(math.Pow10(size - 1))
	var max int = int(math.Pow10(size)) - 1

	return f.NumberBetween(min, max)
}

func (f Faker) RandomFloat(maxDecimals, min, max int) float64 {
	s := fmt.Sprintf("%d.%d", f.NumberBetween(min, max-1), f.NumberBetween(1, maxDecimals))
	value, _ := strconv.ParseFloat(s, 10)
	return value
}

func (f Faker) NumberBetween(min, max int) int {
	step := 1
	for i := 0; i < f.RandomDigitNotNull(); i++ {
		step = step * f.RandomDigitNotNull()
	}

	var value int = min + step

	if value > max {
		return max
	}

	return value
}

func (f Faker) RandomLetter() string {
	return string(f.NumberBetween(97, 122))
}

func (f Faker) RandomStringElement(s []string) string {
	i := f.NumberBetween(0, len(s)-1)
	return s[i]
}

func (f Faker) ShuffleString(s string) string {
	orig := strings.Split(s, "")
	dest := make([]string, len(orig))

	for i := 0; i < len(orig); i++ {
		dest[i] = orig[len(orig)-i-1]
	}

	return strings.Join(dest, "")
}

func (f Faker) Numerify(in string) (out string) {
	for _, c := range strings.Split(in, "") {
		if c == "#" {
			c = strconv.Itoa(f.RandomDigit())
		}

		out = out + c
	}

	return
}

func (f Faker) Lexify(in string) (out string) {
	for _, c := range strings.Split(in, "") {
		if c == "?" {
			c = f.RandomLetter()
		}

		out = out + c
	}

	return
}

func (f Faker) Bothify(in string) (out string) {
	out = f.Lexify(in)
	out = f.Numerify(out)
	return
}

func (f Faker) Asciify(in string) (out string) {
	for _, c := range strings.Split(in, "") {
		if c == "*" {
			c = string(f.NumberBetween(33, 126))
		}

		out = out + c
	}

	return
}

func (f Faker) Lorem() Lorem {
	return Lorem{&f}
}

func (f Faker) Person() Person {
	return Person{&f}
}

func (f Faker) Address() Address {
	return Address{&f}
}

func (f Faker) Phone() Phone {
	return Phone{&f}
}

func (f Faker) Company() Company {
	return Company{&f}
}

func New() (f Faker) {
	seed := rand.NewSource(time.Now().Unix())
	f = NewWithSeed(seed)
	return
}

func NewWithSeed(src rand.Source) (f Faker) {
	generator := rand.New(src)
	f = Faker{Generator: generator}
	return
}
