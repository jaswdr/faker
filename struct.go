package faker

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Struct is a faker struct for generating random data for struct fields
type Struct struct {
	Faker *Faker
}

// MaxRecursionDepth defines the default maximum depth for recursive structs
const MaxRecursionDepth = 32

type fakerFunction func() interface{}

var functions = map[string]fakerFunction{}

func RegisterFunction(name string, function fakerFunction) {
	functions[fmt.Sprintf("fn=%s", name)] = function
}

// Fill populates a struct with random data based on its type and tags
func (s Struct) Fill(v interface{}) {
	if v == nil {
		return
	}

	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return
	}

	s.fillValue(val.Elem(), "", -1, 0, MaxRecursionDepth)
}

func (s Struct) FillWithDepth(v interface{}, maxDepth int) {
	if v == nil {
		return
	}

	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return
	}

	s.fillValue(val.Elem(), "", -1, 0, maxDepth)
}

// fillValue recursively fills a reflect.Value with random data
func (s Struct) fillValue(v reflect.Value, function string, size int, depth int, maxDepth int) {
	if !v.CanSet() || depth > maxDepth {
		return
	}

	if strings.HasPrefix(function, "fn=") && s.fillFunction(v.Type(), v, function) {
		return
	}

	switch v.Kind() {
	case reflect.Pointer:
		s.fillPointer(v, function, size, depth, maxDepth)
	case reflect.Struct:
		s.fillStruct(v, depth, maxDepth)
	case reflect.String:
		s.fillString(v.Type(), v, function)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		s.fillUint(v.Type(), v, function)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		s.fillInt(v.Type(), v, function)
	case reflect.Float32, reflect.Float64:
		s.fillFloat(v.Type(), v, function)
	case reflect.Bool:
		s.fillBool(v.Type(), v)
	case reflect.Array, reflect.Slice:
		s.fillSlice(v, function, size, depth, maxDepth)
	case reflect.Map:
		s.fillMap(v.Type(), v, function, size, depth, maxDepth)
	}
}

// fillStruct fills a struct with random data for each field
func (s Struct) fillStruct(v reflect.Value, depth int, maxDepth int) {
	if v.Type().ConvertibleTo(reflect.TypeOf(time.Time{})) {
		v.Set(reflect.ValueOf(time.Unix(int64(s.Faker.Int32()), 0)).Convert(v.Type()))
		return
	}
	if v.Type().ConvertibleTo(reflect.TypeOf(big.Rat{})) {
		v.Set(reflect.ValueOf(*big.NewRat(s.Faker.Int64Between(1, math.MaxInt64), s.Faker.Int64Between(1, math.MaxInt64))).Convert(v.Type()))
		return
	}

	typ := v.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := v.Field(i)

		// Skip fields with "skip" tag
		if tag, ok := field.Tag.Lookup("fake"); ok && tag == "skip" {
			continue
		}

		// Get size from fakesize tag if present
		size := -1
		if sizeStr, ok := field.Tag.Lookup("fakesize"); ok {
			if parsedSize, err := strconv.Atoi(sizeStr); err == nil {
				size = parsedSize
			}
		}

		// Get function from fake tag if present
		function := ""
		if tag, ok := field.Tag.Lookup("fake"); ok {
			function = tag
		}

		s.fillValue(fieldVal, function, size, depth+1, maxDepth)
	}
}

// fillPointer handles pointer types by creating new instances if needed
func (s Struct) fillPointer(v reflect.Value, function string, size int, depth int, maxDepth int) {
	if v.IsNil() {
		v.Set(reflect.New(v.Type().Elem()))
	}
	s.fillValue(v.Elem(), function, size, depth+1, maxDepth)
}

