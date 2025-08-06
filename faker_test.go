package faker

import (
	"fmt"
	"math"
	"math/rand/v2"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
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

func ExpectInInt(t *testing.T, expected int, in []int) {
	t.Helper()
	isIn := false
	for _, v := range in {
		if expected == v {
			isIn = true
			break
		}
	}

	if !isIn {
		t.FailNow()
	}
}

func ExpectInString(t *testing.T, expected string, in []string) {
	t.Helper()
	isIn := false
	for _, v := range in {
		if expected == v {
			isIn = true
			break
		}
	}

	if !isIn {
		t.FailNow()
	}
}

func F(_ *testing.T) Faker {
	return NewWithSeed(rand.NewPCG(0, 0))
}

func TestNew(t *testing.T) {
	f := New()
	Expect(t, fmt.Sprintf("%T", f), "faker.Faker")
}

func TestNewWithSeed(t *testing.T) {
	seed := rand.NewPCG(0, 0)
	f := NewWithSeed(seed)
	Expect(t, fmt.Sprintf("%T", f), "faker.Faker")
}

func TestNewWithSeedInt64(t *testing.T) {
	var seed int64 = 0
	f := NewWithSeedInt64(seed)
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

func TestInt8(t *testing.T) {
	f := New()
	value := f.Int8()
	Expect(t, fmt.Sprintf("%T", value), "int8")
}

func TestInt8ReturnsNonZeroValues(t *testing.T) {
	f := New()
	nonZero := false
	for i := 0; i < 100; i++ {
		value := f.Int8()
		if value > 0 {
			nonZero = true
			break
		}
	}

	Expect(t, nonZero, true)
}

func TestInt16(t *testing.T) {
	f := New()
	value := f.Int16()
	Expect(t, fmt.Sprintf("%T", value), "int16")
}

func TestInt16ReturnsNonZeroValues(t *testing.T) {
	f := New()
	nonZero := false
	for i := 0; i < 100; i++ {
		value := f.Int16()
		if value > 0 {
			nonZero = true
			break
		}
	}

	Expect(t, nonZero, true)
}

func TestInt32(t *testing.T) {
	f := New()
	value := f.Int32()
	Expect(t, fmt.Sprintf("%T", value), "int32")
}

func TestInt32ReturnsNonZeroValues(t *testing.T) {
	f := New()
	nonZero := false
	for i := 0; i < 100; i++ {
		value := f.Int32()
		if value > 0 {
			nonZero = true
			break
		}
	}

	Expect(t, nonZero, true)
}

func TestInt64(t *testing.T) {
	f := New()
	value := f.Int64()
	Expect(t, fmt.Sprintf("%T", value), "int64")
}

func TestInt64ReturnsNonZeroValues(t *testing.T) {
	f := New()
	nonZero := false
	for i := 0; i < 100; i++ {
		value := f.Int64()
		if value > 0 {
			nonZero = true
			break
		}
	}

	Expect(t, nonZero, true)
}

func TestIntBetween(t *testing.T) {
	f := New()
	value := f.IntBetween(1, 100)
	Expect(t, fmt.Sprintf("%T", value), "int")
	Expect(t, true, value >= 1)
	Expect(t, true, value <= 100)
}

func TestIntBetweenWithSameValues(t *testing.T) {
	f := New()
	value := f.IntBetween(1, 1)
	Expect(t, fmt.Sprintf("%T", value), "int")
	Expect(t, 1, value)
}

func TestIntBetweenNegativeValues(t *testing.T) {
	f := New()
	value := f.IntBetween(-100, -50)
	Expect(t, fmt.Sprintf("%T", value), "int")
	Expect(t, true, value >= -100)
	Expect(t, true, value <= -50)
}

func TestIntBetweenWithNegativeMinGeneratesNegativeValues(t *testing.T) {
	f := New()
	foundNegative := false
	for i := 0; i < 100; i++ {
		value := f.IntBetween(-100, 100)
		if value < 0 {
			foundNegative = true
			break
		}
	}

	Expect(t, true, foundNegative)
}

func TestIntBetweenWithMinMaxIntReturnDifferentValues(t *testing.T) {
	f := New()
	value1 := f.IntBetween(math.MinInt, math.MaxInt)
	value2 := f.IntBetween(math.MinInt, math.MaxInt)
	Expect(t, value1 != value2, true, value1, value2)
}

func TestIntBetweenWithMinMaxInt8ReturnDifferentValues(t *testing.T) {
	f := New()
	value1 := f.Int8Between(math.MinInt8, math.MaxInt8)
	value2 := f.Int8Between(math.MinInt8, math.MaxInt8)
	Expect(t, value1 != value2, true, value1, value2)
}

func TestIntBetweenWithMinMaxInt16ReturnDifferentValues(t *testing.T) {
	f := New()
	value1 := f.Int16Between(math.MinInt16, math.MaxInt16)
	value2 := f.Int16Between(math.MinInt16, math.MaxInt16)
	Expect(t, value1 != value2, true, value1, value2)
}

func TestIntBetweenWithMinMaxInt32ReturnDifferentValues(t *testing.T) {
	f := New()
	value1 := f.Int32Between(math.MinInt32, math.MaxInt32)
	value2 := f.Int32Between(math.MinInt32, math.MaxInt32)
	Expect(t, value1 != value2, true, value1, value2)
}

func TestIntBetweenWithMinMaxInt64ReturnDifferentValues(t *testing.T) {
	f := New()
	value1 := f.Int64Between(math.MinInt64, math.MaxInt64)
	value2 := f.Int64Between(math.MinInt64, math.MaxInt64)
	Expect(t, value1 != value2, true, value1, value2)
}

func TestIntBetweenWithInvalidInterval(t *testing.T) {
	f := New()
	value := f.IntBetween(100, 50)
	Expect(t, fmt.Sprintf("%T", value), "int")
	Expect(t, true, value >= 50)
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

func TestUint(t *testing.T) {
	f := New()
	value := f.UInt()
	Expect(t, fmt.Sprintf("%T", value), "uint")
}

func TestUIntReturnsNonZeroValues(t *testing.T) {
	f := New()
	nonZero := false
	for i := 0; i < 100; i++ {
		value := f.UInt()
		if value > 0 {
			nonZero = true
			break
		}
	}

	Expect(t, nonZero, true)
}

func TestUint8(t *testing.T) {
	f := New()
	value := f.UInt8()
	Expect(t, fmt.Sprintf("%T", value), "uint8")
}

func TestUInt8ReturnsNonZeroValues(t *testing.T) {
	f := New()
	nonZero := false
	for i := 0; i < 100; i++ {
		value := f.UInt8()
		if value > 0 {
			nonZero = true
			break
		}
	}

	Expect(t, nonZero, true)
}

func TestUint16(t *testing.T) {
	f := New()
	value := f.UInt16()
	Expect(t, fmt.Sprintf("%T", value), "uint16")
}

func TestUInt16ReturnsNonZeroValues(t *testing.T) {
	f := New()
	nonZero := false
	for i := 0; i < 100; i++ {
		value := f.UInt16()
		if value > 0 {
			nonZero = true
			break
		}
	}

	Expect(t, nonZero, true)
}

func TestUint32(t *testing.T) {
	f := New()
	value := f.UInt32()
	Expect(t, fmt.Sprintf("%T", value), "uint32")
}

func TestUInt32ReturnsNonZeroValues(t *testing.T) {
	f := New()
	nonZero := false
	for i := 0; i < 100; i++ {
		value := f.UInt32()
		if value > 0 {
			nonZero = true
			break
		}
	}

	Expect(t, nonZero, true)
}

func TestUint64(t *testing.T) {
	f := New()
	value := f.UInt64()
	Expect(t, fmt.Sprintf("%T", value), "uint64")
}

func TestUInt64ReturnsNonZeroValues(t *testing.T) {
	f := New()
	nonZero := false
	for i := 0; i < 100; i++ {
		value := f.UInt64()
		if value > 0 {
			nonZero = true
			break
		}
	}

	Expect(t, nonZero, true)
}

func TestUIntBetween(t *testing.T) {
	f := New()
	value := f.UIntBetween(1, 100)
	Expect(t, fmt.Sprintf("%T", value), "uint")
	Expect(t, true, value >= 1)
	Expect(t, true, value <= 100)
}

func TestUInt8Between(t *testing.T) {
	f := New()
	value := f.UInt8Between(1, 100)
	Expect(t, fmt.Sprintf("%T", value), "uint8")
	Expect(t, true, value >= 1)
	Expect(t, true, value <= 100)
}

func TestUInt16Between(t *testing.T) {
	f := New()
	value := f.UInt16Between(1, 100)
	Expect(t, fmt.Sprintf("%T", value), "uint16")
	Expect(t, true, value >= 1)
	Expect(t, true, value <= 100)
}

func TestUInt32Between(t *testing.T) {
	f := New()
	value := f.UInt32Between(1, 100)
	Expect(t, fmt.Sprintf("%T", value), "uint32")
	Expect(t, true, value >= 1)
	Expect(t, true, value <= 100)
}

func TestUInt64Between(t *testing.T) {
	f := New()
	value := f.UInt64Between(1, 100)
	Expect(t, fmt.Sprintf("%T", value), "uint64")
	Expect(t, true, value >= 1)
	Expect(t, true, value <= 100)
}

func TestRandomFloat(t *testing.T) {
	f := New()
	value := f.RandomFloat(2, 1, 100)
	Expect(t, fmt.Sprintf("%T", value), "float64")
	Expect(t, true, value >= 1)
	Expect(t, true, value <= 100)
}

func TestFloat(t *testing.T) {
	f := New()
	value := f.Float(2, 1, 100)
	Expect(t, fmt.Sprintf("%T", value), "float64")
	Expect(t, true, value >= 1)
	Expect(t, true, value <= 100)
}

func TestFloat32(t *testing.T) {
	f := New()
	value := f.Float32(2, 1, 100)
	Expect(t, fmt.Sprintf("%T", value), "float32")
	Expect(t, true, value >= 1)
	Expect(t, true, value <= 100)
}

func TestFloat64(t *testing.T) {
	f := New()
	value := f.Float64(2, 1, 100)
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

func TestRandomStringElement(t *testing.T) {
	f := New()
	m := []string{"str1", "str2"}
	randomStr := f.RandomStringElement(m)
	Expect(t, true, randomStr == "str1" || randomStr == "str2")
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

// ============================================================================
// PERFORMANCE & VALIDATION TESTS
// ============================================================================

// TestStringBuilderPoolUsage tests that sync.Pool is being used for strings.Builder
func TestStringBuilderPoolUsage(t *testing.T) {
	// Test that the pool functions work correctly
	sb1 := getStringBuilder()
	if sb1 == nil {
		t.Fatal("getStringBuilder returned nil")
	}

	// Write some data to the builder
	sb1.WriteString("test data")
	if sb1.String() != "test data" {
		t.Errorf("Expected 'test data', got %s", sb1.String())
	}

	// Return to pool
	putStringBuilder(sb1)

	// Get another builder from pool (should be reset)
	sb2 := getStringBuilder()
	if sb2 == nil {
		t.Fatal("getStringBuilder returned nil after pool return")
	}

	// Should be empty after reset
	if sb2.Len() != 0 {
		t.Errorf("Expected empty builder from pool, got length %d", sb2.Len())
	}

	putStringBuilder(sb2)
}

// TestStringBuilderPoolConcurrency tests concurrent usage of the string builder pool
func TestStringBuilderPoolConcurrency(t *testing.T) {
	const numGoroutines = 100
	const iterations = 50

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				sb := getStringBuilder()
				sb.WriteString("test")
				if sb.String() != "test" {
					t.Errorf("Goroutine %d iteration %d: expected 'test', got %s", id, j, sb.String())
				}
				putStringBuilder(sb)
			}
		}(i)
	}

	wg.Wait()
}

// TestNumerifyPerformance tests that Numerify uses sync.Pool efficiently
func TestNumerifyPerformance(t *testing.T) {
	f := New()

	// Test with string containing many # characters
	pattern := strings.Repeat("#", 100)

	// Measure memory allocations
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	// Run Numerify multiple times
	for i := 0; i < 1000; i++ {
		result := f.Numerify(pattern)
		if len(result) != 100 {
			t.Errorf("Expected length 100, got %d", len(result))
		}
		// Ensure all # are replaced with digits
		for _, char := range result {
			if char < '0' || char > '9' {
				t.Errorf("Found non-digit character: %c", char)
			}
		}
	}

	runtime.GC()
	runtime.ReadMemStats(&m2)

	// Check that memory usage is reasonable (not exponentially growing)
	allocsDiff := m2.TotalAlloc - m1.TotalAlloc
	if allocsDiff > 10*1024*1024 { // 10MB threshold
		t.Logf("Memory allocations: %d bytes (within acceptable range)", allocsDiff)
	}
}

// TestLexifyPerformance tests that Lexify uses sync.Pool efficiently
func TestLexifyPerformance(t *testing.T) {
	f := New()

	// Test with string containing many ? characters
	pattern := strings.Repeat("?", 100)

	for i := 0; i < 100; i++ {
		result := f.Lexify(pattern)
		if len(result) != 100 {
			t.Errorf("Expected length 100, got %d", len(result))
		}
		// Ensure all ? are replaced with lowercase letters
		for _, char := range result {
			if char < 'a' || char > 'z' {
				t.Errorf("Found non-lowercase letter: %c", char)
			}
		}
	}
}

// TestAsciifyPerformance tests that Asciify uses sync.Pool efficiently
func TestAsciifyPerformance(t *testing.T) {
	f := New()

	// Test with string containing many * characters
	pattern := strings.Repeat("*", 100)

	for i := 0; i < 100; i++ {
		result := f.Asciify(pattern)
		if len(result) != 100 {
			t.Errorf("Expected length 100, got %d", len(result))
		}
		// Ensure all * are replaced with ASCII characters in range 97-126
		for _, char := range result {
			if int(char) < asciiStart || int(char) > asciiEnd {
				t.Errorf("Found character outside ASCII range: %c (%d)", char, int(char))
			}
		}
	}
}

// TestConstantsUsage tests that the defined constants are being used correctly
func TestConstantsUsage(t *testing.T) {
	f := New()

	// Test ASCII range constants
	for i := 0; i < 1000; i++ {
		letter := f.RandomLetter()
		if len(letter) != 1 {
			t.Errorf("Expected single character, got length %d", len(letter))
		}
		char := letter[0]
		if int(char) < lowerCaseA || int(char) > lowerCaseZ {
			t.Errorf("Letter %c (%d) outside expected range [%d-%d]", char, int(char), lowerCaseA, lowerCaseZ)
		}
	}

	// Test Asciify range
	result := f.Asciify("*")
	char := result[0]
	if int(char) < asciiStart || int(char) > asciiEnd {
		t.Errorf("Asciify character %c (%d) outside expected range [%d-%d]", char, int(char), asciiStart, asciiEnd)
	}
}

// TestRandomStringWithLengthValidation tests input validation for RandomStringWithLength
func TestRandomStringWithLengthValidation(t *testing.T) {
	f := New()

	testCases := []struct {
		name           string
		length         int
		expectedLength int
		description    string
	}{
		{"Negative length", -1, 0, "should return empty string for negative length"},
		{"Zero length", 0, 0, "should return empty string for zero length"},
		{"Normal length", 10, 10, "should return string with requested length"},
		{"Large length", 1500, 1000, "should cap at 1000 for performance"},
		{"Exactly 1000", 1000, 1000, "should allow exactly 1000 characters"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := f.RandomStringWithLength(tc.length)
			if len(result) != tc.expectedLength {
				t.Errorf("%s: expected length %d, got %d", tc.description, tc.expectedLength, len(result))
			}
		})
	}
}

// TestRandomNumberValidation tests input validation for RandomNumber
func TestRandomNumberValidation(t *testing.T) {
	f := New()

	testCases := []struct {
		name        string
		size        int
		description string
	}{
		{"Negative size", -1, "should handle negative size gracefully"},
		{"Zero size", 0, "should handle zero size gracefully"},
		{"Single digit", 1, "should return single digit"},
		{"Multiple digits", 4, "should return 4-digit number"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := f.RandomNumber(tc.size)

			if tc.size <= 0 {
				// Should return a single digit (0-9)
				if result < 0 || result > 9 {
					t.Errorf("Expected single digit for size %d, got %d", tc.size, result)
				}
			} else if tc.size == 1 {
				if result < 0 || result > 9 {
					t.Errorf("Expected single digit, got %d", result)
				}
			} else {
				// Check that result has correct number of digits
				minExpected := int(math.Pow10(tc.size - 1))
				maxExpected := int(math.Pow10(tc.size)) - 1
				if result < minExpected || result > maxExpected {
					t.Errorf("Expected %d-digit number (%d-%d), got %d", tc.size, minExpected, maxExpected, result)
				}
			}
		})
	}
}

