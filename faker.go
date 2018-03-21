package faker

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Faker struct {
	Generator *rand.Rand
}

func (f Faker) RandomDigit() int {
	return f.Generator.Int() % 10
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

func (f Faker) RandomElements(s interface{}, count int) interface{} {
	//@TODO
	return s
}

func (f Faker) RandomElement(s interface{}) interface{} {
	//@TODO
	return s
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
