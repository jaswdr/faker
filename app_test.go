package faker

import (
	"regexp"
	"slices"
	"testing"
)

func TestAppName(t *testing.T) {
	v := New().App().Name()
	Expect(t, true, slices.Contains(appNames, v))
}

func TestAppVersion(t *testing.T) {
	v := New().App().Version()
	re := regexp.MustCompile(`^v\d\.\d\.\d$`)
	Expect(t, true, re.MatchString(v))
}
