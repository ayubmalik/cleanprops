package props

import (
	"path/filepath"
	"sort"
)

type Key string
type Value string
type byKey []Key

func (b byKey) Len() int           { return len(b) }
func (b byKey) Less(i, j int) bool { return b[i] < b[j] }
func (b byKey) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type Props struct {
	FileName string
	props    map[Key]Value
}

func (p Props) SortedKeys() []Key {
	keys := make([]Key, 0, len(p.props))
	for k := range p.props {
		keys = append(keys, k)
	}
	sort.Sort(byKey(keys))
	return keys
}

func LoadProps(file string) Props {
	abs, err := filepath.Abs(file)
	if err != nil {
		panic(err)
	}

	props, err := load(file)
	if err != nil {
		panic(err)
	}

	return Props{abs, props}
}
