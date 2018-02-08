package cleanprops

import (
	"testing"
)

func TestListFiles(t *testing.T) {
	files := listFiles("testdata")
	if len(files) < 1 {
		t.Errorf("expected some files")
	}
}
