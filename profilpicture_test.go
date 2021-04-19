package faker

import (
	"fmt"
	"strings"
	"testing"
)

func TestImageProfil(t *testing.T) {

	f := New()
	value := f.ProfilPicture().Image()
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".jfif"), true, value.Name())
}
