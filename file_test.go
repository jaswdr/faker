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

func TestDirectory(t *testing.T) {
	p := New().File()
	Expect(t, true, len(strings.Split(p.Directory(2), string(filepath.Separator))) == 3)
}

func TestDriveLetter(t *testing.T) {
	p := New().File()
	reg := regexp.MustCompile(`^[a-zA-Z]:\\`)
	Expect(t, true, reg.MatchString(p.DriveLetter()))
}

func TestAbsolutePath(t *testing.T) {
	p := New().File()
	if runtime.GOOS == "windows" {
		Expect(t, true, filepath.IsAbs(p.AbsoluteWin32Path(3)))
	} else {
		Expect(t, true, filepath.IsAbs(p.AbsoluteUnixPath(3)))
	}
}
