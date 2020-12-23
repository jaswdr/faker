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

// Float returns a fake random float number for Faker
func (f Faker) Float(maxDecimals, min, max int) float64 {
	s := fmt.Sprintf("%d.%d", f.IntBetween(min, max-1), f.IntBetween(1, maxDecimals))
	value, _ := strconv.ParseFloat(s, 10)
	return value
}

// Float32 returns a fake random float64 number for Faker
func (f Faker) Float32(maxDecimals, min, max int) float32 {
	s := fmt.Sprintf("%d.%d", f.IntBetween(min, max-1), f.IntBetween(1, maxDecimals))
	value, _ := strconv.ParseFloat(s, 10)
	return float32(value)
}

// Float64 returns a fake random float64 number for Faker
func (f Faker) Float64(maxDecimals, min, max int) float64 {
	s := fmt.Sprintf("%d.%d", f.IntBetween(min, max-1), f.IntBetween(1, maxDecimals))
	value, _ := strconv.ParseFloat(s, 10)
	return float64(value)
}

// Int returns a fake Int number for Faker
func (f Faker) Int() int {
	maxU := ^uint(0) >> 1
	max := int(maxU)
	min := -max - 2
	return f.IntBetween(min, max)
}

// Int8 returns a fake Int8 number for Faker
func (f Faker) Int8() int8 {
	return int8(f.Int())
}

// Int16 returns a fake Int16 number for Faker
func (f Faker) Int16() int16 {
	return int16(f.Int())
}

// Int32 returns a fake Int32 number for Faker
func (f Faker) Int32() int32 {
	return int32(f.Int())
}

// Int64 returns a fake Int64 number for Faker
func (f Faker) Int64() int64 {
	return int64(f.Int())
}

// UInt returns a fake UInt number for Faker
func (f Faker) UInt() uint {
	maxU := ^uint(0) >> 1
	max := int(maxU)
	return uint(f.IntBetween(0, max))
}

// UInt8 returns a fake UInt8 number for Faker
func (f Faker) UInt8() uint8 {
	return uint8(f.Int())
}

// UInt16 returns a fake UInt16 number for Faker
func (f Faker) UInt16() uint16 {
	return uint16(f.Int())
}

// UInt32 returns a fake UInt32 number for Faker
func (f Faker) UInt32() uint32 {
	return uint32(f.Int())
}

// UInt64 returns a fake UInt64 number for Faker
func (f Faker) UInt64() uint64 {
	return uint64(f.Int())
}

// IntBetween returns a fake Int between a given minimum and maximum values for Faker
func (f Faker) IntBetween(min, max int) int {
	diff := max - min

	if diff == 0 {
		return min
	}

	return rand.Intn(diff+1) + min
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
	i := f.IntBetween(0, len(s)-1)
	return s[i]
}

// RandomIntElement returns a fake random int element form a given list of ints for Faker
func (f Faker) RandomIntElement(a []int) int {
	i := f.IntBetween(0, len(a)-1)
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

// Bool returns a fake bool for Faker
func (f Faker) Bool() bool {
	return f.Boolean().Bool()
}

// BoolWithChance returns true with a given percentual chance that the value is true, otherwise returns false
func (f Faker) BoolWithChance(chanceTrue int) bool {
	return f.Boolean().BoolWithChance(chanceTrue)
}

// Boolean returns a fake Boolean instance for Faker
func (f Faker) Boolean() Boolean {
	return Boolean{&f}
}

// Map returns a fake Map instance for Faker
func (f Faker) Map() map[string]interface{} {
	m := map[string]interface{}{}
	lorem := f.Lorem()

	randWordType := func() string {
		s := f.RandomStringElement([]string{"lorem", "bs", "job", "name", "address"})
		switch s {
		case "bs":
			return f.Company().BS()
		case "job":
			return f.Company().JobTitle()
		case "name":
			return f.Person().Name()
		case "address":
			return f.Address().Address()
		}
		return lorem.Word()
	}

	randSlice := func() []string {
		var sl []string
		for ii := 0; ii < f.IntBetween(3, 10); ii++ {
			sl = append(sl, lorem.Word())
		}
		return sl
	}

	for i := 0; i < f.IntBetween(3, 10); i++ {
		t := f.RandomStringElement([]string{"string", "int", "float", "slice", "map"})
		switch t {
		case "string":
			m[lorem.Word()] = randWordType()
		case "int":
			m[lorem.Word()] = f.IntBetween(1, 10000000)
		case "float":
			m[lorem.Word()] = f.Float64(f.IntBetween(1, 4), 1, 1000000)
		case "slice":
			m[lorem.Word()] = randSlice()
		case "map":
			mm := map[string]interface{}{}
			tt := f.RandomStringElement([]string{"string", "int", "float", "slice"})
			switch tt {
			case "string":
				mm[lorem.Word()] = randWordType()
			case "int":
				mm[lorem.Word()] = f.IntBetween(1, 10000000)
			case "float":
				mm[lorem.Word()] = f.Float64(f.IntBetween(1, 4), 1, 1000000)
			case "slice":
				mm[lorem.Word()] = randSlice()
			}
			m[lorem.Word()] = mm
		}
	}

	return m
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

// File returns a fake File instance for Faker
func (f Faker) File() File {
	return File{&f}
}

// YouTube returns a fake YouTube instance for Faker
func (f Faker) YouTube() YouTube {
	return YouTube{&f}
}

// Struct returns a fake Struct instance for Faker
func (f Faker) Struct() Struct {
	return Struct{&f}
}

// Gamer returns a fake Gamer instance for Faker
func (f Faker) Gamer() Gamer {
	return Gamer{&f}
}

// Language returns a fake Language instance for Faker
func (f Faker) Language() Language {
	return Language{&f}
}

// Beer returns a fake Beer instance for Faker
func (f Faker) Beer() Beer {
	return Beer{&f}
}

// Car returns a fake Car instance for Faker
func (f Faker) Car() Car {
	return Car{&f}
}

// Food returns a fake Food instance for Faker
func (f Faker) Food() Food {
	return Food{&f}
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
