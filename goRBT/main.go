package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	tree := redblacktree.NewWithIntComparator()

	tree.Put(1, true)
	tree.Put(2, true)
	tree.Put(5, true)
	tree.Put(8, true)
	tree.Put(11, true)
	tree.Put(16, true)
	tree.Put(21, true)

	recur(tree.Root, 11, 22)

}

func recur(node *redblacktree.Node, left, right int) {
	if node == nil {
		return
	}
	val := node.Key.(int)

	if val > left {
		recur(node.Left, left, right)
	}

	if val >= left && val <= right {
		fmt.Println(val)
	}

	if val < right {
		recur(node.Right, left, right)
	}

}
