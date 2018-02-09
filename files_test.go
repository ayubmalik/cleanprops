package props

import (
	"testing"
)

func TestListFiles(t *testing.T) {
	files := listFiles("testdata")
	if len(files) != 6 {
		t.Errorf("expected %d files got %d", 6, len(files))
	}
}

func TestListFilesWithExt(t *testing.T) {
	files := listFiles("testdata", ".xml")
	if len(files) != 1 {
		t.Errorf("expected %d files got %d", 1, len(files))
	}

	files = listFiles("testdata", ".JaVa")
	if len(files) != 3 {
		t.Errorf("expected %d files got %d", 3, len(files))
	}
}
