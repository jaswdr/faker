package faker

import (
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"runtime"
)

const (
	maxRetries = maxRetriesDefault
)

// HTTPClient does HTTP requests to remote servers
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

// HTTPClientImpl is the default implementation of HTTPClient
type HTTPClientImpl struct{}

// Get do a GET request and returns a *http.Response
func (HTTPClientImpl) Get(url string) (resp *http.Response, err error) {
	for i := 0; i < maxRetries; i++ {
		resp, err = http.Get(url)
		if err == nil {
			return resp, nil
		}
	}

	return resp, err
}

// TempFileCreator creates temporary files
type TempFileCreator interface {
	TempFile(prefix string) (f *os.File, err error)
}

// TempFileCreatorImpl is the default implementation of TempFileCreator
type TempFileCreatorImpl struct{}

// TempFile creates a temporary file
func (TempFileCreatorImpl) TempFile(prefix string) (f *os.File, err error) {
	return os.CreateTemp(os.TempDir(), prefix)
}

// OSResolver returns the GOOS value for operating an operating system
type OSResolver interface {
	OS() string
}

// OSResolverImpl is the default implementation of OSResolver
type OSResolverImpl struct{}

// OS returns the runtime.GOOS value for the host operating system
func (OSResolverImpl) OS() string {
	return runtime.GOOS
}

// Shuffle shuffles the slice in place.
func Shuffle[T any](slice []T) []T {
	if len(slice) <= 1 {
		return slice
	}

	original := make([]T, len(slice))
	copy(original, slice)

	maxAttempts := 10
	for attempts := 0; attempts < maxAttempts; attempts++ {
		rand.Shuffle(len(slice), func(i, j int) {
			slice[i], slice[j] = slice[j], slice[i]
		})

		if !reflect.DeepEqual(original, slice) {
			break
		}

		if attempts >= 3 && len(slice) <= 3 {
			break
		}
	}

	return slice
}

// ShuffleWith shuffles the slice in place using the Faker's thread-safe random source.
func ShuffleWith[T any](f Faker, slice []T) []T {
	if len(slice) <= 1 {
		return slice
	}

	original := make([]T, len(slice))
	copy(original, slice)

	maxAttempts := 10
	for attempts := 0; attempts < maxAttempts; attempts++ {
		for i := len(slice) - 1; i > 0; i-- {
			j := f.IntBetween(0, i)
			slice[i], slice[j] = slice[j], slice[i]
		}

		if !reflect.DeepEqual(original, slice) {
			break
		}

		if attempts >= 3 && len(slice) <= 3 {
			break
		}
	}

	return slice
}
