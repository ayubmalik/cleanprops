package main

import (
	"fmt"
	"os"
)

func main() {
	checkArgs()
	for _, file := range os.Args[1:] {
		fmt.Printf("file: '%s'", file)
	}
}

func checkArgs() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: diffprops file1 file2...")
	}
}
