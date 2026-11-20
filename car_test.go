package faker

import (
	"regexp"
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

func TestCarPlate(t *testing.T) {
	v := New().Car().Plate()
	NotExpect(t, "", v)
	Expect(t, true, regexp.MustCompile(`^[A-Z]{2}[0-9]{4}[A-Z]{2}$`).MatchString(v), v)
}

func TestCarPlateLettersAreLatin(t *testing.T) {
	re := regexp.MustCompile(`^[A-Z]{2}$`)
	for _, code := range append(append([]string{}, countryRegions...), carSeries...) {
		Expect(t, true, re.MatchString(code), code)
	}
}
