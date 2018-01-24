package cleanprops

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const whiteSpace = "\t "

func checkIsProp(a []string) {
	if len(a) == 1 {
		fmt.Fprintf(os.Stderr, "Not valid prop %q", a[0])
		os.Exit(1)
	}
}

func parse(line string) (Key, string) {
	s := regexp.MustCompile(":|\\s*=\\s*|\\s").Split(strings.TrimLeft(line, whiteSpace), 2)
	checkIsProp(s)
	k := Key(strings.Trim(s[0], whiteSpace))
	v := strings.Trim(s[1], whiteSpace)
	return k, v
}

func skip(line string) bool {
	if len(line) == 0 {
		return true
	}

	comment, _ := regexp.MatchString(".*(#|!).*", line)
	if comment {
		return true
	}
	return false
}

func Load(file string) (map[Key]string, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	props := make(map[Key]string)
	s := bufio.NewScanner(f)
	for s.Scan() {
		if skip(s.Text()) {
			continue
		}
		k, v := parse(s.Text())
		props[k] = v
	}
	return props, nil
}
