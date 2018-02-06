package cleanprops

import (
	"testing"
)

func TestLoad(t *testing.T) {
	props, err := load("testdata/hello.properties")
	if err != nil {
		t.Errorf("returned nil %s", err)
	}

	value := props["some.key1"]
	if value != "some.value1" {
		t.Errorf("value is incorrect: %s", value)
	}

	if len(props) != 5 {
		t.Errorf("expected 5 props in hello.props")
	}
}

func TestTrimWhitespace(t *testing.T) {
	props, _ := load("testdata/hello.properties")
	if props == nil {
		t.Errorf("returned nil")
	}

	value := props["this.is.a.long.key"]
	if value != "value.with.leading.space" {
		t.Errorf("value is incorrect %q", value)
	}
}