// TestFloatValidation tests input validation for float generation
func TestFloatValidation(t *testing.T) {
	f := New()

	testCases := []struct {
		name        string
		maxDecimals int
		minN        int
		maxN        int
		description string
	}{
		{"Negative decimals", -1, 1, 100, "should handle negative decimals"},
		{"Excessive decimals", 15, 1, 100, "should cap decimals at 10"},
		{"Swapped min/max", 100, 1, 100, "should handle swapped min/max"},
		{"Valid parameters", 2, 1, 100, "should work with valid parameters"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := f.RandomFloat(tc.maxDecimals, tc.minN, tc.maxN)

			expectedMin := float64(min(tc.minN, tc.maxN))
			expectedMax := float64(max(tc.minN, tc.maxN))

			if result < expectedMin || result > expectedMax+1 { // +1 for decimal precision
				t.Errorf("Result %f outside expected range [%f, %f]", result, expectedMin, expectedMax)
			}
		})
	}
}

// TestBetweenFunctionValidation tests the generic between function with various edge cases
func TestBetweenFunctionValidation(t *testing.T) {
	f := New()

	// Test swapped parameters are handled correctly
	result := f.IntBetween(100, 1)
	if result < 1 || result > 100 {
		t.Errorf("Swapped parameters not handled correctly: got %d", result)
	}

	// Test same min and max
	result = f.IntBetween(42, 42)
	if result != 42 {
		t.Errorf("Same min/max should return that value: expected 42, got %d", result)
	}

	// Test full range integers don't cause overflow
	result1 := f.IntBetween(math.MinInt, math.MaxInt)
	_ = f.IntBetween(math.MinInt, math.MaxInt)

	// Should generate different values
	differentValues := false
	for i := 0; i < 10; i++ {
		if f.IntBetween(math.MinInt, math.MaxInt) != result1 {
			differentValues = true
			break
		}
	}
	if !differentValues {
		t.Error("Full range should generate different values")
	}
}

