package faker

import (
	"strings"
	"testing"
)

func TestExtension(t *testing.T) {
	p := New().File()
	Expect(t, true, p.Extension() != "")
}

func TestFileWithExtension(t *testing.T) {
	p := New().File()
	Expect(t, true, len(strings.Split(p.FilenameWithExtension(), ".")) == 2)
}
