package faker

import (
	"testing"
)

func TestBloodName(t *testing.T) {
	v := New().Blood().Name()
	NotExpect(t, "", v)
	Expect(t, true, v == "A+" || v == "A-" || v == "B+" || v == "B-" || v == "AB+" || v == "AB-" || v == "O+" || v == "O-")
}

func TestBloodAbbr(t *testing.T) {
	v := New().Blood().Abbr()
	NotExpect(t, "", v)
	Expect(t, true, v == "A+" || v == "A-" || v == "B+" || v == "B-" || v == "AB+" || v == "AB-" || v == "O+" || v == "O-")
}
