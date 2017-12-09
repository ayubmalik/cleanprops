package propsclean

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadProps(file string) {
	body, err := ioutil.ReadFile(file)
	check(err)
	fmt.Println(string(body))
}
