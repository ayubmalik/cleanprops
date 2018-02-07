package cleanprops

import (
	"bufio"
	"os"
	"strings"
)

func Find(keys []Key, srcFile string, format string) []Key {
	file, err := os.Open(srcFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	result := make([]Key, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		findKeys(diff(keys, result), scanner.Text(), format, &result)
	}
	return result
}

func findKeys(keys []Key, text string, format string, result *[]Key) {
	for _, k := range keys {
		formatted := strings.Replace(format, "key", string(k), 1)
		if strings.Contains(text, formatted) {
			*result = append(*result, k)
		}
	}
}

func diff(a, b []Key) []Key {
	mb := map[Key]bool{}
	for _, x := range b {
		mb[x] = true
	}
	ab := []Key{}
	for _, x := range a {
		if _, ok := mb[x]; !ok {
			ab = append(ab, x)
		}
	}
	return ab
}

// TODO: move
func dedupe(elements []Key) []Key {
	found := map[Key]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		found[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []Key{}
	for key, _ := range found {
		result = append(result, key)
	}
	return result
}
