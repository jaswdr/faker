package faker

import (
	"fmt"
	"strings"
	"testing"
)

func TestImageLorem(t *testing.T) {

	f := New()
	value := f.LoremFlickr().Image(100, 100, nil, nil, nil)
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())
}
