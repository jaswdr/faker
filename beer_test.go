package faker

import (
	"testing"
)

func TestBeerName(t *testing.T) {
	v := New().Beer().Name()
	NotExpect(t, "", v)
}

func TestBeerStyle(t *testing.T) {
	v := New().Beer().Style()
	NotExpect(t, "", v)
}

func TestBeerHop(t *testing.T) {
	v := New().Beer().Hop()
	NotExpect(t, "", v)
}

func TestBeerMalt(t *testing.T) {
	v := New().Beer().Malt()
	NotExpect(t, "", v)
}

func TestBeerAlcohol(t *testing.T) {
	v := New().Beer().Alcohol()
	NotExpect(t, "", v)
}

func TestBeerIbu(t *testing.T) {
	v := New().Beer().Ibu()
	NotExpect(t, "", v)
}

func TestBeerBlg(t *testing.T) {
	v := New().Beer().Blg()
	NotExpect(t, "", v)
}
