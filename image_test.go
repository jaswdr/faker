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

func (creator ErrorRaiserPngEncoder) Encode(w io.Writer, m image.Image) error {
	return creator.err
}

func TestImage(t *testing.T) {
	f := New()
	value := f.Image().Image(100, 100)
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".png"), true, value.Name())
}

func TestImagePanicIfTempFileCreationFails(t *testing.T) {
	f := New()
	img := f.Image()
	expected := fmt.Errorf("temp file creation failed")
	img.TempFileCreator = ErrorRaiserTempFileCreator{err: expected}
	defer func() {
		Expect(t, recover(), expected)
	}()
	img.Image(100, 100)
}

func TestImagePanicIfEncodingFails(t *testing.T) {
	f := New()
	img := f.Image()
	expected := fmt.Errorf("png encoding failed")
	img.PngEncoder = ErrorRaiserPngEncoder{err: expected}
	defer func() {
		Expect(t, recover(), expected)
	}()
	img.Image(100, 100)
}
