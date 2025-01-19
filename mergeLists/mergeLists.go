package main

import (
	"container/heap"
	"fmt"
	_ "math/bits"
)

/**
 * Definition for singly-linked list.

 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	list2 := &ListNode{4, &ListNode{4, &ListNode{7, nil}}}
	list3 := &ListNode{2, &ListNode{6, nil}}
	ans := mergeKLists([]*ListNode{list, list2, list3})

	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}

	//list := PQ{{Val: 10}, {Val: 11}, {Val: 5}}
	//lists := PQ{list, list2, list3}

}

type PQ []*ListNode

func (pq *PQ) Push(x any) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *PQ) Pop() any {
	old := *pq
	last := old[len(*pq)-1]
	*pq = old[:len(*pq)-1]
	return last
}

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
}

func (pq PQ) Less(a, b int) bool {
	return pq[a].Val < pq[b].Val
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PQ, 0)

	for _, node := range lists {
		if node != nil {
			pq = append(pq, node)
		}
	}

	heap.Init(&pq)
	ans := &ListNode{-1, nil}
	dummy := ans
	for len(pq) > 0 {
		min := heap.Pop(&pq)
		minl := min.(*ListNode)
		ans.Next = minl
		ans = ans.Next

		if minl.Next != nil {
			heap.Push(&pq, minl.Next)
		}
	}

	return dummy.Next
}
