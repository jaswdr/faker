package faker

import (
	"testing"
)

func TestPetName(t *testing.T) {
	a := New().Pet()
	NotExpect(t, "", a.Name())
}

func TestPetCat(t *testing.T) {
	a := New().Pet()
	NotExpect(t, "", a.Cat())
}

func TestPetDog(t *testing.T) {
	a := New().Pet()
	NotExpect(t, "", a.Dog())
}
