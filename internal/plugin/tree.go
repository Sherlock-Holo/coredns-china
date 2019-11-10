package plugin

import (
	"strings"
)

type node struct {
	value    string
	children []*node
}

type tree struct {
	root *node
}

func (t *tree) add(words []string) {
	root := t.root
	// ignore eg: com, cn, org, so that we can make the loop clearly
	words = words[1:]

WORDS:
	for _, word := range words {
		for _, child := range root.children {
			if child.value == word {
				root = child
				continue WORDS
			}
		}
		child := &node{
			value:    word,
			children: nil,
		}
		root.children = append(root.children, child)
		root = child
	}
}

func (t *tree) get(words []string) (match string) {
	node := t.root
	matchWords := []string{words[0]}
	// ignore eg: com, cn, org, so that we can make the loop clearly
	words = words[1:]

WORDS:
	for _, word := range words {
		for _, child := range node.children {
			if child.value == word {
				matchWords = append(matchWords, word)
				node = child
				continue WORDS
			}
		}
		break
	}

	return strings.Join(reverse(matchWords), ".")
}
