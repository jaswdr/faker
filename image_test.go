package faker

import (
	"fmt"
	"image"
	"io"
	"strings"
	"testing"
)

type ErrorRaiserPngEncoder struct {
	err error
}

func (creator ErrorRaiserPngEncoder) Encode(_ io.Writer, _ image.Image) error {
	return creator.err
}

func TestImage(t *testing.T) {
	f := New()
	value, err := f.Image().Image(100, 100)
	Expect(t, nil, err)
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".png"), true, value.Name())
}

func TestImageErrorIfTempFileCreationFails(t *testing.T) {
	f := New()
	img := f.Image()
	expected := fmt.Errorf("temp file creation failed")
	img.TempFileCreator = ErrorRaiserTempFileCreator{err: expected}
	_, err := img.Image(100, 100)
	Expect(t, expected, err)
}

func TestImageErrorIfEncodingFails(t *testing.T) {
	f := New()
	img := f.Image()
	expected := fmt.Errorf("png encoding failed")
	img.PngEncoder = ErrorRaiserPngEncoder{err: expected}
	_, err := img.Image(100, 100)
	Expect(t, expected, err)
}
