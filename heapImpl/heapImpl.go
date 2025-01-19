package main

import (
	"container/heap"
	"fmt"
)

type Heap []intP

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	value := x.(int)
	*h = append(*h, value)
}

func (h *Heap) Pop() any {
	myHeap := *h
	last := myHeap[len(myHeap)-1]
	myHeap = myHeap[:len(myHeap)-1]
	*h = myHeap
	return last
}

func main() {
	list := Heap{1, 3, 3, 4, 6, 7, 8}

	heap.Init(&list)

	fmt.Println(heap.Pop(&list))
	heap.Push(&list, 11)
	fmt.Println(heap.Pop(&list))
}
