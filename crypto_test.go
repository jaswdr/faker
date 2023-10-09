package faker

import (
	"strings"
	"testing"
)

func TestBitcoinAddress(t *testing.T) {
	c := New().Crypto()
	addr := c.BitcoinAddress()
	Expect(t, false, addr == "")
}

func TestEtheriumAddress(t *testing.T) {
	c := New().Crypto()
	addr := c.EtheriumAddress()
	Expect(t, false, addr == "")
	Expect(t, true, strings.HasPrefix(addr, "0x"))
}

func TestP2PKHAddress(t *testing.T) {
	c := New().Crypto()
	addr := c.P2PKHAddress()
	Expect(t, false, addr == "")
	Expect(t, true, strings.HasPrefix(addr, "1"))
}

func TestP2PKHAddressWithLength(t *testing.T) {
	c := New().Crypto()
	addr := c.P2PKHAddressWithLength(10)
	Expect(t, false, addr == "")
	Expect(t, true, strings.HasPrefix(addr, "1"))
}

func TestP2SHAddress(t *testing.T) {
	c := New().Crypto()
	addr := c.P2SHAddress()
	Expect(t, false, addr == "")
	Expect(t, true, strings.HasPrefix(addr, "3"))
}

func TestP2SHAddressWithLength(t *testing.T) {
	c := New().Crypto()
	addr := c.P2SHAddressWithLength(10)
	Expect(t, false, addr == "")
	Expect(t, true, strings.HasPrefix(addr, "3"))
}

func TestBech32Address(t *testing.T) {
	c := New().Crypto()
	addr := c.Bech32Address()
	Expect(t, false, addr == "")
	Expect(t, true, strings.HasPrefix(addr, "bc1"))
}

func TestBech32AddressWithLength(t *testing.T) {
	c := New().Crypto()
	addr := c.Bech32AddressWithLength(10)
	Expect(t, false, addr == "")
	Expect(t, true, strings.HasPrefix(addr, "bc1"))
}