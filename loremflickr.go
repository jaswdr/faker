package faker

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const BASE_URL = "https://loremflickr.com"

type LoremFlickr struct {
	faker *Faker
}

func (lf LoremFlickr) Image(width, height int, categories []string, prefix string, categoriesStrict bool) *os.File {

	url := BASE_URL

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

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error while requesting", url, ":", err)
	}
	defer resp.Body.Close()

	f, err := ioutil.TempFile(os.TempDir(), "loremflickr-img-*.jpg")
	if err != nil {
		panic(err)
	}

	io.Copy(f, resp.Body)

	return f
}
