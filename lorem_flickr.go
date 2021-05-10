package faker

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const loremFlickrBaseURL = "https://loremflickr.com"

// LoremFlickr is a faker struct for LoremFlickr
type LoremFlickr struct {
	faker *Faker
}

// Image generates a *os.File with a random image using the loremflickr.com service
func (lf LoremFlickr) Image(width, height int, categories []string, prefix string, categoriesStrict bool) *os.File {

	url := loremFlickrBaseURL

	switch prefix {
	case "g":
		url += "/g"
	case "p":
		url += "/p"
	case "red":
		url += "/red"
	case "green":
		url += "/green"
	case "blue":
		url += "/blue"
	}

	url += string('/') + strconv.Itoa(width) + string('/') + strconv.Itoa(height)

	if len(categories) > 0 {

		url += string('/')

		for _, category := range categories {
			url += category + string(',')
		}

		if categoriesStrict {
			url += "/all"
		}
	}

	resp, err := get(url)
	defer resp.Body.Close()

	if err != nil {
		log.Println("Error while requesting", url, ":", err)
	}

	f, err := ioutil.TempFile(os.TempDir(), "loremflickr-img-*.jpg")
	if err != nil {
		panic(err)
	}

	io.Copy(f, resp.Body)

	return f
}
