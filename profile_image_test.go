package faker

import (
	"fmt"
	"strings"
	"testing"
)

func TestProfileImage(t *testing.T) {
	f := New()
	value := f.ProfileImage().Image()
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())
}

func TestProfileImagePanicIfRequestFails(t *testing.T) {
	f := New()
	service := f.ProfileImage()
	expected := fmt.Errorf("request failed")
	service.HTTPClient = ErrorRaiserHTTPClient{err: expected}
	defer func() {
		Expect(t, recover(), expected)
	}()
	service.Image()
}

func TestProfileImagePanicIfTempFileCreationFails(t *testing.T) {
	f := New()
	service := f.ProfileImage()
	expected := fmt.Errorf("temp file creation failed")
	service.TempFileCreator = ErrorRaiserTempFileCreator{err: expected}
	defer func() {
		Expect(t, recover(), expected)
	}()
	service.Image()
}
