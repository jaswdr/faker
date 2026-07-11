package faker

import (
	"slices"
	"testing"
)

func TestEmoji(t *testing.T) {
	v := New().Emoji().Emoji()
	Expect(t, true, slices.Contains(emojis, v))
}

func TestEmojiCode(t *testing.T) {
	v := New().Emoji().EmojiCode()
	Expect(t, true, slices.Contains(emojisCode, v))
}
