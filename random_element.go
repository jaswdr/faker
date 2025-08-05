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
	if len(elements) == 0 {
		var zero T
		return zero
	}

	totalWeight := 0
	for weight := range elements {
		totalWeight += weight
	}

	if totalWeight == 0 {
		for _, value := range elements {
			return value
		}
	}

	target := f.IntBetween(0, totalWeight-1)

	current := 0
	for weight, value := range elements {
		current += weight
		if target < current {
			return value
		}
	}

	var zero T
	for _, value := range elements {
		return value
	}
	return zero
}

func RandomMapKey[K comparable, V any](f Faker, m map[K]V) K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	i := f.IntBetween(0, len(keys)-1)
	return keys[i]
}

func RandomMapValue[K comparable, V any](f Faker, m map[K]V) V {
	values := make([]V, 0, len(m))
	for k := range m {
		values = append(values, m[k])
	}

	i := f.IntBetween(0, len(values)-1)
	return values[i]
}
