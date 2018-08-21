package faker

import (
	"strings"
	"testing"
)

func TestUser(t *testing.T) {
	i := New().Internet()

	user := i.User()
	Expect(t, true, len(user) > 0)
	Expect(t, false, strings.Contains(user, " "))
}

func TestDomain(t *testing.T) {
	i := New().Internet()

	domain := i.Domain()

	Expect(t, true, len(domain) > 0)
	Expect(t, true, strings.Index(domain, ".") > 0)

	split := strings.Split(domain, ".")
	Expect(t, 2, len(split))
}

func TestEmail(t *testing.T) {
	i := New().Internet()

	email := i.Email()
	split := strings.Split(email, "@")

	Expect(t, 2, len(split))
}

func TestFreeEmail(t *testing.T) {
	i := New().Internet()

	email := i.FreeEmail()
	split := strings.Split(email, "@")

	Expect(t, 2, len(split))
}

func TestSafeEmail(t *testing.T) {
	i := New().Internet()

	email := i.SafeEmail()
	split := strings.Split(email, "@")

	Expect(t, 2, len(split))
}

func TestCompanyEmail(t *testing.T) {
	i := New().Internet()

	email := i.CompanyEmail()
	split := strings.Split(email, "@")

	Expect(t, 2, len(split))
}

func TestPassword(t *testing.T) {
	i := New().Internet()

	Expect(t, true, len(i.Password()) >= 6)
}

func TestTLD(t *testing.T) {
	i := New().Internet()

	Expect(t, true, len(i.TLD()) > 0)
}

func TestSlug(t *testing.T) {
	i := New().Internet()

	Expect(t, true, len(i.Slug()) > 0)
}

func TestURL(t *testing.T) {
	i := New().Internet()

	Expect(t, true, len(i.URL()) > 0)
}

func TestIpv4(t *testing.T) {
	i := New().Internet()

	ip := i.Ipv4()
	Expect(t, true, len(ip) > 0)
	split := strings.Split(ip, ".")
	Expect(t, 4, len(split))
}

func TestLocalIpv4(t *testing.T) {
	i := New().Internet()

	ip := i.LocalIpv4()
	Expect(t, true, len(ip) > 0)
	split := strings.Split(ip, ".")
	Expect(t, 4, len(split))
}

func TestIpv6(t *testing.T) {
	i := New().Internet()

	ip := i.Ipv6()
	Expect(t, true, len(ip) > 0)
	Expect(t, 39, len(ip))

	split := strings.Split(ip, ":")
	Expect(t, 8, len(split))
}

func TestMacAddress(t *testing.T) {
	i := New().Internet()

	Expect(t, 17, len(i.MacAddress()))
}

func TestHTTPMethod(t *testing.T) {
	i := New().Internet()

	Expect(t, true, len(i.HTTPMethod()) > 0)
}

func TestQuery(t *testing.T) {
	i := New().Internet()

	query := i.Query()
	Expect(t, 0, strings.Index(query, "?"))
}
