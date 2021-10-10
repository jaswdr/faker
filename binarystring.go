package faker

type BinaryString struct {
	faker *Faker
}

//BinaryNumberOfLength returns random binary string with given input length
func (bn BinaryString) BinaryStringOfLength(l int) string {
	var bs string
	for i := 0; i < l; i++ {
		bs += bn.faker.RandomStringElement([]string{"0", "1"})
	}
	return bs
}
