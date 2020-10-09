package faker

// Boolean is a faker struct for Boolean
type Boolean struct {
	Faker *Faker
}

var (
	intBool = []int{0, 1}
	boolean = []bool{true, false}
)

// Bool returns a fake bool for Faker
func (c Boolean) Bool() bool {
	return c.Faker.IntBetween(0, 100) > 50
}

// BoolWithChance returns true with a given percentual chance that the value is true, otherwise returns false
func (c Boolean) BoolWithChance(chanceTrue int) bool {
	if chanceTrue <= 0 {
		return false
	} else if chanceTrue >= 100 {
		return true
	}

	return c.Faker.IntBetween(0, 100) < chanceTrue
}

// IntBool returns a fake bool for Integer Boolean
func (c Boolean) IntBool() int {
	return c.Faker.RandomIntElement(intBool)
}

// StringBool returns a fake bool for StringBool Boolean
func (c Boolean) StringBool(firstArg string, secondArg string) string {
	boolean := []string{firstArg, secondArg}

	return c.Faker.RandomStringElement(boolean)
}
