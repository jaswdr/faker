package faker

import (
	"testing"
)

func TestLanguage(t *testing.T) {
	v := New().Language().Language()
	NotExpect(t, "", v)
}

func TestLanguageAbbr(t *testing.T) {
	v := New().Language().LanguageAbbr()
	NotExpect(t, "", v)
}

func TestProgrammingLanguage(t *testing.T) {
	v := New().Language().ProgrammingLanguage()
	NotExpect(t, "", v)
}
