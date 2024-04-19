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
