package faker

import (
	"strings"
	"testing"

	"github.com/jaswdr/faker/internal/test"
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
	// Test Title Male
	s := test.NewFakeSource(0)
	f := NewWithSeed(s)
	maleTitle := f.Person().Title()

	Expect(t, 3, len(maleTitle))

	// Test Title Female
	s.Seed(1)
	femaleTitle := f.Person().Title()

	Expect(t, 3, len(femaleTitle))

	if maleTitle == femaleTitle {
		t.Errorf("expected male title '%s' to be different from female title '%s'", maleTitle, femaleTitle)
	}
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
	// Set fake generator int to 6 to grab this maleNameFormat:
	// "{{titleMale}} {{firstNameMale}} {{lastName}} {{suffix}}"
	s := test.NewFakeSource(6)
	f := NewWithSeed(s)
	name := f.Person().Name()

	Expect(t, true, len(name) > 0)
	Expect(t, false, strings.Contains(name, "{{titleMale}}"))
	Expect(t, false, strings.Contains(name, "{{firstNameMale}}"))
	Expect(t, false, strings.Contains(name, "{{titleFemale}}"))
	Expect(t, false, strings.Contains(name, "{{firstNameFemale}}"))
	Expect(t, false, strings.Contains(name, "{{lastName}}"))
	Expect(t, false, strings.Contains(name, "{{suffix}}"))

	// Set int to 13 to grab this femaleNameFormat:
	// "{{titleFemale}} {{firstNameFemale}} {{lastName}} {{suffix}}"
	s.Seed(13)
	name = f.Person().Name()

	Expect(t, true, len(name) > 0)
	Expect(t, false, strings.Contains(name, "{{titleMale}}"))
	Expect(t, false, strings.Contains(name, "{{firstNameMale}}"))
	Expect(t, false, strings.Contains(name, "{{titleFemale}}"))
	Expect(t, false, strings.Contains(name, "{{firstNameFemale}}"))
	Expect(t, false, strings.Contains(name, "{{lastName}}"))
	Expect(t, false, strings.Contains(name, "{{suffix}}"))
}

func TestGender(t *testing.T) {
	p := New().Person()
	gender := p.Gender()
	Expect(t, true, gender == "Male" || gender == "Female")
}

func TestGenderMale(t *testing.T) {
	p := New().Person()
	Expect(t, "Male", p.GenderMale())
}

func TestGenderFemale(t *testing.T) {
	p := New().Person()
	Expect(t, "Female", p.GenderFemale())
}

func TestNameAndGender(t *testing.T) {
	// Test Name and Gender Female.
	s := test.NewFakeSource(0)
	f := NewWithSeed(s)
	name, gender := f.Person().NameAndGender()

	Expect(t, true, name != "")
	Expect(t, true, gender == "Female")

	// Test Name and Gender Male.
	s.Seed(51)
	name, gender = f.Person().NameAndGender()
	Expect(t, true, name != "")
	Expect(t, true, gender == "Male")
}