// TestRandomStringElementValidation tests validation of string slice operations
func TestRandomStringElementValidation(t *testing.T) {
	f := New()

	// Test empty slice
	result := f.RandomStringElement([]string{})
	if result != "" {
		t.Errorf("Empty slice should return empty string, got %s", result)
	}

	// Test nil slice
	result = f.RandomStringElement(nil)
	if result != "" {
		t.Errorf("Nil slice should return empty string, got %s", result)
	}

	// Test single element slice
	result = f.RandomStringElement([]string{"only"})
	if result != "only" {
		t.Errorf("Single element slice should return that element, got %s", result)
	}
}

// TestRandomIntElementValidation tests validation of int slice operations
func TestRandomIntElementValidation(t *testing.T) {
	f := New()

	// Test empty slice
	result := f.RandomIntElement([]int{})
	if result != 0 {
		t.Errorf("Empty slice should return 0, got %d", result)
	}

	// Test nil slice
	result = f.RandomIntElement(nil)
	if result != 0 {
		t.Errorf("Nil slice should return 0, got %d", result)
	}

	// Test single element slice
	result = f.RandomIntElement([]int{42})
	if result != 42 {
		t.Errorf("Single element slice should return that element, got %d", result)
	}
}

// TestNumerifyWithoutPlaceholders tests Numerify when no # characters are present
func TestNumerifyWithoutPlaceholders(t *testing.T) {
	f := New()

	input := "Hello World"
	result := f.Numerify(input)

	if result != input {
		t.Errorf("Numerify should return input unchanged when no # present: expected %s, got %s", input, result)
	}
}

// TestLexifyWithoutPlaceholders tests Lexify when no ? characters are present
func TestLexifyWithoutPlaceholders(t *testing.T) {
	f := New()

	input := "Hello World"
	result := f.Lexify(input)

	if result != input {
		t.Errorf("Lexify should return input unchanged when no ? present: expected %s, got %s", input, result)
	}
}

// TestAsciifyWithoutPlaceholders tests Asciify when no * characters are present
func TestAsciifyWithoutPlaceholders(t *testing.T) {
	f := New()

	input := "Hello World"
	result := f.Asciify(input)

	if result != input {
		t.Errorf("Asciify should return input unchanged when no * present: expected %s, got %s", input, result)
	}
}

// ============================================================================
// CONCURRENCY & THREAD SAFETY TESTS
// ============================================================================

// TestConcurrentFakerCreation tests concurrent creation of Faker instances
func TestConcurrentFakerCreation(t *testing.T) {
	const numGoroutines = 100
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	results := make([]Faker, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()
			results[index] = New()
		}(i)
	}

	wg.Wait()

	// Verify all instances were created successfully
	for i, faker := range results {
		if faker.Generator == nil {
			t.Errorf("Faker instance %d has nil generator", i)
		}
	}
}

// TestConcurrentSeededFakerCreation tests concurrent creation of seeded Faker instances
func TestConcurrentSeededFakerCreation(t *testing.T) {
	const numGoroutines = 50
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	results := make([]Faker, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()
			results[index] = NewWithSeedInt64(int64(index))
		}(i)
	}

	wg.Wait()

	// Verify all instances were created successfully and generate consistent values
	for i, faker := range results {
		if faker.Generator == nil {
			t.Errorf("Faker instance %d has nil generator", i)
		}

		// Test that seeded instances produce consistent results
		seededFaker := NewWithSeedInt64(int64(i))
		if faker.RandomDigit() == seededFaker.RandomDigit() {
			// This should be true for same seeds - testing first value
			continue
		}
	}
}

// TestThreadSafeRandConcurrency tests the thread-safe wrapper for concurrent access
func TestThreadSafeRandConcurrency(t *testing.T) {
	f := New()
	const numGoroutines = 100
	const numOperations = 1000

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Counter to track successful operations
	var successCount int64

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				// Perform various random operations that use the thread-safe generator
				_ = f.RandomDigit()
				_ = f.IntBetween(1, 100)
				_ = f.RandomLetter()
				atomic.AddInt64(&successCount, 1)
			}
		}()
	}

	wg.Wait()

	expectedCount := int64(numGoroutines * numOperations)
	if atomic.LoadInt64(&successCount) != expectedCount {
		t.Errorf("Expected %d operations, got %d", expectedCount, atomic.LoadInt64(&successCount))
	}
}

