package faker

import (
	"strings"
	"testing"
)

func TestGenerateVideoID(t *testing.T) {
	y := New().YouTube()
	got := len(y.GenerateVideoID())
	expect := VideoLength

	if got != expect {
		t.Errorf("expected length to be: %d, got %d", expect, got)
	}
}

func TestGenerateFullURL(t *testing.T) {
	y := New().YouTube()
	got := y.GenerateFullURL()
	parts := strings.Split(got, "v=")

	if parts[0] != "www.youtube.com/watch?" || len(parts[1]) != VideoLength {
		t.Errorf("received unexpected format for a full URL: %s", got)
	}
}

func TestGenerateShareURL(t *testing.T) {
	y := New().YouTube()
	got := y.GenerateShareURL()
	parts := strings.Split(got, "/")

	if parts[0] != "youtu.be" || len(parts[1]) != VideoLength {
		t.Errorf("received unexpected format for a share URL: %s", got)
	}
}

func TestGenerateEmbededURL(t *testing.T) {
	y := New().YouTube()
	got := y.GenerateEmbededURL()
	parts := strings.Split(got, "embed/")

	if parts[0] != "www.youtube.com/" || len(parts[1]) != VideoLength {
		t.Errorf("received unexpected format for a embeded URL: %s", got)
	}
}
