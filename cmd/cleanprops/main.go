package main

import (
	"fmt"
	"github.com/ayubmalik/cleanprops/internal/pkg/read"
	"os"
	"path/filepath"
)

func main() {
	pwd()
	fmt.Println("cleaning props")
	props := read.Props("/tmp/hello.properties")
	for k, v := range props {
		fmt.Printf("k:'%s', v:'%s'\n", k, v)
	}
}

func pwd() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
}