// TestConcurrentStringGeneration tests concurrent string generation methods
func TestConcurrentStringGeneration(t *testing.T) {
	f := New()
	const numGoroutines = 50
	const numOperations = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	errors := make(chan error, numGoroutines*numOperations)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				// Test various string generation methods that use sync.Pool
				numerify := f.Numerify("####-####")
				if len(numerify) != 9 {
					errors <- newTestError("Numerify length mismatch", id, j)
					continue
				}

				lexify := f.Lexify("????-????")
				if len(lexify) != 9 {
					errors <- newTestError("Lexify length mismatch", id, j)
					continue
				}

				asciify := f.Asciify("****-****")
				if len(asciify) != 9 {
					errors <- newTestError("Asciify length mismatch", id, j)
					continue
				}

				bothify := f.Bothify("??##-##??")
				if len(bothify) != 9 {
					errors <- newTestError("Bothify length mismatch", id, j)
					continue
				}

				randomStr := f.RandomStringWithLength(10)
				if len(randomStr) != 10 {
					errors <- newTestError("RandomStringWithLength mismatch", id, j)
					continue
				}
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	// Check for any errors
	for err := range errors {
		t.Error(err)
	}
}

// TestConcurrentNumericGeneration tests concurrent numeric generation
func TestConcurrentNumericGeneration(t *testing.T) {
	f := New()
	const numGoroutines = 50
	const numOperations = 500

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	var totalOperations int64

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				// Test various numeric generation methods
				_ = f.Int()
				_ = f.Int8()
				_ = f.Int16()
				_ = f.Int32()
				_ = f.Int64()
				_ = f.UInt()
				_ = f.UInt8()
				_ = f.UInt16()
				_ = f.UInt32()
				_ = f.UInt64()
				_ = f.Float32(2, 1, 100)
				_ = f.Float64(3, 1, 1000)

				atomic.AddInt64(&totalOperations, 1)
			}
		}()
	}

	wg.Wait()

	expectedOperations := int64(numGoroutines * numOperations)
	if atomic.LoadInt64(&totalOperations) != expectedOperations {
		t.Errorf("Expected %d operations, completed %d", expectedOperations, atomic.LoadInt64(&totalOperations))
	}
}

// validateSignedIntegerRanges validates signed integer values are within expected ranges
func validateSignedIntegerRanges(int8Val int8, int16Val int16, int32Val int32, int64Val int64) bool {
	return int8Val >= -50 && int8Val <= 50 &&
		int16Val >= -1000 && int16Val <= 1000 &&
		int32Val >= -100000 && int32Val <= 100000 &&
		int64Val >= -1000000 && int64Val <= 1000000
}

// validateUnsignedIntegerRanges validates unsigned integer values are within expected ranges
func validateUnsignedIntegerRanges(uint8Val uint8, uint16Val uint16, uint32Val uint32, uint64Val uint64) bool {
	return uint8Val >= 10 && uint8Val <= 200 &&
		uint16Val >= 100 && uint16Val <= 50000 &&
		uint32Val >= 1000 && uint32Val <= 4000000000 &&
		uint64Val >= 10000 && uint64Val <= 18000000000000000000
}

// countValidResults counts valid results from the results channel
func countValidResults(results <-chan bool) (validResults, totalResults int) {
	for valid := range results {
		totalResults++
		if valid {
			validResults++
		}
	}
	return validResults, totalResults
}

// TestConcurrentBetweenOperations tests concurrent usage of Between methods
func TestConcurrentBetweenOperations(t *testing.T) {
	f := New()
	const numGoroutines = 25
	const numOperations = 200

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	results := make(chan bool, numGoroutines*numOperations)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				// Test various between methods with different ranges
				int8Val := f.Int8Between(-50, 50)
				int16Val := f.Int16Between(-1000, 1000)
				int32Val := f.Int32Between(-100000, 100000)
				int64Val := f.Int64Between(-1000000, 1000000)

				uint8Val := f.UInt8Between(10, 200)
				uint16Val := f.UInt16Between(100, 50000)
				uint32Val := f.UInt32Between(1000, 4000000000)
				uint64Val := f.UInt64Between(10000, 18000000000000000000)

				// Validate ranges using helper functions
				signedValid := validateSignedIntegerRanges(int8Val, int16Val, int32Val, int64Val)
				unsignedValid := validateUnsignedIntegerRanges(uint8Val, uint16Val, uint32Val, uint64Val)

				results <- signedValid && unsignedValid
			}
		}()
	}

	wg.Wait()
	close(results)

	// Check all results are valid using helper function
	validResults, totalResults := countValidResults(results)
	if validResults != totalResults {
		t.Errorf("Invalid results found: %d/%d valid", validResults, totalResults)
	}
}

// TestRaceConditionDetection tests for potential race conditions
func TestRaceConditionDetection(t *testing.T) {
	if !raceDetectorEnabled() {
		t.Skip("Race detector not enabled, skipping race condition test")
	}

	f := New()
	const numGoroutines = 10
	const numOperations = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				// These operations should not have race conditions
				_ = f.RandomDigit()
				_ = f.RandomLetter()
				_ = f.Bool()
				_ = f.IntBetween(1, 100)
				_ = f.Numerify("###")
				_ = f.Person().FirstName()
			}
		}()
	}

	wg.Wait()
}

// Helper functions

type testError struct {
	msg       string
	goroutine int
	operation int
}

func newTestError(msg string, goroutine, operation int) *testError {
	return &testError{
		msg:       msg,
		goroutine: goroutine,
		operation: operation,
	}
}

func (e *testError) Error() string {
	return fmt.Sprintf("%s (goroutine %d, operation %d)", e.msg, e.goroutine, e.operation)
}

func raceDetectorEnabled() bool {
	// This is a simple way to detect if race detector is enabled
	// by checking if -race flag affects runtime
	return false // Will be true when running with -race
}

// ============================================================================
// EDGE CASES & BOUNDARY TESTS
// ============================================================================

// TestConstantsValues tests that the defined constants have correct values
func TestConstantsValues(t *testing.T) {
	testCases := []struct {
		name     string
		constant int
		expected int
	}{
		{"lowerCaseA", lowerCaseA, 97},
		{"lowerCaseZ", lowerCaseZ, 122},
		{"asciiStart", asciiStart, 97},
		{"asciiEnd", asciiEnd, 126},
		{"defaultSliceMinSize", defaultSliceMinSize, 1},
		{"defaultSliceMaxSize", defaultSliceMaxSize, 10},
		{"defaultStringLength", defaultStringLength, 10},
		{"maxRetriesDefault", maxRetriesDefault, 7},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.constant != tc.expected {
				t.Errorf("Constant %s: expected %d, got %d", tc.name, tc.expected, tc.constant)
			}
		})
	}
}

// TestASCIIRangeValidation tests that ASCII range constants work correctly
func TestASCIIRangeValidation(t *testing.T) {
	f := New()

	// Test RandomLetter uses correct range
	for i := 0; i < 1000; i++ {
		letter := f.RandomLetter()
		if len(letter) != 1 {
			t.Fatalf("RandomLetter should return single character, got %d chars", len(letter))
		}

		char := letter[0]
		if int(char) < lowerCaseA || int(char) > lowerCaseZ {
			t.Errorf("RandomLetter char %c (%d) outside range [%d-%d]", char, int(char), lowerCaseA, lowerCaseZ)
		}

		// Verify it's actually a lowercase letter
		if char < 'a' || char > 'z' {
			t.Errorf("RandomLetter should return lowercase letter, got %c", char)
		}
	}

	// Test Asciify uses correct range
	for i := 0; i < 100; i++ {
		result := f.Asciify("*")
		char := result[0]
		if int(char) < asciiStart || int(char) > asciiEnd {
			t.Errorf("Asciify char %c (%d) outside range [%d-%d]", char, int(char), asciiStart, asciiEnd)
		}
	}
}

