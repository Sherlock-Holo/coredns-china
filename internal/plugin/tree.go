package plugin

import (
	"strings"
)

type node struct {
	value    string
	children map[string]*node
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
			children: make(map[string]*node),
		}

		root.children[word] = child
		root = child
	}
}

func (t *tree) get(words []string) (match string) {
	node := t.root
	matchWords := []string{words[0]}
	// ignore eg: com, cn, org, so that we can make the loop clearly
	words = words[1:]

	for _, word := range words {
		if child, ok := node.children[word]; ok {
			matchWords = append(matchWords, word)
			node = child
			continue
		}

		break
	}

	return strings.Join(reverse(matchWords), ".")
}
