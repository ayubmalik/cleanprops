package cleanprops

import (
	"os"
	"path/filepath"
)

func listFiles(dir string) []string {
	files := make([]string, 0)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	return files
}
