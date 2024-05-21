package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/doublylinkedlist"
)

type day struct {
	value int
	next  *day
	prev  *day
}

func main() {
	list := doublylinkedlist.New()

	d1 := day{value: 1}
	d2 := day{value: 6}
	d3 := day{value: 3}
	d4 := day{value: 7}

	list.Add(d1, d2, d3)
	list.d1.next = &d2
	d2.next = &d3
	d3.next = &d4

	d4.next = nil

	curr := &d1
	count := 0
	for curr != nil {
		fmt.Println(curr)
		curr = curr.next
		count++
	}

}
