package faker

import (
	"fmt"
	"strings"
	"testing"
)

func TestLoremFlickrImage(t *testing.T) {
	f := New()
	value, err := f.LoremFlickr().Image(300, 200, []string{}, "", false)
	Expect(t, nil, err)
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())
}

func TestLoremFlickrImageWithPrefix(t *testing.T) {
	f := New()
	for _, prefix := range []string{"g", "p", "red", "green", "blue"} {
		value, err := f.LoremFlickr().Image(300, 200, []string{}, prefix, false)
		Expect(t, nil, err)
		Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())
	}
}

func TestLoremFlickrImageWithCategories(t *testing.T) {
	f := New()
	value, err := f.LoremFlickr().Image(300, 200, []string{"cat", "dog"}, "", false)
	Expect(t, nil, err)
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())

	value, err = f.LoremFlickr().Image(300, 200, []string{"cat", "dog"}, "", true)
	Expect(t, nil, err)
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())
}

func TestLoremFlickrImageErrorIfRequestFails(t *testing.T) {
	f := New()
	service := f.LoremFlickr()
	expected := fmt.Errorf("request failed")
	service.HTTPClient = ErrorRaiserHTTPClient{err: expected}
	_, err := service.Image(300, 200, []string{}, "", false)
	Expect(t, expected, err)
}

func TestLoremFlickrImageErrorIfTempFileCreationFails(t *testing.T) {
	f := New()
	service := f.LoremFlickr()
	expected := fmt.Errorf("temp file creation failed")
	service.TempFileCreator = ErrorRaiserTempFileCreator{err: expected}
	_, err := service.Image(300, 200, []string{}, "", false)
	Expect(t, expected, err)
}
