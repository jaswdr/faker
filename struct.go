package faker

import (
	"fmt"
	"reflect"
	"strconv"
)

// Struct is a faker struct for Struct
type Struct struct {
	Faker *Faker
}

// Fill elements of a struct with random data
func (s Struct) Fill(value interface{}) {
	reflectedType := reflect.TypeOf(value)
	reflectedValue := reflect.ValueOf(value)
	s.typeHandler(reflectedType, reflectedValue, "", 0, "", map[string]bool{})
}

func (s Struct) typeHandler(valueType reflect.Type, value reflect.Value, function string, size int, parentName string, typesSeen map[string]bool) bool {
	kind := valueType.Kind()
	switch kind {
	case reflect.Ptr:
		return s.pointerHandler(valueType, value, function, parentName, typesSeen)
	case reflect.Struct:
		return s.structHandler(valueType, value, parentName, typesSeen)
	case reflect.String:
		return s.stringHandler(valueType, value, function)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return s.uintHandler(valueType, value, function)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return s.intHandler(valueType, value, function)
	case reflect.Float32, reflect.Float64:
		return s.floatHandler(valueType, value, function)
	case reflect.Bool:
		return s.boolHandler(valueType, value)
	case reflect.Array, reflect.Slice:
		return s.sliceHandler(valueType, value, function, size, parentName, typesSeen)
	default:
		return false
	}
}

func (s Struct) structHandler(valueType reflect.Type, value reflect.Value, parentName string, typesSeen map[string]bool) bool {
	result := true
	numberOfFields := valueType.NumField()
	for fieldIndex := 0; fieldIndex < numberOfFields; fieldIndex++ {
		structType := valueType.Field(fieldIndex)
		structValue := value.Field(fieldIndex)

		// Check if type was seen before
		if structType.Type.Name() == "" {
			if parentName == "" {
				parentName = structType.Name
			}

			valueToCompare := fmt.Sprintf("%v.%v", parentName, structType.Name)
			found := false
			_, found = typesSeen[valueToCompare]
			if found {
				result = false
				continue
			}
			typesSeen[valueToCompare] = true
		}

		tag, ok := structType.Tag.Lookup("fake")
		if ok && tag == "skip" {
			continue
		}

		if structValue.CanSet() {
			// Check if fakesize is set
			size := -1 // Set to -1 to indicate fakesize was not set
			fakeSize, ok := structType.Tag.Lookup("fakesize")
			if ok {
				var err error
				size, err = strconv.Atoi(fakeSize)
				if err != nil {
					size = s.Faker.IntBetween(1, 10)
				}
			}

			if !s.typeHandler(structType.Type, structValue, tag, size, parentName, typesSeen) {
				result = false
			}
		}

		if structType.Name == parentName {
			parentName = ""
		}
	}

	return result
}

func (s Struct) pointerHandler(valueType reflect.Type, value reflect.Value, function string, parentName string, typesSeen map[string]bool) bool {
	elementType := valueType.Elem()
	if value.IsNil() {
		newValue := reflect.New(elementType)
		if !s.typeHandler(elementType, newValue.Elem(), function, 0, parentName, typesSeen) {
			return false
		}
		fmt.Printf("Setting %v to %v\n", valueType, newValue)
		value.Set(newValue)
	}

	return s.typeHandler(elementType, value.Elem(), function, 0, parentName, typesSeen)
}

func (s Struct) sliceHandler(valueType reflect.Type, value reflect.Value, function string, size int, parentName string, typesSeen map[string]bool) bool {
	// If you cant even set it dont even try
	if !value.CanSet() {
		return false
	}

	// Grab original size to use if needed for sub arrays
	originalSize := size

	// If the value has a cap and is less than the size
	// use that instead of the requested size
	elementCapacity := value.Cap()
	if elementCapacity == 0 && size == -1 {
		size = s.Faker.IntBetween(1, 10)
	} else if elementCapacity != 0 && (size == -1 || elementCapacity < size) {
		size = elementCapacity
	}

	// If the value is empty and the size is not -1, create a new slice
	if elementCapacity == 0 && size > 0 {
		value.Set(reflect.MakeSlice(valueType, size, size))
	}

	// Get the element type
	elementType := valueType.Elem()
	newValue := reflect.New(elementType)
	elementValue := newValue.Elem()

	for i := 0; i < size; i++ {
		if !s.typeHandler(elementType, elementValue, function, originalSize, parentName, typesSeen) {
			return false
		}

		value.Index(i).Set(reflect.Indirect(newValue))
		if value.Len() != 0 {
			value.Index(i).Set(reflect.Indirect(newValue))
		} else {
			value.Set(reflect.Append(reflect.Indirect(value), reflect.Indirect(newValue)))
		}
	}

	return true
}

func (s Struct) stringHandler(_ reflect.Type, value reflect.Value, function string) bool {
	if function == "" {
		value.SetString(s.Faker.UUID().V4())
		return true
	}

	value.SetString(s.Faker.Bothify(function))
	return true
}

func (s Struct) intHandler(valueType reflect.Type, value reflect.Value, function string) bool {
	if function != "" {
		intValue, err := strconv.ParseInt(s.Faker.Numerify(function), 10, 64)
		if err == nil {
			value.SetInt(intValue)
			return true
		}
	}

	// If no function or error converting to int, set with random value
	switch valueType.Kind() {
	case reflect.Int:
		value.SetInt(s.Faker.Int64())
	case reflect.Int8:
		value.SetInt(int64(s.Faker.Int8()))
	case reflect.Int16:
		value.SetInt(int64(s.Faker.Int16()))
	case reflect.Int32:
		value.SetInt(int64(s.Faker.Int32()))
	case reflect.Int64:
		value.SetInt(s.Faker.Int64())
	default:
		return false
	}

	return true
}

func (s Struct) uintHandler(valueType reflect.Type, value reflect.Value, function string) bool {
	if function != "" {
		uintValue, err := strconv.ParseUint(s.Faker.Numerify(function), 10, 64)
		if err == nil {
			value.SetUint(uintValue)
			return true
		}
	}

	// If no function or error converting to uint, set with random value
	switch valueType.Kind() {
	case reflect.Uint:
		value.SetUint(s.Faker.UInt64())
	case reflect.Uint8:
		value.SetUint(uint64(s.Faker.UInt8()))
	case reflect.Uint16:
		value.SetUint(uint64(s.Faker.UInt16()))
	case reflect.Uint32:
		value.SetUint(uint64(s.Faker.UInt32()))
	case reflect.Uint64:
		value.SetUint(s.Faker.UInt64())
	default:
		return false
	}

	return true
}

func (s Struct) floatHandler(valueType reflect.Type, value reflect.Value, function string) bool {
	if function != "" {
		floatValue, err := strconv.ParseFloat(s.Faker.Numerify(function), 64)
		if err == nil {
			value.SetFloat(floatValue)
			return true
		}
	}

	// If no function or error converting to float, set with random value
	switch valueType.Kind() {
	case reflect.Float64:
		value.SetFloat(s.Faker.Float64(2, 0, 100))
	case reflect.Float32:
		value.SetFloat(s.Faker.Float64(2, 0, 100))
	default:
		return false
	}

	return true
}

func (s Struct) boolHandler(_ reflect.Type, value reflect.Value) bool {
	value.SetBool(s.Faker.Bool())
	return true
}
