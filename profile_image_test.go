package faker

import (
	"fmt"
	"strings"
	"testing"
)

func TestProfileImage(t *testing.T) {
	f := New()
	value, err := f.ProfileImage().Image()
	Expect(t, nil, err)
	Expect(t, fmt.Sprintf("%T", value), "*os.File")
	Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())
}

func TestProfileImageErrorIfRequestFails(t *testing.T) {
	f := New()
	service := f.ProfileImage()
	expected := fmt.Errorf("request failed")
	service.HTTPClient = ErrorRaiserHTTPClient{err: expected}
	_, err := service.Image()
	Expect(t, expected, err)
}

func TestProfileImageErrorIfTempFileCreationFails(t *testing.T) {
	f := New()
	service := f.ProfileImage()
	expected := fmt.Errorf("temp file creation failed")
	service.TempFileCreator = ErrorRaiserTempFileCreator{err: expected}
	_, err := service.Image()
	Expect(t, expected, err)
}
