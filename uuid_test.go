package faker

import (
	"regexp"
	"testing"
)

func TestUUIDv4(t *testing.T) {
	f := New()
	value := f.UUID().V4()
	match, err := regexp.MatchString("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$", value)
	Expect(t, true, err == nil)
	Expect(t, true, match)
}

func TestUUIDV4UniqueInSequence(t *testing.T) {
	f := New()
	before := f.UUID().V4()
	for i := 0; i < 100; i++ {
		after := f.UUID().V4()
		Expect(t, true, before != after)
		before = after
	}
}
