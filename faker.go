package faker

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Faker is the Generator struct for Faker
type Faker struct {
	Generator *rand.Rand
}

// RandomDigit returns a fake random digit for Faker
func (f Faker) RandomDigit() int {
	return f.Generator.Int() % 10
}

// RandomDigitNot returns a fake random digit for Faker that is not in a list of ignored
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

// RandomDigitNotNull returns a fake random digit that is not null for Faker
func (f Faker) RandomDigitNotNull() int {
	return f.Generator.Int()%8 + 1
}

// RandomNumber returns a fake random integer number for Faker
func (f Faker) RandomNumber(size int) int {
	if size == 1 {
		return f.RandomDigit()
	}

	var min int = int(math.Pow10(size - 1))
	var max int = int(math.Pow10(size)) - 1

	return f.IntBetween(min, max)
}

// RandomFloat returns a fake random float number for Faker
func (f Faker) RandomFloat(maxDecimals, min, max int) float64 {
	s := fmt.Sprintf("%d.%d", f.IntBetween(min, max-1), f.IntBetween(1, maxDecimals))
	value, _ := strconv.ParseFloat(s, 10)
	return value
}

// Int returns a fake Int number for Faker
func (f Faker) Int() int {
	maxU := ^uint(0) >> 1
	max := int(maxU)
	min := -max - 1
	return f.IntBetween(min, max)
}

// Int64 returns a fake Int64 number for Faker
func (f Faker) Int64() int64 {
	return int64(f.Int())
}

// Int32 returns a fake Int32 number for Faker
func (f Faker) Int32() int32 {
	return int32(f.Int())
}

// IntBetween returns a fake Int between a given minimum and maximum values for Faker
func (f Faker) IntBetween(min, max int) int {
	diff := max - min

	if diff == 0 {
		return min
	}

	return min + (f.Generator.Int() % diff)
}

// Int64Between returns a fake Int64 between a given minimum and maximum values for Faker
func (f Faker) Int64Between(min, max int64) int64 {
	return int64(f.IntBetween(int(min), int(max)))
}

// Int32Between returns a fake Int32 between a given minimum and maximum values for Faker
func (f Faker) Int32Between(min, max int32) int32 {
	return int32(f.IntBetween(int(min), int(max)))
}

// Letter returns a fake single letter for Faker
func (f Faker) Letter() string {
	return f.RandomLetter()
}

// RandomLetter returns a fake random string with a random number of letters for Faker
func (f Faker) RandomLetter() string {
	return fmt.Sprintf("%c", f.IntBetween(97, 122))
}

// RandomStringElement returns a fake random string element from a given list of strings for Faker
func (f Faker) RandomStringElement(s []string) string {
	i := f.IntBetween(0, len(s))
	return s[i]
}

// RandomIntElement returns a fake random int element form a given list of ints for Faker
func (f Faker) RandomIntElement(a []int) int {
	i := f.IntBetween(0, len(a))
	return a[i]
}

// ShuffleString returns a fake shuffled string from a given string for Faker
func (f Faker) ShuffleString(s string) string {
	orig := strings.Split(s, "")
	dest := make([]string, len(orig))

	for i := 0; i < len(orig); i++ {
		dest[i] = orig[len(orig)-i-1]
	}

	return strings.Join(dest, "")
}

// Numerify returns a fake string that replace all "#" characters with numbers from a given string for Faker
func (f Faker) Numerify(in string) (out string) {
	for _, c := range strings.Split(in, "") {
		if c == "#" {
			c = strconv.Itoa(f.RandomDigit())
		}

		out = out + c
	}

	return
}

// Lexify  returns a fake string that replace all "?" characters with random letters from a given string for Faker
func (f Faker) Lexify(in string) (out string) {
	for _, c := range strings.Split(in, "") {
		if c == "?" {
			c = f.RandomLetter()
		}

		out = out + c
	}

	return
}

// Bothify returns a fake string that apply Lexify() and Numerify() on a given string for Faker
func (f Faker) Bothify(in string) (out string) {
	out = f.Lexify(in)
	out = f.Numerify(out)
	return
}

// Asciify   returns a fake string that replace all "*" characters with random ASCII values from a given string for Faker
func (f Faker) Asciify(in string) (out string) {
	for _, c := range strings.Split(in, "") {
		if c == "*" {
			c = fmt.Sprintf("%c", f.IntBetween(97, 126))
		}

		out = out + c
	}

	return
}

// Lorem returns a fake Lorem instance for Faker
func (f Faker) Boolean() Boolean {
	return Boolean{&f}
}

// Lorem returns a fake Lorem instance for Faker
func (f Faker) Lorem() Lorem {
	return Lorem{&f}
}

// Person returns a fake Person instance for Faker
func (f Faker) Person() Person {
	return Person{&f}
}

// Address returns a fake Address instance for Faker
func (f Faker) Address() Address {
	return Address{&f}
}

// Phone returns a fake Phone instance for Faker
func (f Faker) Phone() Phone {
	return Phone{&f}
}

// Company returns a fake Company instance for Faker
func (f Faker) Company() Company {
	return Company{&f}
}

// Time returns a fake Time instance for Faker
func (f Faker) Time() Time {
	return Time{&f}
}

// Internet returns a fake Internet instance for Faker
func (f Faker) Internet() Internet {
	return Internet{&f}
}

// UserAgent returns a fake UserAgent instance for Faker
func (f Faker) UserAgent() UserAgent {
	return UserAgent{&f}
}

// Payment returns a fake Payment instance for Faker
func (f Faker) Payment() Payment {
	return Payment{&f}
}

// MimeType returns a fake MimeType instance for Faker
func (f Faker) MimeType() MimeType {
	return MimeType{&f}
}

// Color returns a fake Color instance for Faker
func (f Faker) Color() Color {
	return Color{&f}
}

// UUID returns a fake UUID instance for Faker
func (f Faker) UUID() UUID {
	return UUID{&f}
}

// Image returns a fake Image instance for Faker
func (f Faker) Image() Image {
	return Image{&f}
}

// YouTube returns a fake YouTube instance for Faker
func (f Faker) YouTube() YouTube {
	return YouTube{&f}
}

// New returns a new instance of Faker instance with a random seed
func New() (f Faker) {
	seed := rand.NewSource(time.Now().Unix())
	f = NewWithSeed(seed)
	return
}

// NewWithSeed returns a new instance of Faker instance with a given seed
func NewWithSeed(src rand.Source) (f Faker) {
	generator := rand.New(src)
	f = Faker{Generator: generator}
	return
}
