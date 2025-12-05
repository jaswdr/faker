package faker

import "testing"

type TestStructDeterminism struct {
	Name  string
	Email string
	ID    string
}

func TestDeterminism(t *testing.T) {
	seed := int64(12345)

	f1 := NewWithSeedInt64(seed)
	var s1 TestStructDeterminism
	f1.Struct().Fill(&s1)
	t.Logf("Run 1: Name: %s, Email: %s, ID: %s", s1.Name, s1.Email, s1.ID)

	f2 := NewWithSeedInt64(seed)
	var s2 TestStructDeterminism
	f2.Struct().Fill(&s2)
	t.Logf("Run 2: Name: %s, Email: %s, ID: %s", s2.Name, s2.Email, s2.ID)

	if s1.Name != s2.Name {
		t.Errorf("Expected same Name but got different values")
	}
	if s1.Email != s2.Email {
		t.Errorf("Expected same Email but got different values")
	}
	if s1.ID != s2.ID {
		t.Errorf("Expected same ID but got different values")
	}
}
