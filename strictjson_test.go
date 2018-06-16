package strictjson

import (
	"testing"
)

type testStruct struct {
	A int
	B string
	c float64
}

func TestNormalConditions(t *testing.T) {
	actual := testStruct{A: 4, B: "q", c: 5.6}
	err := UnmarshalStruct([]byte(`{"A":8,"B":"y"}`), &actual)
	if err != nil {
		t.Error(err)
	}
	expected := testStruct{A: 8, B: "y", c: 5.6}
	if actual != expected {
		t.Errorf("%v != %v", actual, expected)
	}
}

func TestNilPointer(t *testing.T) {
	var ptr *testStruct
	err := UnmarshalStruct([]byte(`{"A":8,"B":"y"}`), ptr)
	if err == nil {
		t.Error("Expected error")
	}
	t.Logf("Returned:\n%v", err)
}

func TestMissingField(t *testing.T) {
	actual := testStruct{A: 4, B: "q", c: 5.6}
	expected := testStruct{A: 4, B: "q", c: 5.6}
	err := UnmarshalStruct([]byte(`{"B":"y"}`), &actual)
	if err == nil {
		t.Error("Expected error")
	}
	t.Logf("Returned:\n%v", err)
	if actual != expected {
		t.Errorf("%v != %v", actual, expected)
	}
}

func TestExtraField(t *testing.T) {
	actual := testStruct{A: 4, B: "q", c: 5.6}
	expected := testStruct{A: 4, B: "q", c: 5.6}
	err := UnmarshalStruct([]byte(`{"A":8,"B":"y","R":5}`), &actual)
	if err == nil {
		t.Error("Expected error")
	}
	t.Logf("Returned:\n%v", err)
	if actual != expected {
		t.Errorf("%v != %v", actual, expected)
	}
}

func TestUnexportedField(t *testing.T) {
	actual := testStruct{A: 4, B: "q", c: 5.6}
	expected := testStruct{A: 4, B: "q", c: 5.6}
	err := UnmarshalStruct([]byte(`{"A":8,"B":"y","c":8.2}`), &actual)
	if err == nil {
		t.Error("Expected error")
	}
	t.Logf("Returned:\n%v", err)
	if actual != expected {
		t.Errorf("%v != %v", actual, expected)
	}
}

func TestWrongFieldType(t *testing.T) {
	actual := testStruct{A: 4, B: "q", c: 5.6}
	expected := testStruct{A: 4, B: "q", c: 5.6}
	err := UnmarshalStruct([]byte(`{"A":8,"B":5}`), &actual)
	if err == nil {
		t.Error("Expected error")
	}
	t.Logf("Returned:\n%v", err)
	if actual != expected {
		t.Errorf("%v != %v", actual, expected)
	}
}

func TestInvalidJSON(t *testing.T) {
	actual := testStruct{A: 4, B: "q", c: 5.6}
	expected := testStruct{A: 4, B: "q", c: 5.6}
	err := UnmarshalStruct([]byte(`34r587fiyhie4`), &actual)
	if err == nil {
		t.Error("Expected error")
	}
	t.Logf("Returned:\n%v", err)
	if actual != expected {
		t.Errorf("%v != %v", actual, expected)
	}
}

func TestNonPointer(t *testing.T) {
	actual := testStruct{A: 4, B: "q", c: 5.6}
	expected := testStruct{A: 4, B: "q", c: 5.6}
	err := UnmarshalStruct([]byte(`{"A":8,"B":"y"}`), actual)
	if err == nil {
		t.Error("Expected error")
	}
	t.Logf("Returned:\n%v", err)
	if actual != expected {
		t.Errorf("%v != %v", actual, expected)
	}
}

func TestCaseSensitive(t *testing.T) {
	actual := testStruct{A: 4, B: "q", c: 5.6}
	expected := testStruct{A: 4, B: "q", c: 5.6}
	err := UnmarshalStruct([]byte(`{"A":8,"b":"y"}`), &actual)
	if err == nil {
		t.Error("Expected error")
	}
	t.Logf("Returned:\n%v", err)
	if actual != expected {
		t.Errorf("%v != %v", actual, expected)
	}
}
