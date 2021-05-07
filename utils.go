package faker

import (
	"net/http"
)

const (
	max_retries = 7
)

func get(url string) (resp *http.Response, err error) {
	for i := 0; i < max_retries; i++ {
		resp, err = http.Get(url)
		if err == nil {
			return resp, err
		}
	}

	return resp, err
}
