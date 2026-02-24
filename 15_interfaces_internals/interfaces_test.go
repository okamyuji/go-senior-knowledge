package interfaces

import "testing"

func TestNilInterface(t *testing.T) {
	nilIface, nilPointerIface := NilInterface()

	if !nilIface {
		t.Error("nil interface should be == nil")
	}
	if nilPointerIface {
		t.Error("interface holding nil *Dog should NOT be == nil")
	}
}

func TestNilInterface_StringMethod(t *testing.T) {
	var d *Dog
	var s Stringer = d

	// Even though s holds nil, we can still call methods on it
	// because the type information is present.
	got := s.String()
	if got != "<nil dog>" {
		t.Errorf("nil *Dog String() = %q, want %q", got, "<nil dog>")
	}
}

func TestInterfaceComparison_SameType(t *testing.T) {
	sameType, _, _ := InterfaceComparison()
	if !sameType {
		t.Error("Cat{Mimi} == Cat{Mimi} should be true")
	}
}

func TestInterfaceComparison_DiffType(t *testing.T) {
	_, diffType, _ := InterfaceComparison()
	if diffType {
		t.Error("Cat{Mimi} == &Dog{Mimi} should be false (different types)")
	}
}

func TestInterfaceComparison_BothNil(t *testing.T) {
	_, _, bothNil := InterfaceComparison()
	if !bothNil {
		t.Error("two nil interfaces should be equal")
	}
}

func TestDynamicDispatch_Dog(t *testing.T) {
	d := &Dog{Name: "Rex"}
	got := DynamicDispatch(d)
	if got != "Dog: Rex" {
		t.Errorf("DynamicDispatch(Dog) = %q, want %q", got, "Dog: Rex")
	}
}

func TestDynamicDispatch_Cat(t *testing.T) {
	c := Cat{Name: "Whiskers"}
	got := DynamicDispatch(c)
	if got != "Cat: Whiskers" {
		t.Errorf("DynamicDispatch(Cat) = %q, want %q", got, "Cat: Whiskers")
	}
}

func TestDynamicDispatch_NilDog(t *testing.T) {
	var d *Dog
	got := DynamicDispatch(d)
	if got != "<nil dog>" {
		t.Errorf("DynamicDispatch(nil *Dog) = %q, want %q", got, "<nil dog>")
	}
}

func TestTypeInfo_Dog(t *testing.T) {
	d := &Dog{Name: "Buddy"}
	typeName, reflectType := TypeInfo(d)

	if typeName != "*interfaces.Dog" {
		t.Errorf("Sprintf %%T = %q, want %q", typeName, "*interfaces.Dog")
	}
	if reflectType != "*interfaces.Dog" {
		t.Errorf("reflect.TypeOf = %q, want %q", reflectType, "*interfaces.Dog")
	}
}

func TestTypeInfo_Cat(t *testing.T) {
	c := Cat{Name: "Luna"}
	typeName, reflectType := TypeInfo(c)

	if typeName != "interfaces.Cat" {
		t.Errorf("Sprintf %%T = %q, want %q", typeName, "interfaces.Cat")
	}
	if reflectType != "interfaces.Cat" {
		t.Errorf("reflect.TypeOf = %q, want %q", reflectType, "interfaces.Cat")
	}
}

func TestTypeInfo_Primitive(t *testing.T) {
	typeName, reflectType := TypeInfo(42)
	if typeName != "int" {
		t.Errorf("Sprintf %%T = %q, want %q", typeName, "int")
	}
	if reflectType != "int" {
		t.Errorf("reflect.TypeOf = %q, want %q", reflectType, "int")
	}
}
