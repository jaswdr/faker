package faker

// Gender is a faker struct for Gender
type Blood struct {
	Faker *Faker
}

// Name returns a Gender name for Gender
func (f Blood) Name() string {
	return f.Faker.RandomStringElement([]string{"A+", "A-", "B+", "B-", "AB+", "AB-", "O+", "O-"})
}

// Abbr returns a Gender name for Gender
func (f Blood) Abbr() string {
	return f.Faker.RandomStringElement([]string{"A+", "A-", "B+", "B-", "AB+", "AB-", "O+", "O-"})
}
