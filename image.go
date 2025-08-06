package faker

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
)

// PngEncoder encodes a image.Image to a io.Writer
type PngEncoder interface {
	Encode(w io.Writer, m image.Image) error
}

// PngEncoderImpl is the default implementation of the PngEncoder
type PngEncoderImpl struct{}

// Encode does the encoding of the image.Image to an io.Writer
func (PngEncoderImpl) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

// Image is a faker struct for Image
type Image struct {
	faker           *Faker
	TempFileCreator TempFileCreator
	PngEncoder      PngEncoder
}

// Image returns a fake image file.
// Width and height must be positive values. Panics on file creation or encoding errors
// to maintain backward compatibility with existing API contract.
func (i Image) Image(width, height int) *os.File {
	// Input validation - use reasonable defaults instead of nil to avoid breaking changes
	if width <= 0 {
		width = 100
	}
	if height <= 0 {
		height = 100
	}

	// Reasonable limits to prevent excessive memory usage
	if width > 10000 {
		width = 10000
	}
	if height > 10000 {
		height = 10000
	}
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	black := color.RGBA{0, 0, 0, 0xff}
	white := color.RGBA{0xff, 0xff, 0xff, 0xff}
	step := 4
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if y > 0 {
				if x%step == 0 {
					if y%step == 0 {
						img.Set(x, y, black)
					} else {
						img.Set(x, y, white)
					}
				} else {
					img.Set(x, y, white)
				}
			} else {
				img.Set(x, y, white)
			}
		}
	}

	f, err := i.TempFileCreator.TempFile("fake-img-*.png")
	if err != nil {
		panic(err)
	}

	err = i.PngEncoder.Encode(f, img)
	if err != nil {
		panic(err)
	}

	return f
}
