package faker

import (
	"regexp"
	"testing"
)

func TestUUIDv4(t *testing.T) {
	f := New()
	value := f.UUID().V4()
	match, err := regexp.MatchString("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$", value)
	Expect(t, true, err == nil)
	Expect(t, true, match)
}

func TestUUIDV4UniqueInSequence(t *testing.T) {
	f := New()
	before := f.UUID().V4()
	for i := 0; i < 100; i++ {
		after := f.UUID().V4()
		Expect(t, true, before != after)
		before = after
	}
}

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

func TestDeterminismWithMultipleFields(t *testing.T) {
	type ComplexStruct struct {
		StringField1 string
		StringField2 string
		IntField     int
		FloatField   float64
		BoolField    bool
		StringArray  []string `fakesize:"3"`
	}

	seed := int64(98765)

	f1 := NewWithSeedInt64(seed)
	var s1 ComplexStruct
	f1.Struct().Fill(&s1)

	f2 := NewWithSeedInt64(seed)
	var s2 ComplexStruct
	f2.Struct().Fill(&s2)

	if s1.StringField1 != s2.StringField1 {
		t.Errorf("StringField1 mismatch: %s != %s", s1.StringField1, s2.StringField1)
	}
	if s1.StringField2 != s2.StringField2 {
		t.Errorf("StringField2 mismatch: %s != %s", s1.StringField2, s2.StringField2)
	}
	if s1.IntField != s2.IntField {
		t.Errorf("IntField mismatch: %d != %d", s1.IntField, s2.IntField)
	}
	if s1.FloatField != s2.FloatField {
		t.Errorf("FloatField mismatch: %f != %f", s1.FloatField, s2.FloatField)
	}
	if s1.BoolField != s2.BoolField {
		t.Errorf("BoolField mismatch: %v != %v", s1.BoolField, s2.BoolField)
	}
	if len(s1.StringArray) != len(s2.StringArray) {
		t.Errorf("StringArray length mismatch: %d != %d", len(s1.StringArray), len(s2.StringArray))
	}
	for i := range s1.StringArray {
		if s1.StringArray[i] != s2.StringArray[i] {
			t.Errorf("StringArray[%d] mismatch: %s != %s", i, s1.StringArray[i], s2.StringArray[i])
		}
	}
}

func TestDeterminismWithUUIDDirectCall(t *testing.T) {
	seed := int64(54321)

	f1 := NewWithSeedInt64(seed)
	uuid1 := f1.UUID().V4()

	f2 := NewWithSeedInt64(seed)
	uuid2 := f2.UUID().V4()

	if uuid1 != uuid2 {
		t.Errorf("UUID mismatch: %s != %s", uuid1, uuid2)
	}
}

func TestDeterminismAcrossMultipleCalls(t *testing.T) {
	seed := int64(11111)

	f1 := NewWithSeedInt64(seed)
	uuids1 := make([]string, 5)
	for i := 0; i < 5; i++ {
		uuids1[i] = f1.UUID().V4()
	}

	f2 := NewWithSeedInt64(seed)
	uuids2 := make([]string, 5)
	for i := 0; i < 5; i++ {
		uuids2[i] = f2.UUID().V4()
	}

	for i := 0; i < 5; i++ {
		if uuids1[i] != uuids2[i] {
			t.Errorf("UUID[%d] mismatch: %s != %s", i, uuids1[i], uuids2[i])
		}
	}
}

func TestUUIDV4FormatValidation(t *testing.T) {
	f := New()

	for i := 0; i < 10; i++ {
		uuid := f.UUID().V4()

		if len(uuid) != 36 {
			t.Errorf("UUID has incorrect length: %d (expected 36)", len(uuid))
		}

		if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
			t.Errorf("UUID has incorrect dash positions: %s", uuid)
		}

		if uuid[14] != '4' {
			t.Errorf("UUID version should be 4, got: %c in UUID: %s", uuid[14], uuid)
		}

		variant := uuid[19]
		if variant != '8' && variant != '9' && variant != 'a' && variant != 'b' {
			t.Errorf("UUID variant should be 8, 9, a, or b, got: %c in UUID: %s", variant, uuid)
		}

		t.Logf("Valid UUID: %s", uuid)
	}
}