// fillSlice handles array and slice types
func (s Struct) fillSlice(v reflect.Value, function string, size int, depth int, maxDepth int) {
	if !v.CanSet() {
		return
	}

	// Determine the size to use
	actualSize := size
	if actualSize == -1 {
		actualSize = s.Faker.IntBetween(defaultSliceMinSize, defaultSliceMaxSize)
	}
	if v.Cap() > 0 && (size == -1 || v.Cap() < size) {
		actualSize = v.Cap()
	}

	// Get element type
	elemType := v.Type().Elem()

	// Handle existing elements
	if v.Len() > 0 {
		for i := 0; i < actualSize; i++ {
			if i < v.Len() {
				s.fillValue(v.Index(i), function, size, depth+1, maxDepth)
			} else {
				elem := reflect.New(elemType).Elem()
				s.fillValue(elem, function, size, depth+1, maxDepth)
				v.Set(reflect.Append(v, elem))
			}
		}
		return
	}

	// Create new elements
	for i := 0; i < actualSize; i++ {
		elem := reflect.New(elemType).Elem()
		s.fillValue(elem, function, size, depth+1, maxDepth)
		v.Set(reflect.Append(v, elem))
	}
}

func (s Struct) fillString(_ reflect.Type, v reflect.Value, function string) {
	if function == "" {
		v.SetString(s.Faker.UUID().V4())
		return
	}

	v.SetString(s.Faker.Bothify(function))
}

func (s Struct) fillInt(t reflect.Type, v reflect.Value, function string) {
	if function != "" {
		i, err := strconv.ParseInt(s.Faker.Numerify(function), 10, 64)
		if err == nil {
			v.SetInt(i)
			return
		}
	}

	// If no function or error converting to int, set with random value
	switch t.Kind() {
	case reflect.Int:
		v.SetInt(s.Faker.Int64())
	case reflect.Int8:
		v.SetInt(int64(s.Faker.Int8()))
	case reflect.Int16:
		v.SetInt(int64(s.Faker.Int16()))
	case reflect.Int32:
		v.SetInt(int64(s.Faker.Int32()))
	case reflect.Int64:
		v.SetInt(s.Faker.Int64())
	}
}

func (s Struct) fillUint(t reflect.Type, v reflect.Value, function string) {
	if function != "" {
		u, err := strconv.ParseUint(s.Faker.Numerify(function), 10, 64)
		if err == nil {
			v.SetUint(u)
			return
		}
	}

	// If no function or error converting to uint, set with random value
	switch t.Kind() {
	case reflect.Uint:
		v.SetUint(s.Faker.UInt64())
	case reflect.Uint8:
		v.SetUint(uint64(s.Faker.UInt8()))
	case reflect.Uint16:
		v.SetUint(uint64(s.Faker.UInt16()))
	case reflect.Uint32:
		v.SetUint(uint64(s.Faker.UInt32()))
	case reflect.Uint64:
		v.SetUint(s.Faker.UInt64())
	}
}

func (s Struct) fillFloat(t reflect.Type, v reflect.Value, function string) {
	if function != "" {
		f, err := strconv.ParseFloat(s.Faker.Numerify(function), 64)
		if err == nil {
			v.SetFloat(f)
			return
		}
	}

	// If no function or error converting to float, set with random value
	switch t.Kind() {
	case reflect.Float64:
		v.SetFloat(s.Faker.Float64(2, 0, 100))
	case reflect.Float32:
		v.SetFloat(s.Faker.Float64(2, 0, 100))
	}
}

func (s Struct) fillBool(_ reflect.Type, v reflect.Value) {
	v.SetBool(s.Faker.Bool())
}

func (Struct) fillFunction(t reflect.Type, v reflect.Value, function string) bool {
	f, ok := functions[function]

	if !ok {
		return false
	}

	val := f()
	vType := reflect.TypeOf(val)

	if !vType.AssignableTo(t) {
		return false
	}

	vVal := reflect.ValueOf(val)
	v.Set(vVal)
	return true
}

func (s Struct) fillMap(t reflect.Type, v reflect.Value, function string, size int, depth int, maxDepth int) {
	if !v.CanSet() {
		return
	}
	mapType := reflect.MapOf(t.Key(), t.Elem())
	newMap := reflect.MakeMap(mapType)

	newSize := size
	if newSize == -1 {
		newSize = s.Faker.IntBetween(defaultSliceMinSize, defaultSliceMaxSize)
	}

	for i := 0; i < newSize; i++ {
		mapIndex := reflect.New(t.Key())
		s.fillValue(mapIndex.Elem(), function, size, depth+1, maxDepth)
		mapValue := reflect.New(t.Elem())
		s.fillValue(mapValue.Elem(), function, size, depth+1, maxDepth)
		newMap.SetMapIndex(mapIndex.Elem(), mapValue.Elem())
	}
	v.Set(newMap)
}
