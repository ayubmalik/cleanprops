package main

import (
	"flag"
	"fmt"
	"github.com/ayubmalik/props"
	"os"
)

var (
	propsFile = flag.String("props", "", ".properties file to read keys from (required)")
	src       = flag.String("src", ".", "source directory containing your code")
	ext       = flag.String("ext", "", "file extensions to include e.g. -ext .java, defaults to all files")
)

func main() {
	flag.Parse()
	if *propsFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	var extensions []string
	if *ext != "" {
		extensions = []string{*ext}
	}

	fmt.Println("Props file ", *propsFile)
	fmt.Println("Source dir ", *src)
	fmt.Println("Extensions ", extensions)
	p := props.Load(*propsFile)
	matches := props.FindInFiles(p.SortedKeys(), *src, extensions)
	fmt.Println("Matches ", len(matches))

	unused := props.Diff(p.SortedKeys(), matches)
	fmt.Println("Unused ", len(unused))
	for i, k := range unused {
		fmt.Println(i, k)
	}
}
