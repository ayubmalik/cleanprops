package main

import (
	"flag"
	"fmt"
	"github.com/ayubmalik/props"
	"os"
	"strings"
)

var (
	propsFile  = flag.String("props", "", ".properties file to read keys from (required)")
	src        = flag.String("src", ".", "source directory containing your code")
	ext        = flag.String("ext", "", "file extensions to include, e.g. -ext .java or -ext .java,.xml (default all files)")
	extensions []string
)

func main() {
	flag.Parse()
	if *propsFile == "" {
		fmt.Printf("Usage: %s <options>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *ext != "" {
		extensions = strings.Split(*ext, ",")
	}

	fmt.Printf("%-20s %s\n", "Props file", *propsFile)
	fmt.Printf("%-20s %s\n", "Source dir", *src)
	fmt.Printf("%-20s %s\n", "Extensions", extensions)
	p := props.Load(*propsFile)
	matches := props.FindInFiles(p.SortedKeys(), *src, extensions)
	fmt.Printf("%-20s %d\n", "Matches", len(matches))

	unused := props.Diff(p.SortedKeys(), matches)
	fmt.Printf("%-20s %d\n", "Unused", len(unused))
	for i, k := range unused {
		fmt.Println(i, k)
	}
}
