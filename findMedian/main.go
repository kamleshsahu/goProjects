package main

import (
	heap "github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)
import "fmt"

type MedianFinder struct {
	left  *heap.Heap
	right *heap.Heap
}

func Constructor() MedianFinder {
	inverseIntComparator := func(a, b interface{}) int {
		return -utils.IntComparator(a, b)
	}
	return MedianFinder{
		left:  heap.NewWith(inverseIntComparator),
		right: heap.NewWithIntComparator(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	rpeek, _ := this.right.Peek()
	if this.right.Empty() || num >= rpeek.(int) {
		this.right.Push(num)
	} else {
		this.left.Push(num)
	}

	if this.right.Size() > this.left.Size() {
		popped, _ := this.right.Pop()
		this.left.Push(popped)
	} else if (this.right.Size() + 1) < this.left.Size() {
		popped, _ := this.left.Pop()
		this.right.Push(popped)
	}
	fmt.Println(num, "\nleft : ", this.left, "\nright : ", this.right)
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Size() == this.right.Size() {
		p1, _ := this.left.Peek()
		p2, _ := this.right.Peek()
		return (float64(p1.(int)) + float64(p2.(int))) / float64(2)
	}
	lpeek, _ := this.left.Peek()

	return float64(lpeek.(int))
}

func main() {
	medianFinder := Constructor()

	medianFinder.AddNum(1)
	medianFinder.AddNum(2)
	medianFinder.FindMedian()
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
