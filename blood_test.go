package faker

import (
	"testing"
)

func TestBloodName(t *testing.T) {
	v := New().Blood().Name()
	NotExpect(t, "", v)
	ExpectInString(t, v, bloodTypes)
}


