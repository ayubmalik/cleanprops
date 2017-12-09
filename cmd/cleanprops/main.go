package main

import (
	"fmt"
	"github.com/ayubmalik/cleanprops"
)

func main() {
	fmt.Println("cleaning props")
	props := propsclean.ReadProps("/tmp/hello.props")
	for k, v := range props {
		fmt.Printf("k:'%s', v:'%s'\n", k, v)
	}
}
