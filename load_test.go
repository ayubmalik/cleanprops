package cleanprops

import (
	"testing"
)

func TestLoad(t *testing.T) {
	props, err := Load("testdata/hello.properties")
	if err != nil {
		t.Errorf("returned nil %s", err)
	}

	if props == nil {
		t.Errorf("returned nil")
	}

	prop1 := props["some.key1"]
	if prop1 != "some.value1" {
		t.Errorf("prop1 value is incorrect: %s", prop1)
	}
}
