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

func (g GeneratorMock) Intn(n int) int {
	return g.local
}

func (g GeneratorMock) Int() int {
	return g.local
}

type TestCaseAlnum struct {
	desc     string
	localInt int
	assert   func(t *testing.T, a int, b int)
}

func TestInExclusionBitcoin(t *testing.T) {
	for _, c := range bannedBitcoin {
		Expect(t, true, inExclusionBitcoin(int(rune(c[0]))))
	}
	// take any banned rune and + 1 it to get a valid character
	Expect(t, false, inExclusionBitcoin(int(rune(bannedBitcoin[0][0]))+1))
}

func TestRandBitcoin(t *testing.T) {
	c := New().Crypto()
	length := c.Faker.IntBetween(5, 10)
	randAddr := randBitcoin(length, "a", c.Faker)
	Expect(t, length+1, len(randAddr))
}

func TestP2PKH(t *testing.T) {
	c := New().Crypto()
	addr := c.P2PKH()
	Expect(t, true, len(addr) >= bitcoinMin)
	Expect(t, true, len(addr) <= bitcoinMax)
	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["p2pkh"]))
	for i := 0; i < len(bannedBitcoin); i++ {
		Expect(t, true, !strings.Contains(addr, bannedBitcoin[i]))
	}
}

func TestP2PKHWithLength(t *testing.T) {
	c := New().Crypto()
	length := c.Faker.IntBetween(26, 62)
	addr := c.P2PKHWithLength(length)
	Expect(t, true, len(addr) == length)
	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["p2pkh"]))
}

func TestP2SH(t *testing.T) {
	c := New().Crypto()
	addr := c.P2SH()
	Expect(t, true, len(addr) >= bitcoinMin)
	Expect(t, true, len(addr) <= bitcoinMax)
	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["p2sh"]))
	for i := 0; i < len(bannedBitcoin); i++ {
		Expect(t, true, !strings.Contains(addr, bannedBitcoin[i]))
	}
}

func TestP2SHWithLength(t *testing.T) {
	c := New().Crypto()
	length := c.Faker.IntBetween(26, 62)
	addr := c.P2SHWithLength(length)
	Expect(t, true, len(addr) == length)
	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["p2sh"]))
}

func TestBech32(t *testing.T) {
	c := New().Crypto()
	addr := c.Bech32()
	Expect(t, true, len(addr) >= bitcoinMin)
	Expect(t, true, len(addr) <= bitcoinMax)
	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["bech32"]))
	for i := 0; i < len(bannedBitcoin); i++ {
		Expect(t, true, !strings.Contains(addr, bannedBitcoin[i]))
	}
}

func TestBech32WithLength(t *testing.T) {
	c := New().Crypto()
	length := c.Faker.IntBetween(26, 62)
	addr := c.Bech32WithLength(length)
	Expect(t, true, len(addr) == length)
	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["bech32"]))
}

func TestRandomBitcoin(t *testing.T) {
	c := New().Crypto()
	addr := c.RandomBitcoin()
	Expect(t, true, len(addr) >= bitcoinMin)
	Expect(t, true, len(addr) <= bitcoinMax)
	in := false
	for _, pfx := range validBitcoinPrefix {
		if strings.HasPrefix(addr, pfx) {
			in = true
			break
		}
	}
	Expect(t, true, in)
}

func TestRandomEth(t *testing.T) {
	c := New().Crypto()
	addr := c.RandomEth()
	Expect(t, true, len(addr) == ethLen)
	Expect(t, true, strings.HasPrefix(addr, ethPrefix))
}

func TestGetAlnumRange(t *testing.T) {
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
			a, b := getAlnumRange(c.Crypto().Faker)

			tc.assert(t, a, b)
		})
	}
}
