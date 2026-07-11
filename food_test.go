package faker

import (
	"slices"
	"testing"
)

func TestFruit(t *testing.T) {
	v := New().Food().Fruit()
	Expect(t, true, slices.Contains(fruits, v))
}

func TestVegetable(t *testing.T) {
	v := New().Food().Vegetable()
	Expect(t, true, slices.Contains(vegetables, v))
}
