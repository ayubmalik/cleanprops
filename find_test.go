package props

import (
	"fmt"
	"testing"
)

func TestFindHashFormat(t *testing.T) {
	props := Load("testdata/find.properties")
	result := Find(props.SortedKeys(), "testdata/HashPropsOnly.java", "#{key}")
	expectKeys(result, t, 2)
}

func TestFindDollarFormat(t *testing.T) {
	props := Load("testdata/find.properties")
	result := Find(props.SortedKeys(), "testdata/DollarPropsOnly.xml", "\\${key}")
	expectKeys(result, t, 4)
}

func TestFindMixedFormat(t *testing.T) {
	props := Load("testdata/find.properties")
	result := Find(props.SortedKeys(), "testdata/MixedProps.java", "\\${key}|#{key}")
	expectKeys(result, t, 2)
}

func TestFindInFiles(t *testing.T) {
	props := Load("testdata/find.properties")
	result := FindInFiles(props.SortedKeys(), "testdata/")
	expectKeys(result, t, 5)
}

func BenchmarkFind1(b *testing.B) {
	props := Load("testdata/find.properties")
	for n := 0; n < b.N; n++ {
		Find(props.SortedKeys(), "testdata/DollarPropsOnly.xml", "\\${key}|#{key}")
	}
}

func BenchmarkFind2(b *testing.B) {
	props := Load("testdata/find.properties")
	for n := 0; n < b.N; n++ {
		Find2(props.SortedKeys(), "testdata/DollarPropsOnly.xml", "\\${key}|#{key}")
	}
}

func expectKeys(result []Key, t *testing.T, count int) {
	if len(result) != count {
		t.Errorf("expected to find %d keys but was %d", count, len(result))
		fmt.Println(result)
	}
}
