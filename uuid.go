package faker

type UUID struct {
	Faker *Faker
}

func (u UUID) randomAFNumber() string {
	group := []string{
		"a", "b", "c", "d", "e", "f",
		"A", "B", "C", "D", "E", "F",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}

	return u.Faker.RandomStringElement(group)
}

func (u UUID) randomAFNumberCount(count int) (result string) {
	for i := 0; i < count; i++ {
		result += u.randomAFNumber()
	}

	return
}

func (u UUID) V4() (uuid string) {
	uuid += u.randomAFNumberCount(8)
	uuid += "-"
	uuid += u.randomAFNumberCount(4)
	uuid += "-"
	uuid += "4"
	uuid += u.randomAFNumberCount(3)
	uuid += "-"
	uuid += u.Faker.RandomStringElement([]string{"8", "9", "a", "A", "b", "B"})
	uuid += u.randomAFNumberCount(3)
	uuid += "-"
	uuid += u.randomAFNumberCount(12)
	return
}
