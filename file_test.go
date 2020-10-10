package faker

import (
	"strings"
	"testing"
)

func TestExtension(t *testing.T) {
	p := New().File()
	Expect(t, true, len(p.Extension()) > 2)
}

func TestFileWithExtension(t *testing.T) {
	p := New().File()
	Expect(t, true, len(strings.Split(p.FilenameWithExtension(), ".")) == 2)
}
