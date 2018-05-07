package faker

import (
	"strings"
	"testing"
)

func TestHex(t *testing.T) {
	c := New().Color()

	color := c.Hex()
	Expect(t, 7, len(color))
	Expect(t, true, strings.Contains(color, "#"))
}

func TestRGB(t *testing.T) {
	c := New().Color()

	color := c.RGB()
	Expect(t, true, len(color) > 6)
	Expect(t, true, len(color) <= 11)
	Expect(t, true, strings.Contains(color, ","))
}

func TestRGBAsArray(t *testing.T) {
	c := New().Color()

	color := c.RGBAsArray()
	Expect(t, 3, len(color))
}

func TestCSS(t *testing.T) {
	c := New().Color()

	color := c.CSS()
	Expect(t, true, len(color) > 10)
	Expect(t, true, len(color) <= 16)
	Expect(t, true, strings.Contains(color, "rgb("))
	Expect(t, true, strings.Contains(color, ","))
	Expect(t, true, strings.Contains(color, ")"))
}

func TestSafeColorName(t *testing.T) {
	c := New().Color()

	color := c.SafeColorName()
	Expect(t, true, len(color) > 0)
}

func TestColorName(t *testing.T) {
	c := New().Color()

	color := c.ColorName()
	Expect(t, true, len(color) > 0)
}
