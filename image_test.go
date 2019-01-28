package faker

import (
	"fmt"
	"strings"
	"testing"
)

func TestImage(t *testing.T) {
	f := New()
	value := f.Image().Image(100, 100)
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".png"), true)
}
