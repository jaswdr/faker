package test

import "testing"

func TestFakeGenerator(t *testing.T) {
	s := NewFakeSource(100)

	if n := s.Int63(); n != 100 {
		t.Errorf("expected 100 but got %d", n)
	}

	s.Seed(200)

	if n := s.Int63(); n != 200 {
		t.Errorf("expected 200 but got %d", n)
	}
}
