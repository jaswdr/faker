package faker

// Crypto is a faker struct for generating bitcoin data
type Crypto struct {
	Faker *Faker
}

var (
	bitcoinAddresses = []string{
		"1JyxpLZzvYP2TyXaQV3J3vwajJz4hbxtRC",
		"1JyVFJVUNx8RQrjyNCGDe7BQ62wPyxU8bC",
		"1JVcPDeBfGP5PmNZwnJSfms7hKLncMSenV",
		"1FwiYqdgLH6w5XdB9QXgZGi7ZHGiyUYucT",
		"1H3nknk2Pdav9LjXyfLd8umqPC57ZQbN4",
		"1DMwH1FMx2yJ67az7ZTJqViBD6iz3vQkUK",
		"1JzvBX9Q86LbEcBxT58npYXS31QexVVMGG",
		"1HBTfs2QLK459tQrdpeQs4stR25GWwJTui",
		"1NZMxDpqB1ehEdvGJeAs9Gdmh3dXfUfAZB",
		"19ePAsmdkM4u9e3euzfnQa1AXEoD2UgmEj",
	}

	etheriumAddresses = []string{
		"0x83e1e8f10092d42db425D81c2e99f312a7E011aA",
		"0xd3E823D4C999e4ef9c92835eEF6906E519C13251",
		"0x4e3adfcdD456DDe868B1225aA0ed103Dd188B5F6",
		"0xbC39DCa632f8f7f2A94B095d48bcEE779d961728",
		"0xE0A6c75e545947E7Bb4dde2D8182762a4C698E5c",
		"0x45C20Fd8F6B07359750f92B79f1C41754Bd09Ac3",
		"0xA882bE0b4C10E91c3565EE01878A48F9B940f2c5",
		"0x8c2B7B23f01fcAD2946A3C214c4D96338A5eFD6D",
		"0x271253c6B815a07506719116262c1673692eD76E",
		"0x1e887dC08ba56e369E68987F9D82b44065677c87",
	}
)

// BitcoinAddress returns a valid address of either Bech32, P2PKH, or P2SH type.
func (c Crypto) BitcoinAddress() string {
	return c.Faker.RandomStringElement(bitcoinAddresses)
}

// EtheriumAddress returns a valid hexadecimal ethereum address of 42 characters.
func (c Crypto) EtheriumAddress() string {
	return c.Faker.RandomStringElement(etheriumAddresses)
}

// P2PKHAddress generates a P2PKH bitcoin address.
// Deprecated: Use BitcoinAddress instead.
func (c Crypto) P2PKHAddress() string {
	return "1" + c.BitcoinAddress()[1:]
}

// P2PKHAddressWithLength generates a P2PKH bitcoin address with specified length.
// Deprecated: Use BitcoinAddress instead.
func (c Crypto) P2PKHAddressWithLength(length int) string {
	return "1" + c.P2PKHAddress()[1:length-1]
}

// P2SHAddress generates a P2SH bitcoin address.
// Deprecated: Use BitcoinAddress instead.
func (c Crypto) P2SHAddress() string {
	return "3" + c.BitcoinAddress()[1:]
}

// P2SHAddressWithLength generates a P2PKH bitcoin address with specified length.
// Deprecated: Use BitcoinAddress instead.
func (c Crypto) P2SHAddressWithLength(length int) string {
	return "3" + c.P2SHAddress()[1:length-1]
}

// Bech32Address generates a Bech32 bitcoin address
// Deprecated: Use BitcoinAddress instead.
func (c Crypto) Bech32Address() string {
	return "bc1" + c.BitcoinAddress()[3:]
}

// Bech32AddressWithLength generates a Bech32 bitcoin address with specified length.
// Deprecated: Use BitcoinAddress instead.
func (c Crypto) Bech32AddressWithLength(length int) string {
	return "bc1" + c.Bech32Address()[3:length-3]
}