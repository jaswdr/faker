package faker

import (
	"fmt"
	"testing"
)

//Check if input is hex encoded string
func checkIfHexString(s string) bool {
	for i := 0; i < len(s); i++ {
		if !((s[i] >= 97 && s[i] <= 102) || (s[i] >= 48 && s[i] <= 57)) {
			return false
		}
	}
	return true
}

//tests SHA256()
func TestSHA256(t *testing.T) {
	hash := New().Hash()
	s := hash.SHA256()
	Expect(t, fmt.Sprintf("%T", s), "string")
	Expect(t, 64, len(s))
	Expect(t, true, checkIfHexString(s))
}

//tests SHA512()
func TestSHA512(t *testing.T) {
	hash := New().Hash()
	s := hash.SHA512()
	Expect(t, fmt.Sprintf("%T", s), "string")
	Expect(t, 128, len(s))
	Expect(t, true, checkIfHexString(s))
}

//tests MD5()
func TestMD5(t *testing.T) {
	hash := New().Hash()
	s := hash.MD5()
	Expect(t, fmt.Sprintf("%T", s), "string")
	Expect(t, 32, len(s))
	Expect(t, true, checkIfHexString(s))
}
