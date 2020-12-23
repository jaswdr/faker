package faker

import (
	"testing"
)

func TestFoodFruit(t *testing.T) {
	v := New().Food().Fruit()
	NotExpect(t, "", v)
}

func TestFoodVegetable(t *testing.T) {
	v := New().Food().Vegetable()
	NotExpect(t, "", v)
}