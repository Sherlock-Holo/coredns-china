package plugin

import (
	"strings"
)

type forest struct {
	trees []*tree
}

func (f *forest) Add(domain string) {
	words := reverse(strings.Split(domain, "."))

	for _, t := range f.trees {
		if t.root.value == words[0] {
			t.add(words)
			return
		}
	}

	t := newTree(words[0])
	f.trees = append(f.trees, t)

	t.add(words)
}

func (f *forest) Get(domain string) (match, rootWord string) {
	words := reverse(strings.Split(domain, "."))

	for _, t := range f.trees {
		if t.root.value == words[0] {
			return t.get(words), t.root.value
		}
	}

	return "", ""
}
