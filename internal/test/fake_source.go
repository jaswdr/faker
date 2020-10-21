package test

// FakeSource implements the rand.Source interface. This is useful for when we
// want to supply dummy/predictable values to test specific scenarios that depend
// on rand.Rand.
type FakeSource struct {
	seed int64
}

// NewFakeSource returns a new instance of FakeSource. FakeSource does not
// generates random numbers, so the passed in seed number is the returned "random"
// number.
func NewFakeSource(seed int64) *FakeSource {
	return &FakeSource{seed: seed}
}

// Seed sets the internal int to the passed in seed value. This will be the returned
// value when using FakeSource.Int63().
func (s *FakeSource) Seed(seed int64) {
	s.seed = seed
}

// Int63 returns the dummy seed value.
func (s *FakeSource) Int63() int64 {
	return s.seed
}
