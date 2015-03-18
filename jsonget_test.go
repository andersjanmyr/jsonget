package jsonget

import (
	"reflect"
	"testing"
)

const data = `{"name":"Wednesday", "age":6, "color": { "red": "#ff0000", "green": "#00ff00", "blue": "#0000ff" }, "parents":["Gomez","Morticia"], "pets":[{"name": "Fido"}, {"name": "Misse"}] }`

func TestJsonGetString(t *testing.T) {
	value, _ := JsonGet(data, "name")
	expected := "Wednesday"
	if value != expected {
		t.Errorf("JsonGet() = %v, expected: %v", value, expected)
	}
}

func TestJsonGetFloat64(t *testing.T) {
	value, _ := JsonGet(data, "age")
	expected := 6.0
	if value != expected {
		t.Errorf("JsonGet() = %v, expected: %v", value, expected)
	}
}

func TestJsonGetArray(t *testing.T) {
	value, _ := JsonGet(data, "parents")
	expected := []interface{}{"Gomez", "Morticia"}
	if !reflect.DeepEqual(value, expected) {
		t.Errorf("JsonGet() = %#v, expected: %#v", value, expected)
	}
}

func TestJsonGetMap(t *testing.T) {
	value, _ := JsonGet(data, "color")
	expected := map[string]interface{}{"red": "#ff0000", "green": "#00ff00", "blue": "#0000ff"}
	if !reflect.DeepEqual(value, expected) {
		t.Errorf("JsonGet() = %#v, expected: %#v", value, expected)
	}
}

func TestJsonGetNestedMap(t *testing.T) {
	value, _ := JsonGet(data, "color.red")
	expected := "#ff0000"
	if value != expected {
		t.Errorf("JsonGet() = %#v, expected: %#v", value, expected)
	}
}

func TestJsonGetNestedArray(t *testing.T) {
	value, _ := JsonGet(data, "parents.1")
	expected := "Morticia"
	if value != expected {
		t.Errorf("JsonGet() = %#v, expected: %#v", value, expected)
	}
}

func TestJsonGetNestedArrayMap(t *testing.T) {
	value, _ := JsonGet(data, "pets.1.name")
	expected := "Misse"
	if value != expected {
		t.Errorf("JsonGet() = %#v, expected: %#v", value, expected)
	}
}

func TestJsonGetNestedArrayMapStar(t *testing.T) {
	value, _ := JsonGet(data, "pets.*.name")
	expected := []interface{}{"Fido", "Misse"}
	if !reflect.DeepEqual(value, expected) {
		t.Errorf("JsonGet() = %#v, expected: %#v", value, expected)
	}
}

func TestJsonGetNestedArrayInvalidIndex(t *testing.T) {
	value, err := JsonGet(data, "parents.3")
	if err == nil {
		t.Errorf("JsonGet() = %#v, expected: Error", value)
	}
}

func TestJsonGetNestedArrayStringIndex(t *testing.T) {
	value, err := JsonGet(data, "parents.dingo")
	if err == nil {
		t.Errorf("JsonGet() = %#v, expected: Error", value)
	}
}
