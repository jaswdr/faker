package faker

import (
	"regexp"
	"testing"
)

func TestProgrammingLanguageName(t *testing.T) {
	f := New()
	pl := f.ProgrammingLanguage()
	names := make([]string, 0, len(languageVersions))
	for lang := range languageVersions {
		names = append(names, lang)
	}
	ExpectInString(t, pl.Name(), names)
}

func TestProgrammingLanguageVersion(t *testing.T) {
	f := New()
	pl := f.ProgrammingLanguage()
	name := pl.Name()
	ExpectInString(t, pl.Version(name), languageVersions[name])
}

func TestProgrammingLanguageVariableName(t *testing.T) {
	f := New()
	pl := f.ProgrammingLanguage()
	pattern := regexp.MustCompile(`^[_a-zA-Z][_a-zA-Z0-9]*$`)

	// Test multiple generations
	for i := 0; i < 100; i++ {
		name := pl.VariableName()
		Expect(t, true, pattern.MatchString(name))
	}
}

func TestProgrammingLanguageVariableNameWithLength(t *testing.T) {
	f := New()
	pl := f.ProgrammingLanguage()
	pattern := regexp.MustCompile(`^[_a-zA-Z][_a-zA-Z0-9]*$`)

	// Test different lengths
	lengths := []int{1, 5, 10, 20, 50}
	for _, length := range lengths {
		name := pl.VariableNameWithLength(length)
		Expect(t, true, pattern.MatchString(name))
		Expect(t, length, len(name))
	}

	// Test invalid length (should default to 1)
	name := pl.VariableNameWithLength(0)
	Expect(t, 1, len(name))
}

func TestProgrammingLanguageVariableNameDistribution(t *testing.T) {
	f := New()
	pl := f.ProgrammingLanguage()

	// Generate 1000 names and check their distribution
	names := make(map[string]int)
	totalNames := 1000
	for i := 0; i < totalNames; i++ {
		name := pl.VariableName()
		names[name]++
	}

	// Check that we have a reasonable number of unique names
	uniqueNames := len(names)
	Expect(t, true, uniqueNames >= totalNames/2)

	// Check that no single name appears too frequently
	maxOccurrences := 0
	for _, count := range names {
		if count > maxOccurrences {
			maxOccurrences = count
		}
	}
	Expect(t, true, maxOccurrences <= totalNames/10)
}
