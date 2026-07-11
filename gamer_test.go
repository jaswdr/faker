package faker

import (
	"slices"
	"testing"
)

func TestTag(t *testing.T) {
	v := New().Gamer().Tag()
	Expect(t, true, slices.Contains(gamerTags, v))
}
