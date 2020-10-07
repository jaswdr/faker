package faker

import (
	"testing"
)

func TestMimeType(t *testing.T) {
	p := New().MimeType()
	Expect(t, true, p.MimeType() != "")
}
