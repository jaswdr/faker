package faker

import (
	"os"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

func TestDirectory(t *testing.T) {
	p := New().Directory()

	dir := p.Directory(2)
	parts := strings.Split(dir, string(os.PathSeparator))

	Expect(t, true, len(parts) == 3)
	if runtime.GOOS == "windows" {
		Expect(t, true, regexp.MustCompile(`^[a-zA-Z]:`).MatchString(parts[0]))
	} else {
		Expect(t, true, dir[0] == '/')
	}
}

func TestUnixDirectory(t *testing.T) {
	p := New().Directory()

	dir := p.UnixDirectory(2)
	parts := strings.Split(dir, "/")

	Expect(t, true, len(parts) == 3)
	Expect(t, true, parts[0] == "")
}

func TestWindowsDirectory(t *testing.T) {
	p := New().Directory()

	dir := p.WindowsDirectory(2)
	parts := strings.Split(dir, "\\")

	Expect(t, true, len(parts) == 3)
	Expect(t, true, regexp.MustCompile(`^[a-zA-Z]:`).MatchString(parts[0]))
}

func TestDriveLetter(t *testing.T) {
	p := New().Directory()
	reg := regexp.MustCompile(`^[a-zA-Z]:\\`)
	Expect(t, true, reg.MatchString(p.DriveLetter()))
}
