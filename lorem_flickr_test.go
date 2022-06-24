package faker

import (
	"fmt"
	"strings"
	"testing"
)

func TestLoremFlickrImage(t *testing.T) {
	f := New()
	value := f.LoremFlickr().Image(300, 200, []string{}, "", false)
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())
}

func TestLoremFlickrImageWithPrefix(t *testing.T) {
	f := New()
	for _, prefix := range []string{"g", "p", "red", "green", "blue"} {
		value := f.LoremFlickr().Image(300, 200, []string{}, prefix, false)
		Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())
	}
}

func TestLoremFlickrImageWithCategories(t *testing.T) {
	f := New()
	value := f.LoremFlickr().Image(300, 200, []string{"cat", "dog"}, "", false)
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())

	value = f.LoremFlickr().Image(300, 200, []string{"cat", "dog"}, "", true)
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())
}

func TestLoremFlickrImagePanicIfRequestFails(t *testing.T) {
	f := New()
	service := f.LoremFlickr()
	expected := fmt.Errorf("request failed")
	service.HTTPClient = ErrorRaiserHTTPClient{err: expected}
	defer func() {
		Expect(t, recover(), expected)
	}()
	service.Image(300, 200, []string{}, "", false)
}

func TestLoremFlickrImagePanicIfTempFileCreationFails(t *testing.T) {
	f := New()
	service := f.LoremFlickr()
	expected := fmt.Errorf("temp file creation failed")
	service.TempFileCreator = ErrorRaiserTempFileCreator{err: expected}
	defer func() {
		Expect(t, recover(), expected)
	}()
	service.Image(300, 200, []string{}, "", false)
}
