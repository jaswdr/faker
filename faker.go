// Package faker is a comprehensive Go library for generating fake data for testing,
// database seeding, and data anonymization. It provides thread-safe random data
// generation across various data types including names, addresses, internet data,
// payments, and more.
//
// Basic Usage:
//
//	f := faker.New()
//	name := f.Person().Name()        // "John Doe"
//	email := f.Internet().Email()    // "john.doe@example.com"
//	phone := f.Phone().Number()      // "+1-555-123-4567"
//
// Seeded Generation (for reproducible results):
//
//	f := faker.NewWithSeedInt64(12345)
//	name := f.Person().Name() // Always returns the same name for seed 12345
//
// Struct Filling (using struct tags):
//
//	type User struct {
//		Name  string `fake:"{{person.first_name}} {{person.last_name}}"`
//		Email string `fake:"{{internet.email}}"`
//		Age   int    `fake:"{{number.number_int_between 18 65}}"`
//	}
//
//	var user User
//	f.Struct().Fill(&user)
//
// Performance Characteristics:
//
// - Thread-safe: All Faker instances can be used concurrently
// - Memory efficient: Uses sync.Pool for frequently allocated objects
// - Fast generation: Pre-compiled patterns and cached data structures
// - Supports Go 1.22+ with math/rand/v2 for improved performance
//
// Error Handling:
//
// Most methods in this library are designed to never fail and will return
// reasonable defaults. Methods that can fail (like file operations) return
// errors explicitly. For struct filling, invalid tags are ignored silently.
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

// Constants for commonly used values
const (
	// defaultSliceMinSize is the minimum size for generated slices/arrays/maps
	defaultSliceMinSize = 1
	// defaultSliceMaxSize is the maximum size for generated slices/arrays/maps
	defaultSliceMaxSize = 10

	// ASCII character ranges
	lowerCaseA = 97  // 'a'
	lowerCaseZ = 122 // 'z'
	asciiStart = 97  // start of ASCII printable range for Asciify
	asciiEnd   = 126 // end of ASCII printable range for Asciify

	// defaultStringLength is the default length for random strings
	defaultStringLength = 10

	// maxRetriesDefault is the default number of retries for operations
	maxRetriesDefault = 7
)

// Pool for reusing strings.Builder instances to reduce allocations
var stringBuilderPool = sync.Pool{
	New: func() interface{} {
		return &strings.Builder{}
	},
}

// getStringBuilder gets a strings.Builder from the pool
func getStringBuilder() *strings.Builder {
	return stringBuilderPool.Get().(*strings.Builder)
}

// putStringBuilder returns a strings.Builder to the pool after resetting it
func putStringBuilder(sb *strings.Builder) {
	sb.Reset()
	stringBuilderPool.Put(sb)
}

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

// RandomDigitNot returns a fake random digit (0-9) that is not in the list of ignored values.
// If all digits are ignored, this function will loop indefinitely, so use with caution.
//
// Example:
//
//	digit := f.RandomDigitNot(0, 5, 9) // Returns a digit that's not 0, 5, or 9
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

// RandomNumber returns a fake random integer number for Faker.
// The size parameter must be positive. If size is invalid, returns a single digit.
func (f Faker) RandomNumber(size int) int {
	if size <= 0 {
		return f.RandomDigit()
	}
	if size == 1 {
		return f.RandomDigit()
	}

	minN := int(math.Pow10(size - 1))
	maxN := int(math.Pow10(size)) - 1

	return f.IntBetween(minN, maxN)
}

// generateFloat is a helper function that generates a random float with the specified parameters.
// Validates inputs to ensure reasonable values.
func (f Faker) generateFloat(maxDecimals, minN, maxN int) float64 {
	// Ensure minN <= maxN
	if minN > maxN {
		minN, maxN = maxN, minN
	}

	// Ensure valid decimal places
	if maxDecimals < 0 {
		maxDecimals = 0
	}
	if maxDecimals > 10 { // Reasonable limit
		maxDecimals = 10
	}

	// Generate a value between minN and maxN-1 to leave room for decimals
	value := float64(f.IntBetween(minN, maxN))
	if maxDecimals < 1 {
		return value
	}

	p := int(math.Pow10(maxDecimals))
	decimals := float64(f.IntBetween(0, p)) / float64(p)

	// If we're at the maximum value, ensure decimals don't push us over
	if int(value) == maxN && decimals > 0 {
		// Scale down decimals to keep within bounds
		return value
	}

	result := value + decimals
	// Ensure we don't exceed maxN
	if result > float64(maxN) {
		return float64(maxN)
	}

	return result
}

