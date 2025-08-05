package faker

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Faker is the Generator struct for Faker
type Faker struct {
	Generator GeneratorInterface
}

// GeneratorInterface presents an Interface that allows us to subsequently control
// the returned value more accurately when doing tests by allowing us to use a struct that
// implements these methods to control the returned value. If not in tests, rand.Rand implements
// these methods and fufills the interface requirements.
type GeneratorInterface interface {
	Intn(n int) int
	Int32n(n int32) int32
	Int64n(n int64) int64
	Uintn(n uint) uint
	Uint32n(n uint32) uint32
	Uint64n(n uint64) uint64
	Int() int
}

// threadSafeRand wraps rand.Rand with a mutex for thread safety
type threadSafeRand struct {
	rand *rand.Rand
	mu   sync.Mutex
}

func (t *threadSafeRand) Intn(n int) int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.rand.IntN(n)
}

func (t *threadSafeRand) Int32n(n int32) int32 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.rand.Int32N(n)
}

func (t *threadSafeRand) Int64n(n int64) int64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.rand.Int64N(n)
}

func (t *threadSafeRand) Uintn(n uint) uint {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.rand.UintN(n)
}

func (t *threadSafeRand) Uint32n(n uint32) uint32 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.rand.Uint32N(n)
}

func (t *threadSafeRand) Uint64n(n uint64) uint64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.rand.Uint64N(n)
}

func (t *threadSafeRand) Int() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.rand.Int()
}

// RandomDigit returns a fake random digit for Faker
func (f Faker) RandomDigit() int {
	return f.Generator.Int() % 10
}

