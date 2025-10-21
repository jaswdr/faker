package faker

import (
	"strconv"
	"strings"
)

var (
	colorLetters = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

	safeColorNames = []string{
		"black", "maroon", "green", "navy", "olive",
		"purple", "teal", "lime", "blue", "silver",
		"gray", "yellow", "fuchsia", "aqua", "white",
	}

	allColorNames = []string{
		"AliceBlue", "AntiqueWhite", "Aqua", "Aquamarine",
		"Azure", "Beige", "Bisque", "Black", "BlanchedAlmond",
		"Blue", "BlueViolet", "Brown", "BurlyWood", "CadetBlue",
		"Chartreuse", "Chocolate", "Coral", "CornflowerBlue",
		"Cornsilk", "Crimson", "Cyan", "DarkBlue", "DarkCyan",
		"DarkGoldenRod", "DarkGray", "DarkGreen", "DarkKhaki",
		"DarkMagenta", "DarkOliveGreen", "Darkorange", "DarkOrchid",
		"DarkRed", "DarkSalmon", "DarkSeaGreen", "DarkSlateBlue",
		"DarkSlateGray", "DarkTurquoise", "DarkViolet", "DeepPink",
		"DeepSkyBlue", "DimGray", "DimGrey", "DodgerBlue", "FireBrick",
		"FloralWhite", "ForestGreen", "Fuchsia", "Gainsboro", "GhostWhite",
		"Gold", "GoldenRod", "Gray", "Green", "GreenYellow", "HoneyDew",
		"HotPink", "IndianRed", "Indigo", "Ivory", "Khaki", "Lavender",
		"LavenderBlush", "LawnGreen", "LemonChiffon", "LightBlue", "LightCoral",
		"LightCyan", "LightGoldenRodYellow", "LightGray", "LightGreen", "LightPink",
		"LightSalmon", "LightSeaGreen", "LightSkyBlue", "LightSlateGray", "LightSteelBlue",
		"LightYellow", "Lime", "LimeGreen", "Linen", "Magenta", "Maroon", "MediumAquaMarine",
		"MediumBlue", "MediumOrchid", "MediumPurple", "MediumSeaGreen", "MediumSlateBlue",
		"MediumSpringGreen", "MediumTurquoise", "MediumVioletRed", "MidnightBlue",
		"MintCream", "MistyRose", "Moccasin", "NavajoWhite", "Navy", "OldLace", "Olive",
		"OliveDrab", "Orange", "OrangeRed", "Orchid", "PaleGoldenRod", "PaleGreen",
		"PaleTurquoise", "PaleVioletRed", "PapayaWhip", "PeachPuff", "Peru", "Pink", "Plum",
		"PowderBlue", "Purple", "Red", "RosyBrown", "RoyalBlue", "SaddleBrown", "Salmon",
		"SandyBrown", "SeaGreen", "SeaShell", "Sienna", "Silver", "SkyBlue", "SlateBlue",
		"SlateGray", "Snow", "SpringGreen", "SteelBlue", "Tan", "Teal", "Thistle", "Tomato",
		"Turquoise", "Violet", "Wheat", "White", "WhiteSmoke", "Yellow", "YellowGreen",
	}
)

// Color is a faker struct for Color
type Color struct {
	Faker *Faker
}

// Hex returns a fake hex for Color
func (c Color) Hex() string {
	color := "#"

	for i := 0; i < 6; i++ {
		color = color + c.Faker.RandomStringElement(colorLetters)
	}

	return color
}

// RGB returns a fake rgb for Color
func (c Color) RGB() string {
	color := strconv.Itoa(c.Faker.IntBetween(0, 255))

	for i := 0; i < 2; i++ {
		color = color + "," + strconv.Itoa(c.Faker.IntBetween(0, 255))
	}

	return color
}

// RGBA returns a fake color in rgba format for Color
func (c Color) RGBA() string {
	return c.RGB() + ", " + strconv.Itoa(c.Faker.IntBetween(0, 100)) + "%"
}

// OKLCH returns a fake color in OKLCH format for Color
func (c Color) OKLCH() string {
	oklch := ""

	// lightness
	oklch = oklch + strconv.Itoa(c.Faker.IntBetween(0, 100)) + "%, "

	// chroma
	oklch = oklch + strconv.Itoa(c.Faker.IntBetween(0, 100)) + "%, "

	// hue
	oklch = oklch + strconv.Itoa(c.Faker.IntBetween(0, 360)) + "deg"

	return oklch
}

// RGBAsArray returns a fake rgb color in array format for Color
func (c Color) RGBAsArray() [3]string {
	split := strings.Split(c.RGB(), ",")
	return [3]string{split[0], split[1], split[2]}
}

// RGBAAsArray returns a fake rgba color in array format for Color
func (c Color) RGBAAsArray() [4]string {
	split := strings.Split(c.RGBA(), ",")
	return [4]string{split[0], split[1], split[2], split[3]}
}

// OKLCHAsArray returns a fake OKLCH color in array format for Color
func (c Color) OKLCHAsArray() [3]string {
	split := strings.Split(c.OKLCH(), ",")
	return [3]string{split[0], split[1], split[2]}
}

// CSS returns a fake color in CSS format for Color
func (c Color) CSS() string {
	return "rgb(" + c.RGB() + ")"
}

// CSSRGBA returns a fake color in CSS rgba format for Color
func (c Color) CSSRGBA() string {
	return "rgba(" + c.RGBA() + ")"
}

// CSSOKLCH returns a fake color in CSS oklch format for Color
func (c Color) CSSOKLCH() string {
	return "oklch(" + c.OKLCH() + ")"
}

// SafeColorName returns a fake safe color name for Color
func (c Color) SafeColorName() string {
	return c.Faker.RandomStringElement(safeColorNames)
}

// ColorName returns a fake color name for Color
func (c Color) ColorName() string {
	return c.Faker.RandomStringElement(allColorNames)
}