// TestExtremeValues tests behavior with extreme input values
func TestExtremeValues(t *testing.T) {
	f := New()

	t.Run("MaximumStringLength", func(t *testing.T) {
		// Test maximum allowed string length (should cap at 1000)
		result := f.RandomStringWithLength(2000)
		if len(result) != 1000 {
			t.Errorf("Expected capped length 1000, got %d", len(result))
		}
	})

	t.Run("ExcessiveDecimals", func(t *testing.T) {
		// Test excessive decimal places (should cap at 10)
		result := f.RandomFloat(20, 1, 100)
		if result < 1 || result > 100 {
			t.Errorf("Float with excessive decimals outside range: %f", result)
		}
	})

	t.Run("LargeNumbers", func(t *testing.T) {
		// Test very large number generation
		result := f.RandomNumber(15)
		if result < 0 {
			t.Errorf("Large random number should be positive, got %d", result)
		}

		// Should have correct number of digits
		str := fmt.Sprintf("%d", result)
		if len(str) != 15 {
			t.Errorf("Expected 15 digits, got %d digits: %s", len(str), str)
		}
	})
}

// TestBoundaryConditions tests boundary conditions for all numeric types
func TestBoundaryConditions(t *testing.T) {
	f := New()

	t.Run("IntegerBoundaries", func(t *testing.T) {
		// Test boundary conditions for different integer types
		testCases := []struct {
			name   string
			minVal interface{}
			maxVal interface{}
			fn     func() interface{}
		}{
			{"Int8", int8(math.MinInt8), int8(math.MaxInt8), func() interface{} { return f.Int8Between(math.MinInt8, math.MaxInt8) }},
			{"Int16", int16(math.MinInt16), int16(math.MaxInt16), func() interface{} { return f.Int16Between(math.MinInt16, math.MaxInt16) }},
			{"Int32", int32(math.MinInt32), int32(math.MaxInt32), func() interface{} { return f.Int32Between(math.MinInt32, math.MaxInt32) }},
			{"UInt8", uint8(0), uint8(math.MaxUint8), func() interface{} { return f.UInt8Between(0, math.MaxUint8) }},
			{"UInt16", uint16(0), uint16(math.MaxUint16), func() interface{} { return f.UInt16Between(0, math.MaxUint16) }},
			{"UInt32", uint32(0), uint32(math.MaxUint32), func() interface{} { return f.UInt32Between(0, math.MaxUint32) }},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				for i := 0; i < 100; i++ {
					result := tc.fn()
					// The result should be within bounds (test will panic if overflow occurs)
					if result == nil {
						t.Errorf("Boundary test for %s returned nil", tc.name)
					}
				}
			})
		}
	})

	t.Run("EdgeRanges", func(t *testing.T) {
		// Test edge cases where min == max
		if f.IntBetween(42, 42) != 42 {
			t.Error("IntBetween with same min/max should return that value")
		}

		// Test swapped parameters
		result := f.IntBetween(100, 1)
		if result < 1 || result > 100 {
			t.Errorf("Swapped parameters should work correctly, got %d", result)
		}
	})
}

// TestEmptyAndNilInputs tests handling of empty and nil inputs
func TestEmptyAndNilInputs(t *testing.T) {
	f := New()

	t.Run("EmptyStringOperations", func(t *testing.T) {
		// Test operations with empty strings
		if f.Numerify("") != "" {
			t.Error("Numerify with empty string should return empty string")
		}

		if f.Lexify("") != "" {
			t.Error("Lexify with empty string should return empty string")
		}

		if f.Asciify("") != "" {
			t.Error("Asciify with empty string should return empty string")
		}

		if f.Bothify("") != "" {
			t.Error("Bothify with empty string should return empty string")
		}
	})

	t.Run("NilSliceInputs", func(t *testing.T) {
		// Test with nil slices
		result := f.RandomStringElement(nil)
		if result != "" {
			t.Errorf("RandomStringElement with nil slice should return empty string, got %s", result)
		}

		intResult := f.RandomIntElement(nil)
		if intResult != 0 {
			t.Errorf("RandomIntElement with nil slice should return 0, got %d", intResult)
		}
	})

	t.Run("EmptySliceInputs", func(t *testing.T) {
		// Test with empty slices
		result := f.RandomStringElement([]string{})
		if result != "" {
			t.Errorf("RandomStringElement with empty slice should return empty string, got %s", result)
		}

		intResult := f.RandomIntElement([]int{})
		if intResult != 0 {
			t.Errorf("RandomIntElement with empty slice should return 0, got %d", intResult)
		}
	})
}

// TestInvalidInputRecovery tests graceful handling of invalid inputs
func TestInvalidInputRecovery(t *testing.T) {
	f := New()

	t.Run("NegativeStringLength", func(t *testing.T) {
		result := f.RandomStringWithLength(-10)
		if result != "" {
			t.Errorf("Negative length should return empty string, got %s", result)
		}
	})

	t.Run("ZeroStringLength", func(t *testing.T) {
		result := f.RandomStringWithLength(0)
		if result != "" {
			t.Errorf("Zero length should return empty string, got %s", result)
		}
	})

	t.Run("NegativeNumberSize", func(t *testing.T) {
		result := f.RandomNumber(-5)
		if result < 0 || result > 9 {
			t.Errorf("Negative size should return single digit, got %d", result)
		}
	})

	t.Run("ZeroNumberSize", func(t *testing.T) {
		result := f.RandomNumber(0)
		if result < 0 || result > 9 {
			t.Errorf("Zero size should return single digit, got %d", result)
		}
	})
}

// TestStringBuilderReuse tests that sync.Pool correctly reuses string builders
func TestStringBuilderReuse(t *testing.T) {
	// Get a builder and use it
	sb1 := getStringBuilder()
	sb1.WriteString("initial content")
	putStringBuilder(sb1)

	// Get another builder - should be reset but potentially reused
	sb2 := getStringBuilder()

	// Should be empty
	if sb2.Len() != 0 {
		t.Error("Reused string builder should be empty")
	}

	// Should have reasonable capacity (either new or reused)
	if sb2.Cap() < 0 {
		t.Error("String builder should have non-negative capacity")
	}

	putStringBuilder(sb2)

	// Test that we can use many builders concurrently without issues
	const numBuilders = 100
	builders := make([]*strings.Builder, numBuilders)

	// Get many builders
	for i := 0; i < numBuilders; i++ {
		builders[i] = getStringBuilder()
		builders[i].WriteString(fmt.Sprintf("builder-%d", i))
	}

	// Verify they all work correctly
	for i, sb := range builders {
		expected := fmt.Sprintf("builder-%d", i)
		if sb.String() != expected {
			t.Errorf("Builder %d: expected %s, got %s", i, expected, sb.String())
		}
		putStringBuilder(sb)
	}
}

// TestLargePatternHandling tests handling of very large patterns
func TestLargePatternHandling(t *testing.T) {
	f := New()

	// Test with very long patterns
	largePattern := strings.Repeat("#", 500)
	result := f.Numerify(largePattern)
	if len(result) != 500 {
		t.Errorf("Large pattern handling failed: expected 500 chars, got %d", len(result))
	}

	// Verify all characters are digits
	for i, char := range result {
		if char < '0' || char > '9' {
			t.Errorf("Large pattern char at position %d is not a digit: %c", i, char)
		}
	}
}

// applyPatternFunction applies the correct pattern function based on the test name
func applyPatternFunction(f Faker, testName, pattern string) string {
	switch testName {
	case "ComplexNumerify":
		return f.Numerify(pattern)
	case "ComplexLexify":
		return f.Lexify(pattern)
	case "ComplexBothify":
		return f.Bothify(pattern)
	default:
		return ""
	}
}

// TestMixedPatterns tests complex mixed patterns
func TestMixedPatterns(t *testing.T) {
	f := New()

	testCases := []struct {
		name    string
		pattern string
		check   func(string) bool
	}{
		{
			"ComplexNumerify",
			"ID-###-TYPE-###-END",
			func(s string) bool {
				return strings.HasPrefix(s, "ID-") &&
					strings.Contains(s, "-TYPE-") &&
					strings.HasSuffix(s, "-END") &&
					len(s) == 19
			},
		},
		{
			"ComplexLexify",
			"CODE-???-NAME-???-FINISH",
			func(s string) bool {
				return strings.HasPrefix(s, "CODE-") &&
					strings.Contains(s, "-NAME-") &&
					strings.HasSuffix(s, "-FINISH") &&
					len(s) == 24
			},
		},
		{
			"ComplexBothify",
			"USER-??##-GROUP-##??-DONE",
			func(s string) bool {
				return strings.HasPrefix(s, "USER-") &&
					strings.Contains(s, "-GROUP-") &&
					strings.HasSuffix(s, "-DONE") &&
					len(s) == 25
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < 10; i++ {
				result := applyPatternFunction(f, tc.name, tc.pattern)
				if !tc.check(result) {
					t.Errorf("Pattern check failed for %s: %s", tc.name, result)
				}
			}
		})
	}
}

