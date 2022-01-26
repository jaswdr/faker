package faker

import (
	"testing"
	"strings"
)

var (
	bannedBitcoin = []string{"O", "I", "l", "0"}
)

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