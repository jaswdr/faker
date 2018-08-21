package faker

import (
	"regexp"
	"testing"
)

func TestUUIDv4(t *testing.T) {
	f := New()
	value := f.UUID().V4()
	match, err := regexp.MatchString("^[a-fA-F0-9]{8}[a-fA-F0-9]{4}4[a-fA-F0-9]{3}[8|9|aA|bB][a-fA-F0-9]{3}[a-fA-F0-9]{12}$", value)
	t.Log(value)
	Expect(t, true, err == nil)
	Expect(t, true, match)
}

func TestUUIDV4UniqueInSequence(t *testing.T) {
	f := New()
	last := f.UUID().V4()
	current := f.UUID().V4()
	Expect(t, true, last != current)
}
