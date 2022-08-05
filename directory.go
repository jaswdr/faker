package faker

import (
	"path/filepath"
	"runtime"
	"strings"
)

// Directory is a faker struct for Directory
type Directory struct {
	Faker *Faker
}

// Directory returns a fake directory path (the directory path style is dependent OS dependent)
func (d Directory) Directory(levels int) string {
	prefix := "/"

	// This will only be true on Windows, coverage may be impacted depending
	// on host OS
	if runtime.GOOS == "windows" {
		prefix = d.DriveLetter()
	}

	return prefix + filepath.Join(d.Faker.Lorem().Words(levels)...)
}

// UnixDirectory returns a fake Unix directory path, regardless of the host OS
func (d Directory) UnixDirectory(levels int) string {
	return "/" + strings.Join(d.Faker.Lorem().Words(levels), "/")
}

// WindowsDirectory returns a fake Windows directory path, regardless of the host OS
func (d Directory) WindowsDirectory(levels int) string {
	return d.DriveLetter() + strings.Join(d.Faker.Lorem().Words(levels), "\\")
}

// DriveLetter returns a fake Win32 drive letter
func (d Directory) DriveLetter() string {
	return d.Faker.RandomLetter() + ":\\"
}
