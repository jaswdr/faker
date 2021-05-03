package faker

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const PROFILE_IMAGE_BASE_URL = "https://thispersondoesnotexist.com/image"

// ProfileImage  is a faker struct for ProfileImage
type ProfileImage struct {
	faker *Faker
}

// Image generates a *os.File with a random profile image using the thispersondoesnotexist.com service
func (pi ProfileImage) Image() *os.File {
	resp, err := http.Get(PROFILE_IMAGE_BASE_URL)
	if err != nil {
		log.Println("Error while requesting", PROFILE_IMAGE_BASE_URL, ":", err)
	}

	defer resp.Body.Close()

	f, err := ioutil.TempFile(os.TempDir(), "profil-picture-img-*.jfif")
	if err != nil {
		panic(err)
	}

	io.Copy(f, resp.Body)
	return f
}
