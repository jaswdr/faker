package faker

import (
	"testing"
)

//tests BinaryString()
func TestBinaryString(t *testing.T) {
	f := New().BinaryString()
	inputLength := 11
	str := f.BinaryString(inputLength)

	Expect(t, len(str), inputLength)

	isStringValid := true
	for i := 0; i < len(str); i++ {
		if str[i] != '1' && str[i] != '0' {
			isStringValid = false
			break
		}
	}
	Expect(t, isStringValid, true)
}
