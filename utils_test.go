package faker

import (
	"net/http"
	"os"
	"runtime"
	"strings"
	"testing"
)

type ErrorRaiserHTTPClient struct {
	err error
}

func (client ErrorRaiserHTTPClient) Get(_ string) (*http.Response, error) {
	return nil, client.err
}

type ErrorRaiserTempFileCreator struct {
	err error
}

func (creator ErrorRaiserTempFileCreator) TempFile(_ string) (*os.File, error) {
	return nil, creator.err
}

func TestHTTPClientImplCanDoGetRequests(t *testing.T) {
	client := HTTPClientImpl{}
	resp, err := client.Get("https://www.example.com")
	Expect(t, err, nil)
	Expect(t, resp.StatusCode, http.StatusOK)
}

func TestHTTPClientImplReturnsErrorWhenRequestFails(t *testing.T) {
	client := HTTPClientImpl{}
	_, err := client.Get("https://invalid")
	NotExpect(t, err, nil)
}

func TestTempFileCreatorImplCanCreateTempFiles(t *testing.T) {
	creator := TempFileCreatorImpl{}
	f, err := creator.TempFile("prefix")
	Expect(t, err, nil)
	Expect(t, true, strings.Contains(f.Name(), "prefix"))
	Expect(t, f.Close(), nil)
}

type WindowsOSResolver struct{}

func (WindowsOSResolver) OS() string {
	return "windows"
}

func TestOSResolverImplReturnsGOOS(t *testing.T) {
	resolver := OSResolverImpl{}
	Expect(t, runtime.GOOS, resolver.OS())
}
