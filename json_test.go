package faker

import (
	"testing"
	"encoding/json"
)

func TestJsonString(t *testing.T) {
	faker := New()
	j := faker.Json()

	for i := 0; i < 10; i++ {
		result := j.String()

		// Attempt to unmarshal the result into a map[string]interface{}
		var data map[string]interface{}
		err := json.Unmarshal([]byte(result), &data)
		Expect(t, err, nil)

		// Ensure that the result is a valid JSON object
		Expect(t, len(data) > 0, true)

		// Ensure that all attribute values are valid JSON types
		for _, value := range data {
			switch value.(type) {
			case string, float64, bool, nil:
				// Valid JSON types
			case []interface{}:
				// Valid JSON array type
			case map[string]interface{}:
				// Valid JSON object type
			default:
				t.FailNow()
			}
		}
	}
}

func TestJsonObject(t *testing.T) {
	faker := New()
	j := faker.Json()

	for i := 0; i < 10; i++ {
		data := j.Object()

		// Ensure that the result is not nil
		NotExpect(t, data, nil)

		// Ensure that the result is a valid JSON object
		Expect(t, len(data) > 0, true)

		// Ensure that all attribute values are valid JSON types
		for _, value := range data {
			switch value.(type) {
			case string, float64, bool, nil:
				// Valid JSON types
			case []interface{}:
				// Valid JSON array type
			case map[string]interface{}:
				// Valid JSON object type
			default:
				t.FailNow()
			}
		}
	}
}