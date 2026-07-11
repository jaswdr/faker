package faker

import (
	"strconv"
	"sync"
)

// Unique provides unique value generation, avoiding duplicates within a single Unique instance.
type Unique struct {
	Faker *Faker
	seen  map[string]struct{}
	mu    sync.Mutex
}

// Unique returns a Unique helper bound to this Faker instance.
func (f Faker) Unique() *Unique {
	return &Unique{
		Faker: &f,
		seen:  make(map[string]struct{}),
	}
}

// IntBetween returns a unique int between minN and maxN inclusive.
func (u *Unique) IntBetween(minN, maxN int) int {
	for attempt := 0; attempt < 1000; attempt++ {
		val := u.Faker.IntBetween(minN, maxN)
		key := strconv.Itoa(val)
		u.mu.Lock()
		if _, exists := u.seen[key]; !exists {
			u.seen[key] = struct{}{}
			u.mu.Unlock()
			return val
		}
		u.mu.Unlock()
	}
	return u.Faker.IntBetween(minN, maxN)
}

// StringElement returns a unique element from the given slice.
func (u *Unique) StringElement(elements []string) string {
	for attempt := 0; attempt < len(elements)*10; attempt++ {
		val := u.Faker.RandomStringElement(elements)
		u.mu.Lock()
		if _, exists := u.seen[val]; !exists {
			u.seen[val] = struct{}{}
			u.mu.Unlock()
			return val
		}
		u.mu.Unlock()
	}
	return u.Faker.RandomStringElement(elements)
}

// Reset clears all tracked unique values.
func (u *Unique) Reset() {
	u.mu.Lock()
	u.seen = make(map[string]struct{})
	u.mu.Unlock()
}
