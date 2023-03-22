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

// Image generates a *os.File with a random profile image using the thispersondoesnotexist.com service
func (pi ProfileImage) Image() *os.File {
	gender := pi.faker.RandomStringElement([]string{"women", "men"})
	profileId := pi.faker.IntBetween(1, 99)
	url := fmt.Sprintf("%s/%s/%s/%d.jpg", profileImageBaseURL, portraitsEndpoint, gender, profileId)
	resp, err := pi.HTTPClient.Get(url)
	if err != nil {
		log.Println("Error while requesting", profileImageBaseURL, ":", err)
		panic(err)
	}

	f, err := pi.TempFileCreator.TempFile("profile-picture-img-*.jpg")
	if err != nil {
		log.Println("Error while creating a temp file:", err)
		panic(err)
	}

	io.Copy(f, resp.Body)
	resp.Body.Close()
	return f
}
