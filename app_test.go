package faker

import (
	"testing"
)

func TestAppName(t *testing.T) {
	a := New().App()
	NotExpect(t, "", a.Name())
}

func TestAppVersion(t *testing.T) {
	a := New().App()
	NotExpect(t, "", a.Version())
}