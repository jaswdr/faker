package faker

import (
	"testing"
)

func TestCompanyCatchPhrase(t *testing.T) {
	f := New().Company()
	phrase := f.CatchPhrase()
	Expect(t, true, len(phrase) > 0)
}

func TestCompanyBS(t *testing.T) {
	f := New().Company()
	bs := f.BS()
	Expect(t, true, len(bs) > 0)
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
