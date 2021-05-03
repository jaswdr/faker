package faker

import (
	"fmt"
	"strings"
	"testing"
)

func TestLoremFlickrImage(t *testing.T) {
	f := New()
	value := f.LoremFlickr().Image(300, 200, []string{}, "", false)
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())
}
