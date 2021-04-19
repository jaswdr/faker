package faker

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const URL = "https://thispersondoesnotexist.com/image"

type ProfilPicture struct {
	faker *Faker
}

func (pp ProfilPicture) Image() *os.File {

	resp, err := http.Get(URL)
	if err != nil {
		log.Println("Error while requesting", URL, ":", err)
	}
	defer resp.Body.Close()

	f, err := ioutil.TempFile(os.TempDir(), "profil-picture-img-*.jfif")
	if err != nil {
		panic(err)
	}

	io.Copy(f, resp.Body)

	return f
}
