package faker

import (
	"slices"
	"testing"
)

var expectedBloodTypes = []string{"A+", "A-", "B+", "B-", "AB+", "AB-", "O+", "O-"}

func TestBloodName(t *testing.T) {
	v := New().Blood().Name()
	Expect(t, true, slices.Contains(expectedBloodTypes, v))
}
