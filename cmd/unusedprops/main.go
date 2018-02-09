package main

import (
	"flag"
	"fmt"
	"github.com/ayubmalik/props"
	"os"
)

var (
	propsFile = flag.String("props", "", ".properties file to read keys from (required)")
	srcDir    = flag.String("srcdir", ".", "source directory containing your code")
	ext       = flag.String("ext", "", "file extensions to include e.g. -ext .java, defaults to all files")
)

func main() {
	flag.Parse()
	if *propsFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("src ", *srcDir)
	p := props.Load(*propsFile)
	matches := props.FindInFiles(p.SortedKeys(), *srcDir, *ext)
	fmt.Println(len(matches))

}
