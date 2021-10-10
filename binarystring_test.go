package faker

import (
	"testing"
)

//check if given string is binary or not
func CheckIfBinaryString(binary_string string) bool {

	for i := 0; i < len(binary_string); i++ {
		if binary_string[i] != '1' && binary_string[i] != '0' {
			return false
		}
	}
	return true
}

//tests BinaryStringOfLength()
func TestBinaryStringOfLength(t *testing.T) {
	f := New().BinaryString()
	input_length := 11
	binary_string := f.BinaryStringOfLength(input_length)

	Expect(t, len(binary_string), input_length)
	Expect(t, CheckIfBinaryString(binary_string), true)
}
