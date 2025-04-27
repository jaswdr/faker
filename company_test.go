package faker

import (
	"strings"
	"testing"
)

func TestCompanyCatchPhrase(t *testing.T) {
	f := New().Company()
	phrase := f.CatchPhrase()
	Expect(t, true, len(phrase) > 0)
	Expect(t, 2, strings.Count(phrase, " ")) // 3 words
}

func TestCompanyBS(t *testing.T) {
	f := New().Company()
	bs := f.BS()
	Expect(t, true, len(bs) > 0)
	Expect(t, 2, strings.Count(bs, " ")) // 3 words
}

func TestCompanySuffix(t *testing.T) {
	f := New().Company()
	value := f.Suffix()
	Expect(t, true, len(value) > 0)
}

func TestCompanyName(t *testing.T) {
	f := New().Company()
	value := f.Name()
	Expect(t, true, len(value) > 0)
}

func TestCompanyJobTitle(t *testing.T) {
	f := New().Company()
	value := f.JobTitle()
	Expect(t, true, len(value) > 0)
}

func TestEIN(t *testing.T) {
	f := New().Company()
	value := f.EIN()
	Expect(t, true, value > 0)
}
