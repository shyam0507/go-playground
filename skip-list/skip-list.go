package main

import (
	"fmt"
	"math/rand/v2"
)

const MAX_LEVELS = 3

type Node struct {
	key   int
	value string
	next  []*Node
}

type SkipList struct {
	// head is a sentinel: it is not a user-visible entry and always has every
	// possible level. This lets every insertion use the same linking logic.
	head   *Node
	levels int // highest currently populated level, in the range [1, MAX_LEVELS]
}

func NewList() *SkipList {
	return &SkipList{
		head:   &Node{next: make([]*Node, MAX_LEVELS)},
		levels: 1,
	}
}

// function to decide upto which level key should be set
func levelGenerator() int {
	level := 1

	for level < MAX_LEVELS && rand.N(2) == 0 {
		level++
	}
	return level
}

func (l *SkipList) add(key int, val string) {
	// update[level] is the node immediately before key at that level.
	update := make([]*Node, MAX_LEVELS)
	current := l.head
	for i := l.levels - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].key < key {
			current = current.next[i]
		}
		update[i] = current
	}

	// Keep one entry per key; adding an existing key replaces its value.
	if next := update[0].next[0]; next != nil && next.key == key {
		next.value = val
		return
	}

	level := levelGenerator()
	if level > l.levels {
		for i := l.levels; i < level; i++ {
			update[i] = l.head
		}
		l.levels = level
	}

	node := &Node{key: key, value: val, next: make([]*Node, level)}
	for i := 0; i < level; i++ {
		node.next[i] = update[i].next[i]
		update[i].next[i] = node
	}
}

func (l *SkipList) search(key int) bool {
	current := l.head
	for i := l.levels - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].key < key {
			current = current.next[i]
		}
	}

	current = current.next[0]
	return current != nil && current.key == key
}

func main() {
	data := []int{5, 7, 9, 3, 21}

	l := NewList()
	for _, v := range data {
		l.add(v, fmt.Sprintf("Hello %d", v))
	}

	fmt.Println(l.search(21))
	fmt.Println(l.search(20))

}
