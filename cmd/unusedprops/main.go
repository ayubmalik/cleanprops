package main

import (
	"fmt"
	"github.com/ayubmalik/cleanprops"
	"os"
	"path/filepath"
)

func main() {
	pwd()
	fmt.Println("cleaning props")
	props := cleanprops.LoadProps("testdatahello.properties")
	for k, v := range props.SortedKeys() {
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
