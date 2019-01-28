package faker

import (
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
)

type Image struct {
	faker *Faker
}

func (i Image) Image(width, height int) *os.File {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	dot := color.RGBA{0, 0, 0, 0xff}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, dot)
		}
	}

	f, err := ioutil.TempFile(os.TempDir(), "fake-img-*.png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}

	return f
}
