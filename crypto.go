package faker

import (
	"strings"
)

// Bitcoin is a faker struct for generating bitcoin data
type Crypto struct {
	Faker *Faker
}

var (
	bitcoinMin = 26
	bitcoinMax = 35
)

func inExclusion(check int) bool {
	// switch probably seems to be the fastest as compared to list/map
	// https://stackoverflow.com/questions/15323767/does-go-have-if-x-in-construct-similar-to-python
	switch check {
		// Ascii for uppercase letter "O", uppercase letter "I", lowercase letter "l", and the number "0" 
		case
			48,
			73,
			79,
			108:
			return true
	}
	return false
}

// Decide whether to get digit, uppercase, or lowercase. returns the ascii range to do IntBetween on
func getBitcoinRange(f *Faker) (int, int) {
	dec := f.IntBetween(0, 2)
	if dec == 0{
		return 48, 57
	} else if dec == 1 {
		return 65, 90
	} else {
		return 97,122
	}
}

// Helper function to return a bitcoin
func randBitcoin(length int, prefix string, f *Faker) string {
	address := []string{prefix}

	for i := 0; i < length; i++ {
		asciiStart, asciiEnd := getBitcoinRange(f)
		val := f.IntBetween(asciiStart, asciiEnd)
		if inExclusion(val){
			val += 1
		}
		address = append(address, string(rune(val)))
	}
	return strings.Join(address, "")
}

// Generates P2PKH bitcoin address.
// Based on https://github.com/jaswdr/faker/issues/78#issuecomment-1020662826
func (c Crypto) P2PKH() string{
	length := c.Faker.IntBetween(bitcoinMin, bitcoinMax)
	// subtrace 1 for prefix
	return randBitcoin(length-1, "1", c.Faker)
}
/*
func (c Crypto) P2PKHVariableLength(int length){

}
*/