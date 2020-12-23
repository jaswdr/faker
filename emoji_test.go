package faker

import (
	"testing"
)

func TestEmoji(t *testing.T) {
	e := New().Emoji()
	NotExpect(t, "", e.Emoji())
}

func TestEmojiCode(t *testing.T) {
	e := New().Emoji()
	NotExpect(t, "", e.EmojiCode())
}
