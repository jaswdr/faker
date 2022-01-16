package faker

import (
	"testing"
)

func TestMusicName(t *testing.T) {
	m := New().Music()
	Expect(t, true, len(m.Name()) > 0)
}

func TestMusicAuthor(t *testing.T) {
	m := New().Music()
	Expect(t, true, len(m.Author()) > 0)
}

func TestMusicGenre(t *testing.T) {
	m := New().Music()
	Expect(t, true, len(m.Genre()) > 0)
}

func TestMusicLength(t *testing.T) {
	m := New().Music()
	Expect(t, true, 2 < m.Length().Minutes() && m.Length().Minutes() < 9)
}
