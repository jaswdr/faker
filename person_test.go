package faker

import (
	"strings"
	"testing"
)

func TestTitleMale(t *testing.T) {
	p := New().Person()
	Expect(t, "Mr.", p.TitleMale())
}

func TestTitleFemale(t *testing.T) {
	p := New().Person()
	Expect(t, "Ms.", p.TitleFemale())
}

func TestTitle(t *testing.T) {
	p := New().Person()
	Expect(t, 3, len(p.Title()))
}

func TestSuffix(t *testing.T) {
	p := New().Person()
	suffix := p.Suffix()
	Expect(t, true, len(suffix) > 0)
}

func TestFirstNameMale(t *testing.T) {
	p := New().Person()
	firstName := p.FirstNameMale()
	Expect(t, true, len(firstName) > 0)
}

func TestFirstNameFemale(t *testing.T) {
	p := New().Person()
	firstName := p.FirstNameFemale()
	Expect(t, true, len(firstName) > 0)
}

func TestFirstName(t *testing.T) {
	p := New().Person()
	firstName := p.FirstName()
	Expect(t, true, len(firstName) > 0)
}

func TestLastName(t *testing.T) {
	p := New().Person()
	lastName := p.LastName()
	Expect(t, true, len(lastName) > 0)
}

func TestName(t *testing.T) {
	p := New().Person()
	name := p.Name()
	Expect(t, true, len(name) > 0)
	Expect(t, false, strings.Contains(name, "{{titleMale}}"))
	Expect(t, false, strings.Contains(name, "{{firstNameMale}}"))
	Expect(t, false, strings.Contains(name, "{{titleFemale}}"))
	Expect(t, false, strings.Contains(name, "{{firstNameFemale}}"))
	Expect(t, false, strings.Contains(name, "{{lastName}}"))
	Expect(t, false, strings.Contains(name, "{{suffix}}"))
}
