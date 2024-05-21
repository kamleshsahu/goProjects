package main

import (
	"fmt"
	"github.com/emirpasic/gods/queues/circularbuffer"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

func main() {
	queue := circularbuffer.New(3) // empty (max size is 3)
	queue.Enqueue(1)               // 1
	queue.Enqueue(2)               // 1, 2
	queue.Enqueue(3)               // 1, 2, 3
	_ = queue.Values()             // 1, 2, 3
	queue.Enqueue(5)               // 4, 2, 3
	fmt.Println(queue)
	queue.Enqueue(7) // 4, 2, 3
	fmt.Println(queue)

	_, _ = queue.Peek()    // 4,true
	_, _ = queue.Dequeue() // 4, true
	_, _ = queue.Dequeue() // 2, true
	_, _ = queue.Dequeue() // 3, true
	_, _ = queue.Dequeue() // nil, false (nothing to deque)
	queue.Enqueue(1)       // 1
	queue.Clear()          // empty
	queue.Empty()          // true
	_ = queue.Size()       // 0
	main2()
}

// Element is an entry in the priority queue
type Element struct {
	name     string
	priority int
}

// Comparator function (sort by element's priority value in descending order)
func byPriority(a, b interface{}) int {
	priorityA := a.(Element).priority
	priorityB := b.(Element).priority
	return -utils.IntComparator(priorityA, priorityB) // "-" descending order
}

// PriorityQueueExample to demonstrate basic usage of BinaryHeap
func main2() {
	a := Element{name: "a", priority: 1}
	b := Element{name: "b", priority: 2}
	c := Element{name: "c", priority: 3}

	queue := priorityqueue.NewWith(byPriority) // empty
	queue.Enqueue(a)                           // {a 1}
	queue.Enqueue(c)                           // {c 3}, {a 1}
	queue.Enqueue(b)                           // {c 3}, {b 2}, {a 1}
	_ = queue.Values()                         // [{c 3} {b 2} {a 1}]
	_, _ = queue.Peek()                        // {c 3} true
	_, _ = queue.Dequeue()                     // {c 3} true
	_, _ = queue.Dequeue()                     // {b 2} true
	_, _ = queue.Dequeue()                     // {a 1} true
	_, _ = queue.Dequeue()                     // <nil> false (nothing to dequeue)
	queue.Clear()                              // empty
	_ = queue.Empty()                          // true
	_ = queue.Size()                           // 0
}
