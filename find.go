package props

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"fmt"
)

func FindInFiles(keys []Key, srcDir string, ext ...string) []Key {
	files := listFiles(srcDir, ext...)
	result := make([]Key, 0)
	for _, f := range files {
		if notPropsFile(f) {
			fmt.Println(f)
			r := Find(keys, f, "\\${key}")
			result = append(result, r...)
		}
	}
	return result
}

func notPropsFile(file string) bool {
	return strings.ToLower(filepath.Ext(file)) != ".properties"
}

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
		search := strings.Replace(format, "key", string(k), -1)
		re := regexp.MustCompile(search)
		if re.MatchString(text) {
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
