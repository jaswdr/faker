package faker

import (
	"testing"
)

func TestGenreName(t *testing.T) {
	v := New().Genre().Name()
	NotExpect(t, "", v)
}

func TestGenreNameWithDescription(t *testing.T) {
	name, description := New().Genre().NameWithDescription()
	NotExpect(t, "", name)
	NotExpect(t, "", description)
}