// RandomFloat returns a fake random float number for Faker.
// maxDecimals: number of decimal places (capped at 10 for performance)
// minN, maxN: range boundaries (inclusive)
//
// Example:
//
//	price := f.RandomFloat(2, 10, 1000) // Returns a float like 123.45 between 10.00 and 1000.00
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
	return fmt.Sprintf("%c", f.IntBetween(lowerCaseA, lowerCaseZ))
}

// RandomStringWithLength returns a fake random string with the specified length.
// If length is negative or zero, returns an empty string.
// If length is excessively large (>1000), caps it at 1000 for performance.
func (f Faker) RandomStringWithLength(l int) string {
	if l <= 0 {
		return ""
	}

	// Cap at reasonable limit for performance
	if l > 1000 {
		l = 1000
	}

	builder := getStringBuilder()
	defer putStringBuilder(builder)

	builder.Grow(l)

	for i := 0; i < l; i++ {
		builder.WriteString(f.RandomLetter())
	}
	return builder.String()
}

// RandomStringElement returns a fake random string element from a given list of strings for Faker.
// Returns empty string if slice is nil or empty.
func (f Faker) RandomStringElement(s []string) string {
	if len(s) == 0 {
		return ""
	}
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

// RandomIntElement returns a fake random int element from a given list of ints for Faker.
// Returns 0 if slice is nil or empty.
func (f Faker) RandomIntElement(a []int) int {
	if len(a) == 0 {
		return 0
	}
	i := f.IntBetween(0, len(a)-1)
	return a[i]
}

// ShuffleString returns a fake shuffled string from a given string for Faker
func (Faker) ShuffleString(s string) string {
	orig := strings.Split(s, "")
	return strings.Join(Shuffle(orig), "")
}

// Numerify returns a fake string that replaces all "#" characters with random digits (0-9).
// Uses sync.Pool for efficient string building to minimize allocations.
//
// Example:
//
//	orderID := f.Numerify("ORD-####-###") // Returns something like "ORD-1234-567"
func (f Faker) Numerify(in string) (out string) {
	if !strings.Contains(in, "#") {
		return in
	}

	builder := getStringBuilder()
	defer putStringBuilder(builder)

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

// Lexify returns a fake string that replaces all "?" characters with random lowercase letters (a-z).
// Uses sync.Pool for efficient string building to minimize allocations.
//
// Example:
//
//	code := f.Lexify("???-???") // Returns something like "abc-xyz"
func (f Faker) Lexify(in string) (out string) {
	if !strings.Contains(in, "?") {
		return in
	}

	builder := getStringBuilder()
	defer putStringBuilder(builder)

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

// Bothify returns a fake string that applies both Lexify() and Numerify() transformations.
// First replaces "?" with letters, then "#" with numbers.
//
// Example:
//
//	serial := f.Bothify("??##-??##") // Returns something like "ab12-cd34"
func (f Faker) Bothify(in string) (out string) {
	out = f.Lexify(in)
	out = f.Numerify(out)
	return
}

// Asciify returns a fake string that replaces all "*" characters with random ASCII printable characters.
// Uses characters from ASCII 97-126 (lowercase letters and symbols).
// Uses sync.Pool for efficient string building to minimize allocations.
//
// Example:
//
//	password := f.Asciify("****-****") // Returns something like "a{7~-b}2@"
func (f Faker) Asciify(in string) (out string) {
	if !strings.Contains(in, "*") {
		return in
	}

	builder := getStringBuilder()
	defer putStringBuilder(builder)

	builder.Grow(len(in))

	for _, c := range in {
		if c == '*' {
			builder.WriteByte(byte(f.IntBetween(asciiStart, asciiEnd)))
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
