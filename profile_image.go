package faker

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	profileImageBaseURL = "https://randomuser.me"
	portraitsEndpoint   = "api/portraits"
)

// ProfileImage  is a faker struct for ProfileImage
type ProfileImage struct {
	faker           *Faker
	HTTPClient      HTTPClient
	TempFileCreator TempFileCreator
}

// Image generates a *os.File with a random profile image using the randomuser.me service.
// Panics on failure to maintain backward compatibility.
func (pi ProfileImage) Image() *os.File {
	f, err := pi.ImageFile()
	if err != nil {
		panic(err)
	}
	return f
}

// ImageFile generates a *os.File with a random profile image using the randomuser.me service.
func (pi ProfileImage) ImageFile() (*os.File, error) {
	gender := pi.faker.RandomStringElement([]string{"women", "men"})
	profileId := pi.faker.IntBetween(1, 99)
	url := fmt.Sprintf("%s/%s/%s/%d.jpg", profileImageBaseURL, portraitsEndpoint, gender, profileId)
	resp, err := pi.HTTPClient.Get(url)
	if err != nil {
		log.Println("Error while requesting", profileImageBaseURL, ":", err)
		return nil, err
	}
	defer resp.Body.Close()

	f, err := pi.TempFileCreator.TempFile("profile-picture-img-*.jpg")
	if err != nil {
		log.Println("Error while creating a temp file:", err)
		return nil, err
	}

	if _, err = io.Copy(f, resp.Body); err != nil {
		f.Close()
		return nil, err
	}

	return f, nil
}
