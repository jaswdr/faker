package faker

import "testing"

func TestRandomElement(t *testing.T) {
	f := New()
	m := []string{"str1", "str2"}
	randomStr := RandomElement(f, m...)
	Expect(t, true, randomStr == "str1" || randomStr == "str2")
}

func TestRandomElementWeighted(t *testing.T) {
	f := New()
	m := map[int]string{
		0: "zeroChance",
		1: "someChance",
		5: "moreChance",
	}

	for i := 0; i < 5; i++ {
		got := RandomElementWeighted(f, m)
		Expect(t, true, got == "someChance" || got == "moreChance")
		Expect(t, true, got != "zeroChance")
	}
}

func TestRandomMapKey(t *testing.T) {
	f := New()
	m := map[int]string{
		1:  "one",
		5:  "five",
		42: "forty two",
	}

	randomInt := RandomMapKey(f, m)
	Expect(t, true, randomInt == 1 || randomInt == 5 || randomInt == 42)
}

func TestRandomMapValue(t *testing.T) {
	f := New()
	m := map[int]string{
		1:  "one",
		5:  "five",
		42: "forty two",
	}

	randomStr := RandomMapValue(f, m)
	Expect(t, true, randomStr == "one" || randomStr == "five" || randomStr == "forty two")
}
