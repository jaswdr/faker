package faker

// RandomElement returns a fake random element from a given list of elements
func RandomElement[T any](f Faker, elements ...T) T {
	i := f.IntBetween(0, len(elements)-1)
	return elements[i]
}

// RandomElementWeighted takes faker instance and a list of elements in the form of map[weight]element,
// it then selects one of the elements randomly and returns it,
//
// Elements with higher weight have more chance to be returned.
func RandomElementWeighted[T any](f Faker, elements map[int]T) T {
	arrayOfElements := make([]T, 0)

	for weight, value := range elements {
		// TODO: there is an accepted proposal for generic slices.Repeat function in Go 1.23
		for i := 0; i < weight; i++ {
			arrayOfElements = append(arrayOfElements, value)
		}
	}

	i := f.IntBetween(0, len(arrayOfElements)-1)

	return arrayOfElements[i]
}
