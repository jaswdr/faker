package faker

import (
	"testing"
)

//check if given string is binary or not
func CheckIfBinaryString(str string) bool {

	for i := 0; i < len(str); i++ {
		if str[i] != '1' && str[i] != '0' {
			return false
		}
	}
	return true
}

//tests BinaryStringOfLength()
func TestBinaryStringOfLength(t *testing.T) {
	f := New().BinaryString()
	inputLength := 11
	str := f.BinaryStringOfLength(inputLength)

	Expect(t, len(str), inputLength)
	Expect(t, CheckIfBinaryString(str), true)
}
