package faker

import (
	"strings"
	"testing"
)

var (
	bannedBitcoin = []string{"O", "I", "l", "0"}
)

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
	Expect(t, true, strings.HasPrefix(addr, "1"))
	for i := 0; i < len(bannedBitcoin); i++ {
		Expect(t, true, !strings.Contains(addr, bannedBitcoin[i]))
	}
}

func TestP2PKHWithLength(t *testing.T) {
	c := New().Crypto()
	length := c.Faker.IntBetween(26, 62)
	addr := c.P2PKHWithLength(length)
	Expect(t, true, len(addr) == length)
	Expect(t, true, strings.HasPrefix(addr, "1"))
}

func TestP2SH(t *testing.T) {
	c := New().Crypto()
	addr := c.P2SH()
	Expect(t, true, len(addr) >= bitcoinMin)
	Expect(t, true, len(addr) <= bitcoinMax)
	Expect(t, true, strings.HasPrefix(addr, "3"))
	for i := 0; i < len(bannedBitcoin); i++ {
		Expect(t, true, !strings.Contains(addr, bannedBitcoin[i]))
	}
}

func TestP2SHWithLength(t *testing.T) {
	c := New().Crypto()
	length := c.Faker.IntBetween(26, 62)
	addr := c.P2SHWithLength(length)
	Expect(t, true, len(addr) == length)
	Expect(t, true, strings.HasPrefix(addr, "3"))
}

func TestBech32(t *testing.T) {
	c := New().Crypto()
	addr := c.Bech32()
	Expect(t, true, len(addr) >= bitcoinMin)
	Expect(t, true, len(addr) <= bitcoinMax)
	Expect(t, true, strings.HasPrefix(addr, "bc1"))
	for i := 0; i < len(bannedBitcoin); i++ {
		Expect(t, true, !strings.Contains(addr, bannedBitcoin[i]))
	}
}

func TestBech32WithLength(t *testing.T) {
	c := New().Crypto()
	length := c.Faker.IntBetween(26, 62)
	addr := c.Bech32WithLength(length)
	Expect(t, true, len(addr) == length)
	Expect(t, true, strings.HasPrefix(addr, "bc1"))
}
