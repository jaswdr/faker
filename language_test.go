package faker

import (
	"testing"
)

func TestLanguage(t *testing.T) {
	v := New().Language().Language()
	NotExpect(t, "", v)
}