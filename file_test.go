package faker

import (
	"os"
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
	parts := strings.Split(path, string(os.PathSeparator))
	Expect(t, true, isUnixPath(path) || isWindowsPath(path))
	Expect(t, true, len(parts) == 4)
	Expect(t, true, len(strings.Split(parts[len(parts)-1], ".")) == 2)

	p.OSResolver = WindowsOSResolver{}
	path = p.AbsoluteFilePath(2)
	Expect(t, true, isWindowsPath(p.AbsoluteFilePath(2)))
	Expect(t, true, len(parts) == 4)
	Expect(t, true, len(strings.Split(parts[len(parts)-1], ".")) == 2)
}

func TestAbsoluteFilePathForUnix(t *testing.T) {
	p := New().File()

	path := p.AbsoluteFilePathForUnix(2)
	parts := strings.Split(path, "/")

	Expect(t, true, isUnixPath(path))
	Expect(t, true, len(parts) == 4)
	Expect(t, true, len(strings.Split(parts[len(parts)-1], ".")) == 2)
}

func TestAbsoluteFilePathForWindows(t *testing.T) {
	p := New().File()

	path := p.AbsoluteFilePathForWindows(2)
	parts := strings.Split(path, "\\")

	Expect(t, true, isWindowsPath(path))
	Expect(t, true, len(parts) == 4)
	Expect(t, true, len(strings.Split(parts[len(parts)-1], ".")) == 2)
}
