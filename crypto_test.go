package faker

import (
	"testing"
	"strings"
)

var (
	bannedBitcoin = []string{"O", "I", "l", "0"}
)

func TestInExclusion(t *testing.T){
	for _, c := range bannedBitcoin{
		Expect(t, true, inExclusion(int(rune(c[0]))))
	}
	// take any banned rune and + 1 it to get a valid character
	Expect(t, false, inExclusion(int(rune(bannedBitcoin[0][0]))+1))
}

func TestRandBitcoin(t *testing.T){
	c := New().Crypto()
	length := c.Faker.IntBetween(5,10)
	randAddr := randBitcoin(length, "a", c.Faker)
	Expect(t, len(randAddr), length+1)
}

func TestP2PKH(t *testing.T){
	c := New().Crypto()
	addr := c.P2PKH()
	Expect(t, len(addr) >= bitcoinMin, true)
	Expect(t, len(addr) <= bitcoinMax, true)
	Expect(t, strings.HasPrefix(addr, "1"), true)
	for i := 0; i < len(bannedBitcoin); i++ {
		Expect(t, !strings.Contains(addr, bannedBitcoin[i]), true)
	}
}

func TestP2PKHWithLenght(t *testing.T){
	c := New().Crypto()
	length := c.Faker.IntBetween(26,62)
	addr := c.P2PKHWithLength(length)
	Expect(t, len(addr) == length, true)
}