package faker

import (
	"testing"
)

func TestGenreName(t *testing.T) {
	v := New().Genre().Name()
	_, ok := genres[v]
	Expect(t, true, ok)
}

func TestGenreNameWithDescription(t *testing.T) {
	name, desc := New().Genre().NameWithDescription()
	Expect(t, true, len(name) > 0)
	Expect(t, true, len(desc) > 0)
	Expect(t, genres[name], desc)
}
