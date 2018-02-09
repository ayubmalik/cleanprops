package props

import (
	"os"
	"path/filepath"
	"strings"
)

func listFiles(dir string, ext [] string) []string {
	files := make([]string, 0)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && matches(info.Name(), ext) {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func matches(file string, ext []string) bool {
	if len(ext) == 0 {
		return true
	}

	for _, e := range ext {
		if strings.ToLower(filepath.Ext(file)) == strings.ToLower(e) {
			return true
		}
	}
	return false
}
