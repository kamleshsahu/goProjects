package main

import (
	"container/heap"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

type MinHeap []Student

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i].Age < h[j].Age }

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Student))
}

func (h *MinHeap) Pop() interface{} {
	heapObj := *h
	last := heapObj[len(heapObj)-1]
	*h = heapObj[:len(heapObj)-1]
	return last
}

func main() {
	students := MinHeap{}
	students = append(students, Student{Name: "Kamlesh", Age: 27})
	students = append(students, Student{Name: "Nilesh", Age: 18})
	students = append(students, Student{Name: "Tikesh", Age: 22})

	heap.Init(&students)

	fmt.Println(heap.Pop(&students))
	fmt.Println(heap.Pop(&students))
}