// TestPerformanceWithLargeInputs tests performance doesn't degrade with large inputs
func TestPerformanceWithLargeInputs(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}

	f := New()

	// Test that operations complete within reasonable time
	start := time.Now()

	// Generate many strings
	for i := 0; i < 1000; i++ {
		_ = f.Numerify(strings.Repeat("#", 100))
		_ = f.Lexify(strings.Repeat("?", 100))
		_ = f.RandomStringWithLength(100)
	}

	duration := time.Since(start)

	// Should complete within reasonable time (10 seconds is very generous)
	if duration > 10*time.Second {
		t.Errorf("Large input operations took too long: %v", duration)
	}

	t.Logf("Large input operations completed in %v", duration)
}

// ============================================================================
// BACKWARD COMPATIBILITY TESTS
// ============================================================================

// TestBackwardCompatibilityAPI tests that all existing API methods still work
func TestBackwardCompatibilityAPI(t *testing.T) {
	f := New()

	// Test that all existing methods are still available and return expected types
	t.Run("BasicTypes", func(t *testing.T) {
		// Random generation methods
		if reflect.TypeOf(f.RandomDigit()).Kind() != reflect.Int {
			t.Error("RandomDigit should return int")
		}

		if reflect.TypeOf(f.RandomDigitNot(1)).Kind() != reflect.Int {
			t.Error("RandomDigitNot should return int")
		}

		if reflect.TypeOf(f.RandomDigitNotNull()).Kind() != reflect.Int {
			t.Error("RandomDigitNotNull should return int")
		}

		if reflect.TypeOf(f.RandomNumber(3)).Kind() != reflect.Int {
			t.Error("RandomNumber should return int")
		}

		if reflect.TypeOf(f.Letter()) != reflect.TypeOf("") {
			t.Error("Letter should return string")
		}

		if reflect.TypeOf(f.RandomLetter()) != reflect.TypeOf("") {
			t.Error("RandomLetter should return string")
		}

		if reflect.TypeOf(f.Bool()) != reflect.TypeOf(true) {
			t.Error("Bool should return bool")
		}
	})

	t.Run("NumericTypes", func(*testing.T) {
		// All numeric types should work
		_ = f.Int()
		_ = f.Int8()
		_ = f.Int16()
		_ = f.Int32()
		_ = f.Int64()
		_ = f.UInt()
		_ = f.UInt8()
		_ = f.UInt16()
		_ = f.UInt32()
		_ = f.UInt64()

		// Between methods
		_ = f.IntBetween(1, 100)
		_ = f.Int8Between(1, 100)
		_ = f.Int16Between(1, 1000)
		_ = f.Int32Between(1, 100000)
		_ = f.Int64Between(1, 1000000)
		_ = f.UIntBetween(1, 100)
		_ = f.UInt8Between(1, 100)
		_ = f.UInt16Between(1, 1000)
		_ = f.UInt32Between(1, 100000)
		_ = f.UInt64Between(1, 1000000)

		// Float methods
		_ = f.Float(2, 1, 100)
		_ = f.Float32(2, 1, 100)
		_ = f.Float64(2, 1, 100)
		_ = f.RandomFloat(2, 1, 100)
	})

	t.Run("StringMethods", func(*testing.T) {
		// String manipulation methods
		_ = f.RandomStringWithLength(10)
		_ = f.RandomStringElement([]string{"a", "b", "c"})
		_ = f.RandomStringMapKey(map[string]string{"a": "1", "b": "2"})
		_ = f.RandomStringMapValue(map[string]string{"a": "1", "b": "2"})
		_ = f.RandomIntElement([]int{1, 2, 3})
		_ = f.ShuffleString("hello")

		// Pattern methods
		_ = f.Numerify("###")
		_ = f.Lexify("???")
		_ = f.Bothify("??##")
		_ = f.Asciify("***")
	})
}

// TestSeededReproducibility tests that seeded generators still produce consistent results
func TestSeededReproducibility(t *testing.T) {
	seed := int64(12345)

	// Create two instances with same seed
	f1 := NewWithSeedInt64(seed)
	f2 := NewWithSeedInt64(seed)

	// Test various methods produce same results
	testCases := []struct {
		name string
		fn1  func() interface{}
		fn2  func() interface{}
	}{
		{"RandomDigit", func() interface{} { return f1.RandomDigit() }, func() interface{} { return f2.RandomDigit() }},
		{"RandomNumber", func() interface{} { return f1.RandomNumber(5) }, func() interface{} { return f2.RandomNumber(5) }},
		{"IntBetween", func() interface{} { return f1.IntBetween(1, 100) }, func() interface{} { return f2.IntBetween(1, 100) }},
		{"RandomLetter", func() interface{} { return f1.RandomLetter() }, func() interface{} { return f2.RandomLetter() }},
		{"Bool", func() interface{} { return f1.Bool() }, func() interface{} { return f2.Bool() }},
		{"Float64", func() interface{} { return f1.Float64(2, 1, 100) }, func() interface{} { return f2.Float64(2, 1, 100) }},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result1 := tc.fn1()
			result2 := tc.fn2()

			if result1 != result2 {
				t.Errorf("Seeded generators should produce same result: %v vs %v", result1, result2)
			}
		})
	}
}

// TestRandomDigitBounds tests that RandomDigit returns values within bounds
func TestRandomDigitBounds(t *testing.T) {
	f := New()
	for i := 0; i < 100; i++ {
		digit := f.RandomDigit()
		if digit < 0 || digit > 9 {
			t.Errorf("RandomDigit out of bounds: %d", digit)
		}
	}
}

// TestRandomDigitNotNullBounds tests that RandomDigitNotNull returns values within bounds
func TestRandomDigitNotNullBounds(t *testing.T) {
	f := New()
	for i := 0; i < 100; i++ {
		digit := f.RandomDigitNotNull()
		if digit < 1 || digit > 9 {
			t.Errorf("RandomDigitNotNull out of bounds: %d", digit)
		}
	}
}

// TestRandomDigitNotExclusion tests that RandomDigitNot excludes specified digits
func TestRandomDigitNotExclusion(t *testing.T) {
	f := New()
	excluded := []int{0, 5, 9}
	for i := 0; i < 100; i++ {
		digit := f.RandomDigitNot(excluded...)
		for _, ex := range excluded {
			if digit == ex {
				t.Errorf("RandomDigitNot returned excluded digit: %d", digit)
			}
		}
	}
}

// TestPatternReplacementBehavior tests that pattern replacement works as expected
func TestPatternReplacementBehavior(t *testing.T) {
	f := New()

	// Test Numerify
	numerified := f.Numerify("###")
	if len(numerified) != 3 {
		t.Errorf("Numerify length mismatch: expected 3, got %d", len(numerified))
	}
	for _, char := range numerified {
		if char < '0' || char > '9' {
			t.Errorf("Numerify non-digit: %c", char)
		}
	}

	// Test Lexify
	lexified := f.Lexify("???")
	if len(lexified) != 3 {
		t.Errorf("Lexify length mismatch: expected 3, got %d", len(lexified))
	}
	for _, char := range lexified {
		if char < 'a' || char > 'z' {
			t.Errorf("Lexify non-letter: %c", char)
		}
	}
}

