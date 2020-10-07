package faker

import (
	"strings"
	"testing"
)

func TestGenerateVideoID(t *testing.T) {
	y := New().YouTube()
	Expect(t, VideoLength, len(y.GenerateVideoID()))
}

func TestGenerateFullURL(t *testing.T) {
	y := New().YouTube()
	split := strings.Split(y.GenerateFullURL(), "v=")
	Expect(t, "www.youtube.com/watch?", split[0])
	Expect(t, VideoLength, len(split[1]))
}

func TestGenerateShareURL(t *testing.T) {
	y := New().YouTube()
	split := strings.Split(y.GenerateShareURL(), "/")
	Expect(t, "youtu.be", split[0])
	Expect(t, VideoLength, len(split[1]))
}

func TestGenerateEmbededURL(t *testing.T) {
	y := New().YouTube()
	split := strings.Split(y.GenerateEmbededURL(), "embed/")
	Expect(t, "www.youtube.com/", split[0])
	Expect(t, VideoLength, len(split[1]))
}
