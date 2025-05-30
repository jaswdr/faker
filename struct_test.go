package faker

import (
	"testing"
)

// Test types
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

// Test cases
func TestStructBasic(t *testing.T) {
	var basic Basic
	New().Struct().Fill(&basic)

	// Private field should remain empty
	Expect(t, "", basic.s)
	// Public field should be filled
	NotExpect(t, "", basic.S)
}

func TestStructNested(t *testing.T) {
	var nested Nested
	New().Struct().Fill(&nested)

	// Check nested struct fields
	NotExpect(t, "", nested.A)
	NotExpect(t, nil, nested.B)
	NotExpect(t, "", nested.B.S)
	NotExpect(t, nil, nested.bar)
}

func TestStructBuiltInTypes(t *testing.T) {
	var builtIn BuiltIn
	New().Struct().Fill(&builtIn)

	// Check all built-in types are filled
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

	// Check function tags are respected
	NotExpect(t, "", function.Number)
	NotExpect(t, "", function.Name)
	Expect(t, "ABC", *function.Const)
}

func TestStructArray(t *testing.T) {
	var sa StructArray
	New().Struct().Fill(&sa)

	// Check array fields
	NotExpect(t, 0, len(sa.Bars))
	NotExpect(t, 0, len(sa.Builds))

	// Check strings array with fakesize
	Expect(t, 3, len(sa.Strings))
	for _, s := range sa.Strings {
		NotExpect(t, "", s)
		Expect(t, 4, len(s))
	}

	// Check fixed length array
	Expect(t, 5, len(sa.SetLen))
	for _, s := range sa.SetLen {
		NotExpect(t, "", s)
	}

	// Check nested arrays
	for _, s := range sa.SubStr {
		for _, ss := range s {
			NotExpect(t, "", ss)
		}
	}

	// Check fixed length nested arrays
	for _, s := range sa.SubStrLen {
		Expect(t, 2, len(s))
		for _, ss := range s {
			NotExpect(t, "", ss)
		}
	}

	// Check empty array
	Expect(t, 0, len(sa.Empty))

	// Check skipped array
	Expect(t, 0, len(sa.Skips))

	// Check array with fakesize
	Expect(t, 3, len(sa.Multy))
}

func TestStructNestedArray(t *testing.T) {
	var na NestedArray
	New().Struct().Fill(&na)

	// Check nested array with fakesize
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

	// Test 100 sequential UUIDs
	for i := 0; i < 100; i++ {
		fake.Struct().Fill(&st)
		after := st.UUID

		// UUID should not be empty
		NotExpect(t, true, after == "")
		// UUID should be different from previous
		Expect(t, false, before == after)
		before = after
	}
}

func TestStructWithChildStructures(t *testing.T) {
	type Child struct {
		Name   string
		Age    int
		Active bool
	}

	type Parent struct {
		ID       string `fake`
		Child    Child
		Children []Child `fakesize:"3"`
		Optional *Child
	}

	var p Parent
	New().Struct().Fill(&p)

	// Test parent fields
	NotExpect(t, "", p.ID)

	// Test embedded child struct
	NotExpect(t, "", p.Child.Name)
	NotExpect(t, 0, p.Child.Age)
	NotExpect(t, nil, p.Child.Active)

	// Test slice of child structs
	Expect(t, 3, len(p.Children))
	for _, child := range p.Children {
		NotExpect(t, "", child.Name)
		NotExpect(t, 0, child.Age)
		NotExpect(t, nil, child.Active)
	}

	// Test pointer to child struct
	NotExpect(t, nil, p.Optional)
	NotExpect(t, "", p.Optional.Name)
	NotExpect(t, 0, p.Optional.Age)
	NotExpect(t, nil, p.Optional.Active)
}

func TestNestedStructures(t *testing.T) {
	type Address struct {
		Street string
		City   string
		ZIP    string `fake:"#####"`
	}

	type Contact struct {
		Email   string `fake:"????.????@example.com"`
		Phone   string `fake:"###-###-####"`
		Address Address
	}

	type Person struct {
		ID      uint64
		Name    string
		Contact Contact
	}

	var p Person
	New().Struct().Fill(&p)

	// Test top level fields
	NotExpect(t, 0, p.ID)
	NotExpect(t, "", p.Name)

	// Test nested contact
	NotExpect(t, "", p.Contact.Email)
	NotExpect(t, "", p.Contact.Phone)

	// Test deeply nested address
	NotExpect(t, "", p.Contact.Address.Street)
	NotExpect(t, "", p.Contact.Address.City)
	NotExpect(t, "", p.Contact.Address.ZIP)
}

func TestRecursiveStruct(t *testing.T) {
	type Node struct {
		Value    string
		Parent   *Node
		Children []*Node `fakesize:"2"`
	}

	var root Node
	New().Struct().Fill(&root)

	NotExpect(t, nil, root.Parent)
	Expect(t, 2, len(root.Children))

	for _, child := range root.Children {
		NotExpect(t, nil, child.Parent)
		Expect(t, 2, len(child.Children))
		for _, grandchild := range child.Children {
			NotExpect(t, nil, grandchild.Parent)
			Expect(t, 2, len(grandchild.Children))
		}
	}
}

func TestStructWithDepth(t *testing.T) {
	type Node struct {
		Value string
		Child *Node `fakesize:"2"`
	}

	var n Node
	New().Struct().FillWithDepth(&n, 3)

	NotExpect(t, "", n.Value)
	NotExpect(t, "", n.Child.Value)
	Expect(t, "", n.Child.Child.Value)
}