// RandomDigitNot returns a fake random digit for Faker that is not in a list of ignored
func (f Faker) RandomDigitNot(ignore ...int) int {
	inSlice := func(el int, list []int) bool {
		for _, v := range list {
			if v == el {
				return true
			}
		}

		return false
	}

	for {
		current := f.RandomDigit()
		if !inSlice(current, ignore) {
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

	minN := int(math.Pow10(size - 1))
	maxN := int(math.Pow10(size)) - 1

	return f.IntBetween(minN, maxN)
}

// generateFloat is a helper function that generates a random float with the specified parameters
func (f Faker) generateFloat(maxDecimals, minN, maxN int) float64 {
	value := float64(f.IntBetween(minN, maxN-1))
	if maxDecimals < 1 {
		return value
	}

	p := int(math.Pow10(maxDecimals))
	decimals := float64(f.IntBetween(0, p)) / float64(p)

	return value + decimals
}

// RandomFloat returns a fake random float number for Faker
func (f Faker) RandomFloat(maxDecimals, minN, maxN int) float64 {
	return f.generateFloat(maxDecimals, minN, maxN)
}

// Float returns a fake random float number for Faker
func (f Faker) Float(maxDecimals, minN, maxN int) float64 {
	return f.generateFloat(maxDecimals, minN, maxN)
}

// Float32 returns a fake random float32 number for Faker
func (f Faker) Float32(maxDecimals, minN, maxN int) float32 {
	return float32(f.generateFloat(maxDecimals, minN, maxN))
}

// Float64 returns a fake random float64 number for Faker
func (f Faker) Float64(maxDecimals, minN, maxN int) float64 {
	return f.generateFloat(maxDecimals, minN, maxN)
}

// Int returns a fake Int number for Faker
func (f Faker) Int() int {
	return f.IntBetween(math.MinInt, math.MaxInt)
}

// Int8 returns a fake Int8 number for Faker
func (f Faker) Int8() int8 {
	return f.Int8Between(math.MinInt8, math.MaxInt8)
}

// Int16 returns a fake Int16 number for Faker
func (f Faker) Int16() int16 {
	return f.Int16Between(math.MinInt16, math.MaxInt16)
}

// Int32 returns a fake Int32 number for Faker
func (f Faker) Int32() int32 {
	return f.Int32Between(math.MinInt32, math.MaxInt32)
}

// Int64 returns a fake Int64 number for Faker
func (f Faker) Int64() int64 {
	return f.Int64Between(math.MinInt64, math.MaxInt64)
}

// UInt returns a fake UInt number for Faker
func (f Faker) UInt() uint {
	return f.UIntBetween(0, math.MaxUint)
}

// UInt8 returns a fake UInt8 number for Faker
func (f Faker) UInt8() uint8 {
	return f.UInt8Between(0, math.MaxUint8)
}

// UInt16 returns a fake UInt16 number for Faker
func (f Faker) UInt16() uint16 {
	return f.UInt16Between(0, math.MaxUint16)
}

// UInt32 returns a fake UInt32 number for Faker
func (f Faker) UInt32() uint32 {
	return f.UInt32Between(0, math.MaxUint32)
}

// UInt64 returns a fake UInt64 number for Faker
func (f Faker) UInt64() uint64 {
	return f.UInt64Between(0, math.MaxUint64)
}

type number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// minInt returns the minimum value for a given number
func minInt[T number](num T) T {
	var ret any
	switch any(num).(type) {
	case int:
		ret = int(math.MinInt)
	case int8:
		ret = int8(math.MinInt8)
	case int16:
		ret = int16(math.MinInt16)
	case int32:
		ret = int32(math.MinInt32)
	case int64:
		ret = int64(math.MinInt64)
	case uint, uint8, uint16, uint32, uint64:
		ret = T(0)
	}

	return ret.(T)
}

// maxInt returns the maximum value for a given number
func maxInt[T number](num T) T {
	var ret any
	switch any(num).(type) {
	case int:
		ret = int(math.MaxInt)
	case int8:
		ret = int8(math.MaxInt8)
	case int16:
		ret = int16(math.MaxInt16)
	case int32:
		ret = int32(math.MaxInt32)
	case int64:
		ret = int64(math.MaxInt64)
	case uint:
		ret = uint(math.MaxUint)
	case uint8:
		ret = uint8(math.MaxUint8)
	case uint16:
		ret = uint16(math.MaxUint16)
	case uint32:
		ret = uint32(math.MaxUint32)
	case uint64:
		ret = uint64(math.MaxUint64)
	}

	return ret.(T)
}

// between returns a fake number between a given minimum and maximum value using generator
func between[T number](minN, maxN T, rand GeneratorInterface) T {
	// Ensure minN <= maxN
	if minN > maxN {
		minN, maxN = maxN, minN
	}

	// Handle edge case: full range of integer type
	if minN == minInt(minN) && maxN == maxInt(maxN) {
		return handleFullRange(minN, maxN, rand)
	}

	// Handle equal values
	diff := maxN - minN
	if diff == 0 {
		return minN
	}

	// Generate random value based on type
	var value T
	if diff == maxInt(maxN) {
		// Special case: diff equals max value
		value = generateMaxRangeValue[T](maxN, rand)
	} else {
		// Normal case: generate value in range [0, diff]
		value = generateRangeValue(diff, rand)
	}

	return minN + value
}

// handleFullRange handles the special case when range covers entire type
func handleFullRange[T number](minN, maxN T, rand GeneratorInterface) T {
	// Split range: 50% negative, 50% positive to avoid overflow
	if rand.Intn(2) == 0 {
		// Generate negative number in range [minN, 0]
		maxVal := maxInt(maxN)
		value := generateMaxRangeValue[T](maxN, rand)
		return minN + value
	}
	// Generate positive number in range [0, maxInt]
	return generateMaxRangeValue[T](maxN, rand)
}

// generateMaxRangeValue generates a value when diff equals MaxInt
func generateMaxRangeValue[T number](maxN T, rand GeneratorInterface) T {
	maxVal := maxInt(maxN) - 1
	switch any(maxN).(type) {
	case int:
		return T(rand.Intn(int(maxVal)))
	case int8, int16, int32:
		return T(rand.Int32n(int32(maxVal)))
	case int64:
		return T(rand.Int64n(int64(maxVal)))
	case uint:
		return T(rand.Uintn(uint(maxVal)))
	case uint8, uint16, uint32:
		return T(rand.Uint32n(uint32(maxVal)))
	case uint64:
		return T(rand.Uint64n(uint64(maxVal)))
	default:
		return 0
	}
}

// generateRangeValue generates a value in range [0, diff]
func generateRangeValue[T number](diff T, rand GeneratorInterface) T {
	// Add 1 to diff to make range inclusive
	switch any(diff).(type) {
	case int:
		return T(rand.Intn(int(diff + 1)))
	case int8, int16, int32:
		return T(rand.Int32n(int32(diff + 1)))
	case int64:
		return T(rand.Int64n(int64(diff + 1)))
	case uint:
		return T(rand.Uintn(uint(diff + 1)))
	case uint8, uint16, uint32:
		return T(rand.Uint32n(uint32(diff + 1)))
	case uint64:
		return T(rand.Uint64n(uint64(diff + 1)))
	default:
		return 0
	}
}

// IntBetween returns a fake Int between a given minimum and maximum values for Faker
func (f Faker) IntBetween(minN, maxN int) int {
	return between(minN, maxN, f.Generator)
}

// Int8Between returns a fake Int8 between a given minimum and maximum values for Faker
func (f Faker) Int8Between(minN, maxN int8) int8 {
	return between(minN, maxN, f.Generator)
}

// Int16Between returns a fake Int16 between a given minimum and maximum values for Faker
func (f Faker) Int16Between(minN, maxN int16) int16 {
	return between(minN, maxN, f.Generator)
}

// Int32Between returns a fake Int32 between a given minimum and maximum values for Faker
func (f Faker) Int32Between(minN, maxN int32) int32 {
	return between(minN, maxN, f.Generator)
}

// Int64Between returns a fake Int64 between a given minimum and maximum values for Faker
func (f Faker) Int64Between(minN, maxN int64) int64 {
	return between(minN, maxN, f.Generator)
}

// UIntBetween returns a fake UInt between a given minimum and maximum values for Faker
func (f Faker) UIntBetween(minN, maxN uint) uint {
	return between(minN, maxN, f.Generator)
}

// UInt8Between returns a fake UInt8 between a given minimum and maximum values for Faker
func (f Faker) UInt8Between(minN, maxN uint8) uint8 {
	return between(minN, maxN, f.Generator)
}

// UInt16Between returns a fake UInt16 between a given minimum and maximum values for Faker
func (f Faker) UInt16Between(minN, maxN uint16) uint16 {
	return between(minN, maxN, f.Generator)
}

// UInt32Between returns a fake UInt32 between a given minimum and maximum values for Faker
func (f Faker) UInt32Between(minN, maxN uint32) uint32 {
	return between(minN, maxN, f.Generator)
}

// UInt64Between returns a fake UInt64 between a given minimum and maximum values for Faker
func (f Faker) UInt64Between(minN, maxN uint64) uint64 {
	return between(minN, maxN, f.Generator)
}

// Letter returns a fake single letter for Faker
func (f Faker) Letter() string {
	return f.RandomLetter()
}

// RandomLetter returns a fake random string with a random number of letters for Faker
func (f Faker) RandomLetter() string {
	return fmt.Sprintf("%c", f.IntBetween(97, 122))
}

func (f Faker) RandomStringWithLength(l int) string {
	r := make([]string, 0, l)

	for i := 0; i < l; i++ {
		r = append(r, f.RandomLetter())
	}
	return strings.Join(r, "")
}

// RandomStringElement returns a fake random string element from a given list of strings for Faker
func (f Faker) RandomStringElement(s []string) string {
	i := f.IntBetween(0, len(s)-1)
	return s[i]
}

// RandomStringMapKey returns a fake random string key from a given map[string]string for Faker
func (f Faker) RandomStringMapKey(m map[string]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	i := f.IntBetween(0, len(keys)-1)
	return keys[i]
}

// RandomStringMapValue returns a fake random string value from a given map[string]string for Faker
func (f Faker) RandomStringMapValue(m map[string]string) string {
	values := make([]string, 0, len(m))
	for k := range m {
		values = append(values, m[k])
	}
	sort.Strings(values)

	i := f.IntBetween(0, len(values)-1)
	return values[i]
}

// RandomIntElement returns a fake random int element form a given list of ints for Faker
func (f Faker) RandomIntElement(a []int) int {
	i := f.IntBetween(0, len(a)-1)
	return a[i]
}

// ShuffleString returns a fake shuffled string from a given string for Faker
func (Faker) ShuffleString(s string) string {
	orig := strings.Split(s, "")
	return strings.Join(Shuffle(orig), "")
}

// Numerify returns a fake string that replace all "#" characters with numbers from a given string for Faker
func (f Faker) Numerify(in string) (out string) {
	if !strings.Contains(in, "#") {
		return in
	}

	var builder strings.Builder
	builder.Grow(len(in))

	for _, c := range in {
		if c == '#' {
			builder.WriteString(strconv.Itoa(f.RandomDigit()))
		} else {
			builder.WriteRune(c)
		}
	}

	return builder.String()
}

// Lexify  returns a fake string that replace all "?" characters with random letters from a given string for Faker
func (f Faker) Lexify(in string) (out string) {
	if !strings.Contains(in, "?") {
		return in
	}

	var builder strings.Builder
	builder.Grow(len(in))

	for _, c := range in {
		if c == '?' {
			builder.WriteString(f.RandomLetter())
		} else {
			builder.WriteRune(c)
		}
	}

	return builder.String()
}

// Bothify returns a fake string that apply Lexify() and Numerify() on a given string for Faker
func (f Faker) Bothify(in string) (out string) {
	out = f.Lexify(in)
	out = f.Numerify(out)
	return
}

// Asciify   returns a fake string that replace all "*" characters with random ASCII values from a given string for Faker
func (f Faker) Asciify(in string) (out string) {
	if !strings.Contains(in, "*") {
		return in
	}

	var builder strings.Builder
	builder.Grow(len(in))

	for _, c := range in {
		if c == '*' {
			builder.WriteByte(byte(f.IntBetween(97, 126)))
		} else {
			builder.WriteRune(c)
		}
	}

	return builder.String()
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
	m := make(map[string]interface{})

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
		sl := make([]string, 0, 10)
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
			mm := make(map[string]interface{})

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
	return Image{&f, TempFileCreatorImpl{}, PngEncoderImpl{}}
}

// File returns a fake File instance for Faker
func (f Faker) File() File {
	return File{&f, OSResolverImpl{}}
}

// Directory returns a fake Directory instance for Faker
func (f Faker) Directory() Directory {
	return Directory{&f, OSResolverImpl{}}
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

// App returns a fake App instance for Faker
func (f Faker) App() App {
	return App{&f}
}

// Pet returns a fake Pet instance for Faker
func (f Faker) Pet() Pet {
	return Pet{&f}
}

// Emoji returns a fake Emoji instance for Faker
func (f Faker) Emoji() Emoji {
	return Emoji{&f}
}

// LoremFlickr returns a fake LoremFlickr instance for Faker
func (f Faker) LoremFlickr() LoremFlickr {
	return LoremFlickr{&f, HTTPClientImpl{}, TempFileCreatorImpl{}}
}

// ProfileImage returns a fake ProfileImage instance for Faker
func (f Faker) ProfileImage() ProfileImage {
	return ProfileImage{&f, HTTPClientImpl{}, TempFileCreatorImpl{}}
}

// Genre returns a fake Genre instance for Faker
func (f Faker) Genre() Genre {
	return Genre{&f}
}

// Gender returns a fake Gender instance for Faker
func (f Faker) Gender() Gender {
	return Gender{&f}
}

// BinaryString returns a fake BinaryString instance for Faker
func (f Faker) BinaryString() BinaryString {
	return BinaryString{&f}
}

// Hash returns a fake Hash instance for Faker
func (f Faker) Hash() Hash {
	return Hash{&f}
}

// Music returns a fake Music instance for Faker
func (f Faker) Music() Music {
	return Music{&f}
}

// Currency returns a fake Currency instance for Faker
func (f Faker) Currency() Currency {
	return Currency{&f}
}

// Crypto returns a fake Crypto instance for Faker
func (f Faker) Crypto() Crypto {
	return Crypto{&f}
}

// ProgrammingLanguage returns a fake ProgrammingLanguage instance for Faker
func (f Faker) ProgrammingLanguage() ProgrammingLanguage {
	return ProgrammingLanguage{&f}
}

// New returns a new instance of Faker with a random seed
func New() Faker {
	return NewWithSeed(rand.NewPCG(rand.Uint64(), rand.Uint64()))
}

// NewWithSeed returns a new instance of Faker with a given seed
func NewWithSeed(src rand.Source) Faker {
	generator := &threadSafeRand{
		rand: rand.New(src),
	}
	return Faker{Generator: generator}
}

// NewWithSeedInt64 returns a new instance of Faker seeded with the given value
func NewWithSeedInt64(seed int64) Faker {
	generator := &threadSafeRand{
		rand: rand.New(rand.NewPCG(uint64(seed), uint64(seed))),
	}
	return Faker{Generator: generator}
}

// Blood returns a fake Blood instance for Faker
func (f Faker) Blood() Blood {
	return Blood{&f}
}

// Json returns a fake Json instance for Faker
func (f Faker) Json() Json {
	return Json{&f}
}

// Pokemon returns a fake Pokemon instance for Faker
func (f Faker) Pokemon() Pokemon {
	return Pokemon{Faker: &f}
}
