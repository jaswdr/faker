package faker

import (
	"fmt"
	"strings"
	"testing"
)

func TestProfileImage(t *testing.T) {
	f := New()
	value := f.ProfileImage().Image()
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".jfif"), true, value.Name())
}
