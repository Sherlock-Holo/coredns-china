package plugin

import (
	"strings"
)

type node struct {
	value    string
	children map[string]*node
}

func (n *node) add(words []string) {
	if len(words) == 0 {
		return
	}

	word := words[0]

	child := n.children[word]
	if child == nil {
		n.children[word] = &node{
			value:    word,
			children: make(map[string]*node),
		}

		child = n.children[word]
	}

	child.add(words[1:])
}

func (n *node) get(words []string) []string {
	if len(words) == 0 {
		return []string{n.value}
	}

	if child, ok := n.children[words[0]]; ok {
		return append(child.get(words[1:]), n.value)
	}

	return []string{n.value}
}

type tree struct {
	root *node
}

func newTree(rootWord string) *tree {
	return &tree{root: &node{
		value:    rootWord,
		children: make(map[string]*node),
	}}
}

func (t *tree) add(words []string) {
	words = words[1:]

	t.root.add(words)
}

func (t *tree) get(words []string) (match string) {
	words = words[1:]

	if matchWords := t.root.get(words); len(matchWords) > 0 {
		return strings.Join(matchWords, ".")
	}

	return ""
}
