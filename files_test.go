package cleanprops

import (
	"testing"
)

func TestListFiles(t *testing.T) {
	files := listFiles("testdata")
	if len(files) != 5 {
		t.Errorf("expected %d files got %d", 5, len(files))
	}
}

func TestListFilesWithExt(t *testing.T) {
	files := listFiles("testdata", ".xml")
	if len(files) != 1 {
		t.Errorf("expected %d files got %d", 1, len(files))
	}

	files = listFiles("testdata", ".JaVa")
	if len(files) != 2 {
		t.Errorf("expected %d files got %d", 2, len(files))
	}
}
