package faker

import (
	"fmt"
	"strings"
	"testing"
)

var (
	bannedBitcoin      = []string{"O", "I", "l", "0"}
	validBitcoinPrefix = map[string]string{
		"p2pkh":  "1",
		"p2sh":   "3",
		"bech32": "bc1",
	}
	validEthPrefix = "0x"
)

type GeneratorMock struct {
	local int
}

func (g GeneratorMock) Intn(_ int) int {
	return g.local
}

func (g GeneratorMock) Int32n(_ int32) int32 {
	return int32(g.local)
}

func (g GeneratorMock) Int64n(_ int64) int64 {
	return int64(g.local)
}

func (g GeneratorMock) Uintn(_ uint) uint {
	return uint(g.local)
}

func (g GeneratorMock) Uint32n(_ uint32) uint32 {
	return uint32(g.local)
}

func (g GeneratorMock) Uint64n(_ uint64) uint64 {
	return uint64(g.local)
}

func (g GeneratorMock) Int() int {
	return g.local
}

type TestCaseAlnum struct {
	desc     string
	localInt int
	assert   func(t *testing.T, a int, b int)
}

type TestCaseRandomBitcoin struct {
	desc              string
	localInt          int
	expectedSubstring string
}

func TestIsInExclusionZone(t *testing.T) {
	c := New().Crypto()
	for _, address := range bannedBitcoin {
		Expect(t, true, c.isInExclusionZone(int(rune(address[0]))))
	}
	// take any banned rune and + 1 it to get a valid character
	Expect(t, false, c.isInExclusionZone(int(rune(bannedBitcoin[0][0]))+1))
}

func TestGenerateBicoinAddress(t *testing.T) {
	c := New().Crypto()
	length := c.Faker.IntBetween(5, 10)
	Expect(t, length+1, len(c.generateBicoinAddress(length, "a", c.Faker)))
}

func TestP2PKHAddress(t *testing.T) {
	c := New().Crypto()
	addr := c.P2PKHAddress()
	Expect(t, true, len(addr) >= bitcoinMin)
	Expect(t, true, len(addr) <= bitcoinMax)
	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["p2pkh"]))
	for i := 0; i < len(bannedBitcoin); i++ {
		Expect(t, true, !strings.Contains(addr, bannedBitcoin[i]))
	}
}

func TestP2PKHAddressWithLength(t *testing.T) {
	c := New().Crypto()
	length := c.Faker.IntBetween(26, 62)
	addr := c.P2PKHAddressWithLength(length)
	Expect(t, true, len(addr) == length)
	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["p2pkh"]))
}

func TestP2SHAddress(t *testing.T) {
	c := New().Crypto()
	addr := c.P2SHAddress()
	Expect(t, true, len(addr) >= bitcoinMin)
	Expect(t, true, len(addr) <= bitcoinMax)
	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["p2sh"]))
	for i := 0; i < len(bannedBitcoin); i++ {
		Expect(t, true, !strings.Contains(addr, bannedBitcoin[i]))
	}
}

func TestP2SHAddressWithLength(t *testing.T) {
	c := New().Crypto()
	length := c.Faker.IntBetween(26, 62)
	addr := c.P2SHAddressWithLength(length)
	Expect(t, true, len(addr) == length)
	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["p2sh"]))
}

func TestBech32Address(t *testing.T) {
	c := New().Crypto()
	addr := c.Bech32Address()
	Expect(t, true, len(addr) >= bitcoinMin)
	Expect(t, true, len(addr) <= bitcoinMax)
	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["bech32"]))
	for i := 0; i < len(bannedBitcoin); i++ {
		Expect(t, true, !strings.Contains(addr, bannedBitcoin[i]))
	}
}

func TestBech32AddressWithLength(t *testing.T) {
	c := New().Crypto()
	length := c.Faker.IntBetween(26, 62)
	addr := c.Bech32AddressWithLength(length)
	Expect(t, true, len(addr) == length)
	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["bech32"]))
}

func TestEtheriumAddress(t *testing.T) {
	c := New().Crypto()
	addr := c.EtheriumAddress()
	Expect(t, true, len(addr) == ethLen)
	Expect(t, true, strings.HasPrefix(addr, ethPrefix))
}

func TestAlgorithmRange(t *testing.T) {
	for k, tc := range []TestCaseAlnum{
		{
			// The Description of the test case
			desc:     "Test Get Digit 0-9",
			localInt: 0,
			// Our anticipated result
			assert: func(t *testing.T, a int, b int) {
				Expect(t, true, a == int('0'))
				Expect(t, true, b == int('9'))
			},
		},
		{
			desc:     "Test Get Uppercase A-Z",
			localInt: 1,
			assert: func(t *testing.T, a int, b int) {
				Expect(t, true, a == int('A'))
				Expect(t, true, b == int('Z'))
			},
		},
		{
			desc:     "Test Get Lowercase a-z",
			localInt: 2,
			assert: func(t *testing.T, a int, b int) {
				Expect(t, true, a == int('a'))
				Expect(t, true, b == int('z'))
			},
		},
	} {
		t.Run(fmt.Sprintf("case=%d/description=%s", k, tc.desc), func(t *testing.T) {
			// Use our mock here instead of using a seed.
			gen := GeneratorMock{}
			gen.local = tc.localInt
			// populate the generator with our mock as it is an interface.
			c := Faker{Generator: gen}
			a, b := c.Crypto().algorithmRange()
			tc.assert(t, a, b)
		})
	}
}

func TestRandomBitcoin(t *testing.T) {
	for k, tc := range []TestCaseRandomBitcoin{
		{
			// The Description of the test case
			desc:     "Test Get Bech32",
			localInt: 0,
			// Our anticipated result
			expectedSubstring: "bc1",
		},
		{
			// The Description of the test case
			desc:     "Test Get P2SH",
			localInt: 1,
			// Our anticipated result
			expectedSubstring: "3",
		},
		{
			// The Description of the test case
			desc:     "Test Get P2PKH",
			localInt: 2,
			// Our anticipated result
			expectedSubstring: "1",
		},
	} {
		t.Run(fmt.Sprintf("case=%d/description=%s", k, tc.desc), func(t *testing.T) {
			// Use our mock here instead of using a seed.
			gen := GeneratorMock{}
			gen.local = tc.localInt
			// populate the generator with our mock as it is an interface.
			c := Faker{Generator: gen}
			rs := c.Crypto().BitcoinAddress()
			Expect(t, true, strings.HasPrefix(rs, tc.expectedSubstring))
			Expect(t, true, len(rs) >= bitcoinMin)
			Expect(t, true, len(rs) <= bitcoinMax)
		})
	}
}
