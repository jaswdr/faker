package faker

import (
	"regexp"
	"strings"
	"testing"
)

func isUnixPath(path string) bool {
	return path[0] == '/'
}

func isWindowsPath(path string) bool {
	return regexp.MustCompile(`^[a-zA-Z]:\\`).MatchString(path[:3])
}

func TestDirectory(t *testing.T) {
	p := New().Directory()
	dir := p.Directory(2)
	Expect(t, true, isUnixPath(dir) || isWindowsPath(dir))

	p.OSResolver = WindowsOSResolver{}
	Expect(t, true, isWindowsPath(p.Directory(2)))
}

func TestUnixDirectory(t *testing.T) {
	p := New().Directory()

	dir := p.UnixDirectory(2)
	parts := strings.Split(dir, "/")

	Expect(t, true, len(parts) == 3)
	Expect(t, true, isUnixPath(dir))
}

func TestWindowsDirectory(t *testing.T) {
	p := New().Directory()

	dir := p.WindowsDirectory(2)
	parts := strings.Split(dir, "\\")

	Expect(t, true, len(parts) == 3)
	Expect(t, true, isWindowsPath(dir))
}

func TestDriveLetter(t *testing.T) {
	p := New().Directory()
	Expect(t, true, isWindowsPath(p.DriveLetter()))
}