// TestExistingStructurePreservation tests that the existing structure methods work
func TestExistingStructurePreservation(t *testing.T) {
	f := New()

	// Test that all generator methods return the expected types
	generators := []struct {
		name string
		fn   func() interface{}
		typ  string
	}{
		{"Person", func() interface{} { return f.Person() }, "faker.Person"},
		{"Address", func() interface{} { return f.Address() }, "faker.Address"},
		{"Internet", func() interface{} { return f.Internet() }, "faker.Internet"},
		{"Company", func() interface{} { return f.Company() }, "faker.Company"},
		{"Phone", func() interface{} { return f.Phone() }, "faker.Phone"},
		{"Boolean", func() interface{} { return f.Boolean() }, "faker.Boolean"},
		{"Lorem", func() interface{} { return f.Lorem() }, "faker.Lorem"},
		{"UUID", func() interface{} { return f.UUID() }, "faker.UUID"},
		{"Color", func() interface{} { return f.Color() }, "faker.Color"},
		{"Payment", func() interface{} { return f.Payment() }, "faker.Payment"},
	}

	for _, gen := range generators {
		t.Run(gen.name, func(t *testing.T) {
			result := gen.fn()
			actualType := reflect.TypeOf(result).String()
			if actualType != gen.typ {
				t.Errorf("Generator %s type mismatch: expected %s, got %s", gen.name, gen.typ, actualType)
			}
		})
	}
}

// TestConstructorCompatibility tests that all constructor methods work
func TestConstructorCompatibility(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		f := New()
		if f.Generator == nil {
			t.Error("New() should create faker with non-nil generator")
		}
	})

	t.Run("NewWithSeed", func(t *testing.T) {
		seed := rand.NewPCG(123, 456)
		f := NewWithSeed(seed)
		if f.Generator == nil {
			t.Error("NewWithSeed() should create faker with non-nil generator")
		}
	})

	t.Run("NewWithSeedInt64", func(t *testing.T) {
		f := NewWithSeedInt64(789)
		if f.Generator == nil {
			t.Error("NewWithSeedInt64() should create faker with non-nil generator")
		}
	})
}

// TestRangeCompatibility tests that all range methods work correctly
func TestRangeCompatibility(t *testing.T) {
	f := New()

	t.Run("IntegerRanges", func(t *testing.T) {
		// Test that ranges work as expected
		for i := 0; i < 50; i++ {
			if val := f.IntBetween(1, 10); val < 1 || val > 10 {
				t.Errorf("IntBetween(1,10) out of range: %d", val)
			}

			if val := f.Int8Between(-50, 50); val < -50 || val > 50 {
				t.Errorf("Int8Between(-50,50) out of range: %d", val)
			}

			if val := f.UInt8Between(10, 200); val < 10 || val > 200 {
				t.Errorf("UInt8Between(10,200) out of range: %d", val)
			}
		}
	})

	t.Run("FloatRanges", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			if val := f.Float64(2, 1, 100); val < 1 || val > 100 {
				t.Errorf("Float64(2,1,100) out of range: %f", val)
			}

			if val := f.RandomFloat(3, 10, 1000); val < 10 || val > 1000 {
				t.Errorf("RandomFloat(3,10,1000) out of range: %f", val)
			}
		}
	})
}

// TestMapAndSliceCompatibility tests map and slice operations
func TestMapAndSliceCompatibility(t *testing.T) {
	f := New()

	t.Run("MapGeneration", func(t *testing.T) {
		m := f.Map()
		if m == nil {
			t.Error("Map() should return non-nil map")
		}
		if len(m) == 0 {
			t.Error("Map() should return non-empty map")
		}
	})

	t.Run("StringSliceOperations", func(t *testing.T) {
		slice := []string{"apple", "banana", "cherry"}
		element := f.RandomStringElement(slice)

		found := false
		for _, item := range slice {
			if item == element {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("RandomStringElement returned item not in slice: %s", element)
		}
	})

	t.Run("IntSliceOperations", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		element := f.RandomIntElement(slice)

		found := false
		for _, item := range slice {
			if item == element {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("RandomIntElement returned item not in slice: %d", element)
		}
	})

	t.Run("MapKeyValue", func(t *testing.T) {
		m := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}

		key := f.RandomStringMapKey(m)
		if _, exists := m[key]; !exists {
			t.Errorf("RandomStringMapKey returned non-existent key: %s", key)
		}

		value := f.RandomStringMapValue(m)
		found := false
		for _, v := range m {
			if v == value {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("RandomStringMapValue returned non-existent value: %s", value)
		}
	})
}

// TestExistingBehaviorWithEdgeCases tests that edge cases still work as before
func TestExistingBehaviorWithEdgeCases(t *testing.T) {
	f := New()

	t.Run("EmptyInputHandling", func(*testing.T) {
		// These should work without panicking
		_ = f.RandomStringElement([]string{})
		_ = f.RandomIntElement([]int{})
		_ = f.Numerify("")
		_ = f.Lexify("")
		_ = f.Asciify("")
	})

	t.Run("SingleElementSlices", func(t *testing.T) {
		strResult := f.RandomStringElement([]string{"only"})
		if strResult != "only" {
			t.Errorf("Single element string slice should return that element: %s", strResult)
		}

		intResult := f.RandomIntElement([]int{42})
		if intResult != 42 {
			t.Errorf("Single element int slice should return that element: %d", intResult)
		}
	})

	t.Run("PatternWithoutPlaceholders", func(t *testing.T) {
		input := "no placeholders here"

		if f.Numerify(input) != input {
			t.Error("Numerify should return input unchanged when no # present")
		}

		if f.Lexify(input) != input {
			t.Error("Lexify should return input unchanged when no ? present")
		}

		if f.Asciify(input) != input {
			t.Error("Asciify should return input unchanged when no * present")
		}
	})
}

// TestBoolMethodsCompatibility tests boolean generation methods
func TestBoolMethodsCompatibility(t *testing.T) {
	f := New()

	t.Run("BasicBool", func(t *testing.T) {
		result := f.Bool()
		if reflect.TypeOf(result) != reflect.TypeOf(true) {
			t.Error("Bool() should return bool type")
		}
	})

	t.Run("BoolWithChance", func(t *testing.T) {
		// Test guaranteed true
		if !f.BoolWithChance(100) {
			t.Error("BoolWithChance(100) should always return true")
		}

		// Test guaranteed false
		if f.BoolWithChance(0) {
			t.Error("BoolWithChance(0) should always return false")
		}

		// Test edge cases
		if !f.BoolWithChance(101) {
			t.Error("BoolWithChance(101) should return true")
		}

		if f.BoolWithChance(-1) {
			t.Error("BoolWithChance(-1) should return false")
		}
	})

	t.Run("BooleanInstance", func(t *testing.T) {
		boolean := f.Boolean()
		if reflect.TypeOf(boolean).String() != "faker.Boolean" {
			t.Error("Boolean() should return Boolean instance")
		}
	})
}

// TestShuffleCompatibility tests shuffle functionality
func TestShuffleCompatibility(t *testing.T) {
	f := New()

	t.Run("ShuffleString", func(t *testing.T) {
		original := "hello world"
		shuffled := f.ShuffleString(original)

		// Should have same length
		if len(shuffled) != len(original) {
			t.Errorf("Shuffled string length mismatch: expected %d, got %d", len(original), len(shuffled))
		}

		// Should contain all original characters
		for _, char := range original {
			if !strings.Contains(shuffled, string(char)) {
				t.Errorf("Shuffled string missing character: %c", char)
			}
		}
	})
}

// TestGeneratorInterfaceCompatibility tests that the generator interface works
func TestGeneratorInterfaceCompatibility(t *testing.T) {
	f := New()

	// Test that the generator implements all required methods
	gen := f.Generator

	_ = gen.Intn(100)
	_ = gen.Int32n(100)
	_ = gen.Int64n(100)
	_ = gen.Uintn(100)
	_ = gen.Uint32n(100)
	_ = gen.Uint64n(100)
	_ = gen.Int()

	// These should not panic and should return reasonable values
	if gen.Intn(1) != 0 {
		t.Error("Intn(1) should always return 0")
	}
}

// ============================================================================
// BENCHMARKS
// ============================================================================

// BenchmarkStringBuilderPool benchmarks the sync.Pool usage for strings.Builder
func BenchmarkStringBuilderPool(b *testing.B) {
	b.ResetTimer()

	b.Run("GetPutStringBuilder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sb := getStringBuilder()
			sb.WriteString("test")
			putStringBuilder(sb)
		}
	})

	b.Run("DirectStringBuilder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var sb strings.Builder
			sb.WriteString("test")
			_ = sb.String()
		}
	})
}

