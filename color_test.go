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
	Expect(t, true, len(color) >= 5)
	Expect(t, true, len(color) <= 11)
	Expect(t, true, strings.Contains(color, ","))
}

func TestRGBAsArray(t *testing.T) {
	c := New().Color()

	Expect(t, 3, len(c.RGBAsArray()))
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

func TestRGBA(t *testing.T) {
	c := New().Color()

	color := c.RGBA()

	Expect(t, true, len(color) > 10)
	Expect(t, true, len(color) < 20)
	Expect(t, true, strings.Contains(color, ","))
	Expect(t, true, strings.Contains(color, "%"))
}

func TestOKLCH(t *testing.T) {
	c := New().Color()

	color := c.OKLCH()

	Expect(t, true, len(color) > 11)
	Expect(t, true, len(color) < 19)
	Expect(t, true, strings.Contains(color, "%"))
	Expect(t, true, strings.Contains(color, "deg"))
}

func TestRGBAAsArray(t *testing.T) {
	c := New().Color()

	color := c.RGBAAsArray()

	Expect(t, 4, len(color))
	Expect(t, true, strings.Contains(color[3], "%"))
}

func TestOKLCHAsArray(t *testing.T) {
	c := New().Color()

	Expect(t, 3, len(c.OKLCHAsArray()))
}

func TestCSSOKLCH(t *testing.T) {
	c := New().Color()

	color := c.CSSOKLCH()

	Expect(t, true, len(color) > 17)
	Expect(t, true, len(color) <= 26)
	Expect(t, true, strings.Contains(color, "oklch("))
	Expect(t, true, strings.Contains(color, ","))
	Expect(t, true, strings.Contains(color, ")"))
}

func TestCSSRGBA(t *testing.T) {
	c := New().Color()

	color := c.CSSRGBA()

	Expect(t, true, len(color) > 14)
	Expect(t, true, len(color) <= 25)
	Expect(t, true, strings.Contains(color, "rgba("))
	Expect(t, true, strings.Contains(color, ","))
	Expect(t, true, strings.Contains(color, ")"))
}
