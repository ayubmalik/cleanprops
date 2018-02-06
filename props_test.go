package cleanprops

import (
	"strings"
	"testing"
)

func TestLoadProps(t *testing.T) {
	props := LoadProps("testdata/hello.properties")
	if !strings.HasSuffix(props.FileName, "testdata/hello.properties") {
		t.Errorf("filename is incorrect %q", props.FileName)
	}
}

func TestSortedKeys(t *testing.T) {
	props := Props{"filename", map[Key]Value{"a": "1", "c": "3", "b": "2"}}
	keys := props.SortedKeys()

	for i, v := range []Key{"a", "b", "c"} {
		if keys[i] != v {
			t.Errorf("Keys not sorted")
			break
		}
	}
}
