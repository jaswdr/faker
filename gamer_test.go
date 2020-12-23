package faker

import (
	"testing"
)

func TestTag(t *testing.T) {
	v := New().Gamer().Tag()
	NotExpect(t, "", v)
}