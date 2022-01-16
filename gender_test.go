package faker

import (
	"testing"
)

func TestGenderName(t *testing.T) {
	v := New().Gender().Name()
	NotExpect(t, "", v)
	Expect(t, true, v == "masculine" || v == "feminine")
}

func TestGenderAbbr(t *testing.T) {
	v := New().Gender().Abbr()
	NotExpect(t, "", v)
	Expect(t, true, v == "masc" || v == "fem")
}
