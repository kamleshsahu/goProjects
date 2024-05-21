package main

import "fmt"

func main() {
	node := &Node{key: 1, value: 1}

	node2 := &Node{key: 2, value: 2}

	node.addFirst(node2)
	node3 := &Node{key: 3, value: 3}
	node.addFirst(node3)

	node.DeleteTail()
	fmt.Println(node)
}

type Node struct {
	prev, next *Node
	key, value int
}

func (this *Node) addFirst(head *Node) {
	head.prev = this
	this.next = head
}

func (this *Node) DeleteTail() (key int) {
	temp := this
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	if temp.prev != nil {
		temp.prev = nil
	}
	return temp.key
}
