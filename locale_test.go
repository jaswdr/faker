package faker

import "testing"

func TestDefaultLocale(t *testing.T) {
	f := New()
	Expect(t, LocaleEnUS, f.Locale())
}

func TestWithLocale(t *testing.T) {
	f := New().WithLocale(LocaleDeDE)
	Expect(t, LocaleDeDE, f.Locale())
}

func TestNewWithLocale(t *testing.T) {
	f := NewWithLocale(LocaleDeDE)
	Expect(t, LocaleDeDE, f.Locale())
}

func TestGermanLocaleNames(t *testing.T) {
	f := NewWithSeedInt64(42).WithLocale(LocaleDeDE)
	p := f.Person()

	male := p.FirstNameMale()
	female := p.FirstNameFemale()
	last := p.LastName()

	Expect(t, true, contains(localeDeDE.FirstNameMale, male))
	Expect(t, true, contains(localeDeDE.FirstNameFemale, female))
	Expect(t, true, contains(localeDeDE.LastName, last))
}

func TestRegisterLocale(t *testing.T) {
	custom := LocaleData{
		Code:            "test_XX",
		FirstNameMale:   []string{"TestMale"},
		FirstNameFemale: []string{"TestFemale"},
		LastName:        []string{"TestLast"},
	}
	RegisterLocale(custom)

	f := New().WithLocale("test_XX")
	Expect(t, "TestMale", f.Person().FirstNameMale())
	Expect(t, "TestFemale", f.Person().FirstNameFemale())
	Expect(t, "TestLast", f.Person().LastName())
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
