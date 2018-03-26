package faker

import (
	"strings"
	"testing"
)

func TestCityPrefix(t *testing.T) {
	a := New().Address()
	Expect(t, true, len(a.CityPrefix()) > 0)
}

func TestSecondaryAddress(t *testing.T) {
	a := New().Address()
	Expect(t, true, len(a.SecondaryAddress()) > 0)
}

func TestState(t *testing.T) {
	a := New().Address()
	Expect(t, true, len(a.State()) > 0)
}

func TestStateAbbr(t *testing.T) {
	a := New().Address()
	Expect(t, true, len(a.StateAbbr()) > 0)
}

func TestCitySuffix(t *testing.T) {
	a := New().Address()
	Expect(t, true, len(a.CitySuffix()) > 0)
}

func TestStreetSuffix(t *testing.T) {
	a := New().Address()
	Expect(t, true, len(a.StreetSuffix()) > 0)
}

func TestBuildingNumber(t *testing.T) {
	a := New().Address()
	Expect(t, true, len(a.BuildingNumber()) > 0)
}

func TestCity(t *testing.T) {
	a := New().Address()
	city := a.City()
	Expect(t, true, len(city) > 0)
	Expect(t, false, strings.Contains(city, "{{cityPrefix}}"))
	Expect(t, false, strings.Contains(city, "{{firstName}}"))
	Expect(t, false, strings.Contains(city, "{{lastName}}"))
	Expect(t, false, strings.Contains(city, "{{citySuffix}}"))
	Expect(t, false, strings.Contains(city, "{{"))
	Expect(t, false, strings.Contains(city, "}}"))
}

func TestStreetName(t *testing.T) {
	a := New().Address()
	streetName := a.StreetName()
	Expect(t, true, len(streetName) > 0)
	Expect(t, false, strings.Contains(streetName, "{{firstName}}"))
	Expect(t, false, strings.Contains(streetName, "{{firstName}}"))
	Expect(t, false, strings.Contains(streetName, "{{lastName}}"))
	Expect(t, false, strings.Contains(streetName, "{{streetSuffix}}"))
	Expect(t, false, strings.Contains(streetName, "{{"))
	Expect(t, false, strings.Contains(streetName, "}}"))
}

func TestStreetAddress(t *testing.T) {
	a := New().Address()
	streetAddress := a.StreetAddress()
	Expect(t, true, len(streetAddress) > 0)
	Expect(t, false, strings.Contains(streetAddress, "{{firstName}}"))
	Expect(t, false, strings.Contains(streetAddress, "{{firstName}}"))
	Expect(t, false, strings.Contains(streetAddress, "{{lastName}}"))
	Expect(t, false, strings.Contains(streetAddress, "{{streetSuffix}}"))
	Expect(t, false, strings.Contains(streetAddress, "{{"))
	Expect(t, false, strings.Contains(streetAddress, "}}"))
}

func TestPostCode(t *testing.T) {
	a := New().Address()
	code := a.PostCode()
	Expect(t, true, len(code) > 0)
}

func TestAddress(t *testing.T) {
	a := New().Address()
	address := a.Address()
	Expect(t, true, len(address) > 0)
	Expect(t, false, strings.Contains(address, "{{streetAddress}}"))
	Expect(t, false, strings.Contains(address, "{{city}}"))
	Expect(t, false, strings.Contains(address, "{{stateAbbr}}"))
	Expect(t, false, strings.Contains(address, "{{postCode}}"))
	Expect(t, false, strings.Contains(address, "{{"))
	Expect(t, false, strings.Contains(address, "}}"))
}

func TestCountry(t *testing.T) {
	a := New().Address()
	Expect(t, true, len(a.Country()) > 0)
}

func TestLatitude(t *testing.T) {
	a := New().Address()
	Expect(t, true, a.Latitude() > 0)
}

func TestLongitude(t *testing.T) {
	a := New().Address()
	Expect(t, true, a.Longitude() > 0)
}
