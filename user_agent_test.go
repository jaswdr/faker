package faker

import (
	"testing"
)

func TestInternetExplorer(t *testing.T) {
	u := New().UserAgent()
	Expect(t, "Mozilla/5.0 (compatible; MSIE 7.0; Windows 98; Win 9x 4.90; Trident/3.0)", u.InternetExplorer())
}

func TestOpera(t *testing.T) {
	u := New().UserAgent()
	Expect(t, "Opera/8.25 (Windows NT 5.1; en-US) Presto/2.9.188 Version/10.00", u.Opera())
}

func TestSafari(t *testing.T) {
	u := New().UserAgent()
	Expect(t, "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_7_1 rv:3.0; en-US) AppleWebKit/534.11.3 (KHTML, like Gecko) Version/4.0 Safari/534.11.3", u.Safari())
}

func TestFirefox(t *testing.T) {
	u := New().UserAgent()
	Expect(t, "Mozilla/5.0 (X11; Linuxi686; rv:7.0) Gecko/20101231 Firefox/3.6", u.Firefox())
}

func TestChrome(t *testing.T) {
	u := New().UserAgent()
	Expect(t, "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_6_5) AppleWebKit/5312 (KHTML, like Gecko) Chrome/14.0.894.0 Safari/5312", u.Chrome())
}

func TestUserAgent(t *testing.T) {
	u := New().UserAgent()
	Expect(t, true, len(u.UserAgent()) > 0)
}
