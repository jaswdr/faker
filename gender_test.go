package faker

import (
	"slices"
	"testing"
)

func TestGenderName(t *testing.T) {
	v := New().Gender().Name()
	Expect(t, true, slices.Contains([]string{"masculine", "feminine"}, v))
}

func TestGenderAbbr(t *testing.T) {
	v := New().Gender().Abbr()
	Expect(t, true, slices.Contains([]string{"masc", "fem"}, v))
}
