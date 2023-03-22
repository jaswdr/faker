package faker

import (
	"fmt"
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
	p := New().Person()
	name, gender := p.NameAndGender()
	Expect(t, true, name != "")
	Expect(t, true, gender == "Male" || gender == "Female")
}

func TestSSN(t *testing.T) {
	p := New().Person()
	ssn := p.SSN()
	Expect(t, 9, len(ssn))
}

func TestContact(t *testing.T) {
	p := New().Person()
	contact := p.Contact()
	Expect(t, true, len(contact.Phone) > 0)
	Expect(t, true, len(contact.Email) > 0)
}

func TestPersonImage(t *testing.T) {
	p := New().Person()
	image := p.Image()
	Expect(t, fmt.Sprintf("%T", image), "*os.File")
	Expect(t, strings.HasSuffix(image.Name(), ".jpg"), true, image.Name())
}

func TestPersonaNameMale(t *testing.T) {
	p := New().Person()
	name := p.NameMale()
	Expect(t, true, len(name) > 0)
}
