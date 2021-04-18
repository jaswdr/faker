package faker

import (
	"net/http"
	"os"
)

const BASE_URL = "https://loremflickr.com"

type LoremFlickr struct {
	faker *Faker
}

func (lf LoremFlickr) Image(width, height int, categories []*string, prefix *string, categoriesStrict *bool) *os.File {

	url := BASE_URL

	switch *prefix {
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

	url += string('/') + string(width) + string('/') + string(height)

	if len(categories) > 0 {

		url += string('/')

		for _, category := range categories {
			url += *category + string(',')
		}

		if *categoriesStrict {
			url += "/all"
		}
	}

	http.Get(url)

	return &os.File{}
}
