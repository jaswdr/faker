package faker

type CarPlate struct {
	Faker *Faker
}

var (
	regions = []string{
		"АK", "АB", "АC", "AЕ", "AН", "АM", "АO", "АP", "АT",
		"AА", "АI", "BА", "ВB", "BС", "ВE", "BН", "ВI", "BК",
		"СH", "ВM", "ВO", "АX", "ВT", "ВX", "CА", "CВ", "СE",
	}

	series = []string{
		"KР", "ВI", "ВO", "АA", "EА", "BА", "PЕ", "НA", "IB",
		"KА", "KK", "OМ", "АM", "TА", "HI", "ОA", "CК", "PВ",
		"КC", "CА", "TЕ", "XА", "XО", "XМ", "MА", "МK", "MО",
	}
)

func (c CarPlate) Region() string {
	return c.Faker.RandomStringElement(regions)
}

func (c CarPlate) Code() int {
	return c.Faker.RandomNumber(4)
}

func (c CarPlate) Series() string {
	return c.Faker.RandomStringElement(series)
}
