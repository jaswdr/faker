package faker

import (
	"slices"
	"strings"
	"testing"
)

func TestMimeType(t *testing.T) {
	v := New().MimeType().MimeType()
	Expect(t, true, slices.Contains(mimeType, v))
	Expect(t, true, strings.Contains(v, "/"))
}
