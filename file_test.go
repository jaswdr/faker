package faker

import (
	"path/filepath"
	"regexp"
	"runtime"
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

func TestAbsolutePath(t *testing.T) {
	p := New().File()

	path := p.AbsoluteFilePath(2)
	Expect(t, true, filepath.IsAbs(path))

	if runtime.GOOS == "windows" {
		Expect(t, true, regexp.MustCompile(`^[a-zA-Z]:`).MatchString(path[:2]))
	}
}

func TestAbsoluteFilePathForUnix(t *testing.T) {
	p := New().File()

	path := p.AbsoluteFilePathForUnix(2)
	parts := strings.Split(path, "/")

	Expect(t, true, path[0] == '/')
	Expect(t, true, len(parts) == 4)
	Expect(t, true, len(strings.Split(parts[len(parts)-1], ".")) == 2)
}

func TestAbsoluteFilePathForWindows(t *testing.T) {
	p := New().File()

	path := p.AbsoluteFilePathForWindows(2)
	parts := strings.Split(path, "\\")

	Expect(t, true, regexp.MustCompile(`^[a-zA-Z]:`).MatchString(parts[0]))
	Expect(t, true, len(parts) == 4)
	Expect(t, true, len(strings.Split(parts[len(parts)-1], ".")) == 2)
}
