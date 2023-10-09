package faker

import (
	"encoding/json"
	"strconv"
)

// Json is a faker struct for json files
type Json struct {
	Faker *Faker
}

var (
	attributesTypes                      = []string{"string", "number", "object", "array", "boolean", "null"}
	attributesTypesWithoutArrayAndObject = []string{"string", "number", "boolean", "null"}
)

func (j *Json) randomAttributeValueFromListAsString(validAttributesTypes []string) string {
	attributeType := j.Faker.RandomStringElement(validAttributesTypes)
	switch attributeType {
	case "string":
		return "\"" + j.Faker.Lorem().Word() + "\""
	case "number":
		number := strconv.Itoa(j.Faker.RandomNumber(j.Faker.RandomDigit()))
		return number
	case "object":
		// Avoid having multiple nested objects by not using object and array as valid attribute types
		return j.randomJsonObjectAsString(attributesTypesWithoutArrayAndObject)
	case "array":
		objects := ""
		for i := 0; i < j.Faker.IntBetween(1, 10); i++ {
			if objects != "" {
				objects += ", "
			}
			// Avoid having multiple nested objects by not using object and array as valid attribute types
			objects += j.randomJsonObjectAsString(attributesTypesWithoutArrayAndObject)
		}
		return "[" + objects + "]"
	case "boolean":
		return j.Faker.RandomStringElement([]string{"true", "false"})
	case "null":
		return "null"
	}

	panic("Invalid attribute type: " + attributeType)
}

func (j *Json) randomJsonObjectAsString(validAttributesTypes []string) string {
	numberAttributes := j.Faker.IntBetween(1, 10)
	attributes := make([]string, numberAttributes)
	for i := 0; i < numberAttributes; i++ {
		attributeName := j.Faker.Lorem().Word()
		attributeValue := j.randomAttributeValueFromListAsString(validAttributesTypes)
		attributes[i] = "\"" + attributeName + "\": " + attributeValue
	}

	result := "{"
	for i := 0; i < len(attributes); i++ {
		if i > 0 {
			result += ", "
		}
		result += attributes[i]
	}
	result += "}"
	return result
}

// String generates a random json string
func (j *Json) String() string {
	return j.randomJsonObjectAsString(attributesTypes)
}

// Object generates a random json object
func (j *Json) Object() map[string]interface{} {
	result := j.String()
	var data map[string]interface{}
	json.Unmarshal([]byte(result), &data)
	return data
}
