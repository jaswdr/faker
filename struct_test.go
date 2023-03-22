package faker

import (
	"testing"
)

type Basic struct {
	s string
	S string
}

type Nested struct {
	A   string
	B   *Basic
	bar *Basic
}

type BuiltIn struct {
	Uint8   *uint8
	Uint16  *uint16
	Uint32  *uint32
	Uint64  *uint64
	Int     *int
	Int8    *int8
	Int16   *int16
	Int32   *int32
	Int64   *int64
	Float32 *float32
	Float64 *float64
	Bool    *bool
}

type Function struct {
	Number *string `fake:"#"`
	Name   *string `fake:"####"`
	Const  *string `fake:"ABC"`
}

type StructArray struct {
	Bars      []*Basic
	Builds    []BuiltIn
	Skips     []string  `fake:"skip"`
	Strings   []string  `fake:"####" fakesize:"3"`
	SetLen    [5]string `fake:"????"`
	SubStr    [][]string
	SubStrLen [][2]string
	Empty     []*Basic    `fakesize:"0"`
	Multy     []*Function `fakesize:"3"`
}

type NestedArray struct {
	NA []StructArray `fakesize:"2"`
}

func TestStructBasic(t *testing.T) {
	var basic Basic
	New().Struct().Fill(&basic)
	Expect(t, "", basic.s)
	NotExpect(t, "", basic.S)
}

func TestStructNested(t *testing.T) {
	var nested Nested
	New().Struct().Fill(&nested)
	NotExpect(t, "", nested.A)
	NotExpect(t, nil, nested.B)
	NotExpect(t, "", nested.B.S)
	NotExpect(t, nil, nested.bar)
}

func TestStructBuiltInTypes(t *testing.T) {
	var builtIn BuiltIn
	New().Struct().Fill(&builtIn)
	NotExpect(t, nil, builtIn.Uint8)
	NotExpect(t, nil, builtIn.Uint16)
	NotExpect(t, nil, builtIn.Uint32)
	NotExpect(t, nil, builtIn.Uint64)
	NotExpect(t, nil, builtIn.Int)
	NotExpect(t, nil, builtIn.Int8)
	NotExpect(t, nil, builtIn.Int16)
	NotExpect(t, nil, builtIn.Int32)
	NotExpect(t, nil, builtIn.Int64)
	NotExpect(t, nil, builtIn.Float32)
	NotExpect(t, nil, builtIn.Float64)
	NotExpect(t, nil, builtIn.Bool)
}

func TestStructWithFunction(t *testing.T) {
	var function Function
	New().Struct().Fill(&function)
	NotExpect(t, "", function.Number)
	NotExpect(t, "", function.Name)
	NotExpect(t, "ABC", function.Const)
}

func TestStructArray(t *testing.T) {
	var sa StructArray
	New().Struct().Fill(&sa)
	NotExpect(t, 0, sa.Bars)
	NotExpect(t, 0, sa.Builds)
	Expect(t, 3, len(sa.Strings))
	for _, s := range sa.Strings {
		NotExpect(t, "", s)
		Expect(t, 4, len(s))
	}
	Expect(t, 5, len(sa.SetLen))
	for _, s := range sa.SetLen {
		NotExpect(t, "", s)
	}
	for _, s := range sa.SubStr {
		for _, ss := range s {
			NotExpect(t, "", ss)
		}
	}
	for _, s := range sa.SubStrLen {
		Expect(t, 2, len(s))
		for _, ss := range s {
			NotExpect(t, "", ss)
		}
	}
	NotExpect(t, "", sa.Empty)
	NotExpect(t, "", sa.Skips)
	Expect(t, 3, len(sa.Multy))
}

func TestStructNestedArray(t *testing.T) {
	var na NestedArray
	New().Struct().Fill(&na)
	Expect(t, 2, len(na.NA))
	for _, elem := range na.NA {
		NotExpect(t, 0, len(elem.Builds))
		Expect(t, 0, len(elem.Empty))
		Expect(t, 0, len(elem.Skips))
		Expect(t, 3, len(elem.Multy))
	}
}

func TestStructToInt(t *testing.T) {
	var si struct {
		Int         int
		IntConst    int8  `fake:"-123"`
		IntGenerate int64 `fake:"{number:1,10}"`
	}
	New().Struct().Fill(&si)
	NotExpect(t, 0, si.Int)
	NotExpect(t, 0, si.IntConst)
	NotExpect(t, 0, si.IntGenerate)
}

func TestStructToUint(t *testing.T) {
	var su struct {
		Uint         uint
		UintConst    uint8  `fake:"123"`
		UintGenerate uint64 `fake:"{number:1,10}"`
	}
	New().Struct().Fill(&su)
	NotExpect(t, 0, su.Uint)
	NotExpect(t, 0, su.UintConst)
	NotExpect(t, 0, su.UintGenerate)
}

func TestStructToFloat(t *testing.T) {
	var sf struct {
		Float         float32
		FloatConst    float64 `fake:"123.456789"`
		FloatGenerate float32 `fake:"{latitude}"`
	}
	New().Struct().Fill(&sf)
	NotExpect(t, 0, sf.Float)
	NotExpect(t, 0, sf.FloatConst)
	NotExpect(t, 0, sf.FloatGenerate)
}

func TestStructToBool(t *testing.T) {
	var sf struct {
		Bool         bool
		BoolConst    bool `fake:"true"`
		BoolGenerate bool `fake:"{bool}"`
	}
	New().Struct().Fill(&sf)
	NotExpect(t, nil, sf.Bool)
	NotExpect(t, nil, sf.BoolConst)
	NotExpect(t, nil, sf.BoolGenerate)
}

func TestStructUUID(t *testing.T) {
	var st struct {
		UUID string `fake`
	}
	New().Struct().Fill(&st)
	NotExpect(t, "", st.UUID)
}

func TestStructUUIDInSequence(t *testing.T) {
	var st struct {
		UUID string `fake`
	}
	fake := New()
	before := ""
	for i := 0; i < 100; i++ {
		fake.Struct().Fill(&st)
		after := st.UUID
		NotExpect(t, true, after == "")
		Expect(t, false, before == after)
		before = after
	}
}
