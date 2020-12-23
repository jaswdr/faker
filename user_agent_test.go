package faker

import (
	"testing"
)

func TestInternetExplorer(t *testing.T) {
	u := New().UserAgent()
	NotExpect(t, "", u.InternetExplorer())
}

func TestOpera(t *testing.T) {
	u := New().UserAgent()
	NotExpect(t, "", u.Opera())
}

func TestSafari(t *testing.T) {
	u := New().UserAgent()
	NotExpect(t, "", u.Safari())
}

func TestFirefox(t *testing.T) {
	u := New().UserAgent()
	NotExpect(t, "", u.Firefox())
}

func TestChrome(t *testing.T) {
	u := New().UserAgent()
	NotExpect(t, "", u.Chrome())
}

func TestUserAgent(t *testing.T) {
	u := New().UserAgent()
	Expect(t, true, len(u.UserAgent()) > 0)
}
