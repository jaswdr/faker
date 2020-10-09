package faker

import (
	"reflect"
	"testing"
)

func TestBooleanBool(t *testing.T) {
	f := New().Boolean()
	tp := reflect.TypeOf(f.Bool())
	Expect(t, "bool", tp.String())
}

func TestBooleanBoolWithChance(t *testing.T) {
	f := New().Boolean()
	tp := reflect.TypeOf(f.BoolWithChance(30))
	Expect(t, "bool", tp.String())

	Expect(t, true, f.BoolWithChance(100))
	Expect(t, false, f.BoolWithChance(0))
	Expect(t, true, f.BoolWithChance(101))
	Expect(t, false, f.BoolWithChance(-1))
}

func TestBooleanInt(t *testing.T) {
	p := New().Boolean()
	result := p.IntBool()
	Expect(t, true, result == 1 || result == 0)
}

func TestBooleanString(t *testing.T) {
	p := New().Boolean()
	args := []string{"yes", "no"}
	result := p.StringBool(args[0], args[1])
	Expect(t, true, result == args[0] || result == args[1])
}
