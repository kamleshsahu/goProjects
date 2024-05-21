package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func highFive(items [][]int) [][]int {

	var myMap [1001]*IntHeap

	for _, student := range items {
		h := myMap[student[0]]

		if h == nil {
			h = &IntHeap{}
			heap.Init(h)
		}
		heap.Push(h, student[1])
		if h.Len() > 1 {
			// fmt.Println(heap.Pop(h));
			heap.Pop(h)
		}
		myMap[student[0]] = h
	}

	ans := make([][]int, 0, len(items))
	for i, h := range myMap {
		if h == nil {
			continue
		}
		sum := 0
		for h.Len() > 0 {

			sum += heap.Pop(h).(int)
		}
		stud := []int{i, sum}
		ans = append(ans, stud)
	}
	return ans
}

func main() {
	data := [][]int{{1, 100}, {1, 70}, {1, 50}, {1, 10}}
	ans := highFive(data)
	fmt.Println(ans)
}