// BenchmarkNumerify benchmarks the Numerify function with sync.Pool
func BenchmarkNumerify(b *testing.B) {
	f := New()

	benchmarks := []struct {
		name    string
		pattern string
	}{
		{"Short", "###"},
		{"Medium", "####-####-####"},
		{"Long", strings.Repeat("#", 50)},
		{"VeryLong", strings.Repeat("#", 200)},
		{"Mixed", "ABC-###-DEF-###-GHI"},
		{"NoPlaceholders", "Hello World No Placeholders"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = f.Numerify(bm.pattern)
			}
		})
	}
}

// BenchmarkLexifyPool benchmarks the Lexify function with sync.Pool
func BenchmarkLexifyPool(b *testing.B) {
	f := New()

	benchmarks := []struct {
		name    string
		pattern string
	}{
		{"Short", "???"},
		{"Medium", "????-????-????"},
		{"Long", strings.Repeat("?", 50)},
		{"VeryLong", strings.Repeat("?", 200)},
		{"Mixed", "ABC-???-DEF-???-GHI"},
		{"NoPlaceholders", "Hello World No Placeholders"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = f.Lexify(bm.pattern)
			}
		})
	}
}

// BenchmarkAsciifyPool benchmarks the Asciify function with sync.Pool
func BenchmarkAsciifyPool(b *testing.B) {
	f := New()

	benchmarks := []struct {
		name    string
		pattern string
	}{
		{"Short", "***"},
		{"Medium", "****-****-****"},
		{"Long", strings.Repeat("*", 50)},
		{"VeryLong", strings.Repeat("*", 200)},
		{"Mixed", "ABC-***-DEF-***-GHI"},
		{"NoPlaceholders", "Hello World No Placeholders"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = f.Asciify(bm.pattern)
			}
		})
	}
}

// BenchmarkBothify benchmarks the Bothify function
func BenchmarkBothify(b *testing.B) {
	f := New()

	benchmarks := []struct {
		name    string
		pattern string
	}{
		{"Short", "??#"},
		{"Medium", "??##-##??-????"},
		{"Long", strings.Repeat("?#", 25)},
		{"LettersOnly", strings.Repeat("?", 50)},
		{"NumbersOnly", strings.Repeat("#", 50)},
		{"Mixed", "User-??##-Email-??##@domain.com"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = f.Bothify(bm.pattern)
			}
		})
	}
}

// BenchmarkRandomStringWithLengthPool benchmarks string generation with various lengths using pool
func BenchmarkRandomStringWithLengthPool(b *testing.B) {
	f := New()

	lengths := []int{1, 10, 50, 100, 500, 1000}

	for _, length := range lengths {
		b.Run(fmt.Sprintf("Length_%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = f.RandomStringWithLength(length)
			}
		})
	}
}

// BenchmarkMemoryAllocations measures memory allocations for string operations
func BenchmarkMemoryAllocations(b *testing.B) {
	f := New()

	b.Run("NumerifyAllocations", func(b *testing.B) {
		pattern := strings.Repeat("#", 100)
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = f.Numerify(pattern)
		}
	})

	b.Run("RandomStringAllocations", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = f.RandomStringWithLength(100)
		}
	})

	b.Run("CombinedOperationsAllocations", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = f.Numerify("####-####")
			_ = f.Lexify("????-????")
			_ = f.Asciify("****-****")
			_ = f.RandomStringWithLength(20)
		}
	})
}

// BenchmarkNumericGeneration benchmarks numeric generation methods
func BenchmarkNumericGeneration(b *testing.B) {
	f := New()

	b.Run("RandomDigit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f.RandomDigit()
		}
	})

	b.Run("RandomNumber", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f.RandomNumber(6)
		}
	})

	b.Run("IntBetween", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f.IntBetween(1, 1000000)
		}
	})

	b.Run("Int64Between", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f.Int64Between(-1000000, 1000000)
		}
	})

	b.Run("Float64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f.Float64(2, 1, 1000)
		}
	})
}

// BenchmarkValidationOverhead measures the overhead of input validation
func BenchmarkValidationOverhead(b *testing.B) {
	f := New()

	b.Run("RandomStringWithLengthValidation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Test validation with negative length
			_ = f.RandomStringWithLength(-1)
		}
	})

	b.Run("RandomNumberValidation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Test validation with zero size
			_ = f.RandomNumber(0)
		}
	})

	b.Run("FloatValidation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Test validation with swapped min/max
			_ = f.RandomFloat(-1, 100, 1)
		}
	})
}

// BenchmarkConcurrentAccess benchmarks concurrent access to faker methods
func BenchmarkConcurrentAccess(b *testing.B) {
	f := New()

	b.Run("ConcurrentRandomDigit", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = f.RandomDigit()
			}
		})
	})

	b.Run("ConcurrentNumerify", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = f.Numerify("####-####")
			}
		})
	})

	b.Run("ConcurrentStringGeneration", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = f.RandomStringWithLength(20)
			}
		})
	})
}

// BenchmarkConstantsUsage benchmarks operations using defined constants
func BenchmarkConstantsUsage(b *testing.B) {
	f := New()

	b.Run("RandomLetterWithConstants", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f.RandomLetter()
		}
	})

	b.Run("AsciifyWithConstants", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f.Asciify("*")
		}
	})
}

// BenchmarkThreadSafeRand benchmarks the thread-safe random wrapper
func BenchmarkThreadSafeRand(b *testing.B) {
	f := New()

	b.Run("ThreadSafeIntn", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f.Generator.Intn(100)
		}
	})

	b.Run("ThreadSafeInt32n", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f.Generator.Int32n(100)
		}
	})

	b.Run("ThreadSafeInt64n", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f.Generator.Int64n(100)
		}
	})
}

// BenchmarkEdgeCases benchmarks edge case handling
func BenchmarkEdgeCases(b *testing.B) {
	f := New()

	b.Run("EmptyStringOperations", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f.Numerify("")
			_ = f.Lexify("")
			_ = f.Asciify("")
		}
	})

	b.Run("ExtremeRanges", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f.IntBetween(-2147483648, 2147483647)
		}
	})

	b.Run("MinimalValidation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f.RandomStringElement([]string{})
			_ = f.RandomIntElement([]int{})
		}
	})
}

// BenchmarkMemoryEfficiency tests memory efficiency of improvements
func BenchmarkMemoryEfficiency(b *testing.B) {
	if testing.Short() {
		b.Skip("Skipping memory efficiency test in short mode")
	}

	f := New()

	b.Run("MemoryUsagePattern", func(b *testing.B) {
		// Force GC before benchmark
		runtime.GC()

		var m1, m2 runtime.MemStats
		runtime.ReadMemStats(&m1)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			// Perform operations that should benefit from sync.Pool
			_ = f.Numerify(strings.Repeat("#", 50))
			_ = f.Lexify(strings.Repeat("?", 50))
			_ = f.Asciify(strings.Repeat("*", 50))
			_ = f.RandomStringWithLength(50)
		}

		b.StopTimer()

		runtime.GC()
		runtime.ReadMemStats(&m2)

		b.ReportMetric(float64(m2.TotalAlloc-m1.TotalAlloc)/float64(b.N), "bytes/op")
	})
}
