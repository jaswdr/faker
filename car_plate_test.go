package faker

import (
	"testing"
)

func TestCarPlateRegion(t *testing.T) {
	v := New().CarPlate().Region()
	NotExpect(t, "", v)
}

func TestCarPlateCode(t *testing.T) {
	v := New().CarPlate().Code()
	NotExpect(t, "", v)
}

func TestCarPlateSeries(t *testing.T) {
	v := New().CarPlate().Series()
	NotExpect(t, "", v)
}
