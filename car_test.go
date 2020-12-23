package faker

import (
	"testing"
)

func TestCarMaker(t *testing.T) {
	v := New().Car().Maker()
	NotExpect(t, "", v)
}

func TestCarModel(t *testing.T) {
	v := New().Car().Model()
	NotExpect(t, "", v)
}

func TestCarCategory(t *testing.T) {
	v := New().Car().Category()
	NotExpect(t, "", v)
}

func TestCarFuelType(t *testing.T) {
	v := New().Car().FuelType()
	NotExpect(t, "", v)
}

func TestCarTransmissionGear(t *testing.T) {
	v := New().Car().TransmissionGear()
	NotExpect(t, "", v)
}
