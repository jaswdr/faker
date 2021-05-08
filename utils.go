package faker

import (
	"net/http"
)

const (
	maxRetries = 7
)

func get(url string) (resp *http.Response, err error) {
	for i := 0; i < maxRetries; i++ {
		resp, err = http.Get(url)
		if err == nil {
			return resp, err
		}
	}

	return resp, err
}
