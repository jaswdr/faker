package faker

import (
	"strings"
)

// Crypto is a faker struct for generating bitcoin data
type Crypto struct {
	Faker *Faker
}

var (
	bitcoinMin = 26
	bitcoinMax = 35
	ethLen     = 42
	ethPrefix  = "0x"
)

// Checks whether the ascii value provided is in the exclusion for bitcoin.
func inExclusionBitcoin(ascii int) bool {
	switch ascii {
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
func getAlnumRange(f *Faker) (int, int) {
	dec := f.IntBetween(0, 2)
	if dec == 0 {
		// digit
		return 48, 57
	} else if dec == 1 {
		// upper
		return 65, 90
	}
	// lower
	return 97, 122
}

// Helper function to return a bitcoin address with a given prefix and length
func randBitcoin(length int, prefix string, f *Faker) string {
	address := []string{prefix}

	for i := 0; i < length; i++ {
		asciiStart, asciiEnd := getAlnumRange(f)
		val := f.IntBetween(asciiStart, asciiEnd)
		if inExclusionBitcoin(val) {
			val++
		}
		address = append(address, string(rune(val)))
	}
	return strings.Join(address, "")
}

// Helper function to return an Ethereum address.
func randEth(length int, prefix string, f *Faker) string {
	address := []string{prefix}

	for i := 0; i < length; i++ {
		asciiStart, asciiEnd := getAlnumRange(f)
		val := f.IntBetween(asciiStart, asciiEnd)
		address = append(address, string(rune(val)))
	}
	return strings.Join(address, "")

}

// P2PKH generates a P2PKH bitcoin address.
func (c Crypto) P2PKH() string {
	length := c.Faker.IntBetween(bitcoinMin, bitcoinMax)
	// subtract 1 for prefix
	return randBitcoin(length-1, "1", c.Faker)
}

// P2PKHWithLength generates a P2PKH bitcoin address with specified length.
func (c Crypto) P2PKHWithLength(length int) string {
	return randBitcoin(length-1, "1", c.Faker)
}

// P2SH generates a P2SH bitcoin address.
func (c Crypto) P2SH() string {
	length := c.Faker.IntBetween(bitcoinMin, bitcoinMax)
	// subtract 1 for prefix
	return randBitcoin(length-1, "3", c.Faker)
}

// P2SHWithLength generates a P2PKH bitcoin address with specified length.
func (c Crypto) P2SHWithLength(length int) string {
	return randBitcoin(length-1, "3", c.Faker)
}

// Bech32 generates a Bech32 bitcoin address
func (c Crypto) Bech32() string {
	length := c.Faker.IntBetween(bitcoinMin, bitcoinMax)
	// subtract 1 for prefix
	return randBitcoin(length-3, "bc1", c.Faker)
}

// Bech32WithLength generates a Bech32 bitcoin address with specified length.
func (c Crypto) Bech32WithLength(length int) string {
	return randBitcoin(length-3, "bc1", c.Faker)
}

// RandomBitcoin returns an address of either Bech32, P2PKH, or P2SH type.
func (c Crypto) RandomBitcoin() string {
	dec := c.Faker.IntBetween(0, 2)
	if dec == 0 {
		return c.Bech32()
	} else if dec == 1 {
		return c.P2SH()
	}
	return c.P2PKH()
}

func (c Crypto) RandomEth() string {
	return randEth(ethLen-2, ethPrefix, c.Faker)
}
