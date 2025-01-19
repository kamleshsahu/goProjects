package main

import (
	"container/heap"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
)

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool {
	return prefixSum[h[i]] < prefixSum[h[j]]
}
func (P *IntHeap) Swap(i, j int) {
	(*P)[i], (*P)[j] = (*P)[j], (*P)[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IntHeap []int

var prefixSum map[int]int

func getMaxProfit(nums []int32, _k int32) int64 {
	k := int(_k)
	prefixSum = make(map[int]int)
	prefixSum[-1] = 0
	for i := 0; i < len(nums); i++ {
		prefixSum[i] += int(nums[i])
		if i > 0 {
			prefixSum[i] += prefixSum[i-1]
		}
	}

	//fmt.Println(prefixSum)
	ans := math.MinInt64
	minheap := IntHeap{}

	heap.Init(&minheap)
	heap.Push(&minheap, -1)

	for i := 0; i < len(nums); i++ {
		for minheap.Len() > 0 && minheap[0] < (i-k) {
			heap.Pop(&minheap)
		}
		minIdx := minheap[0]
		//fmt.Printf("minIdx :%d : %d, curridx :%d : %d diff : %d\n", minIdx, prefixSum[minIdx], i, prefixSum[i], prefixSum[i]-prefixSum[minIdx])
		ans = max(ans, prefixSum[i]-prefixSum[minIdx])
		heap.Push(&minheap, i)
	}
	return int64(ans)
}

func main() {
	arr := []int32{4, 3, -2, 9, -4, 2, 7}
	k := int32(6)

	//fmt.Println(getMaxProfit(arr, k))
	assert.Equal(nil, getMaxProfit(arr, k), int64(15))
	fmt.Println()
	arr2 := []int32{5, -7, 8, -6, 4, 1, -9, 5}
	k2 := int32(5)

	assert.Equal(nil, getMaxProfit(arr2, k2), int64(8))
	fmt.Println()

	arr3 := []int32{-3, 4, 3, -2, 2, 5}
	k3 := int32(4)

	assert.Equal(nil, getMaxProfit(arr3, k3), int64(8))

	arr4 := []int32{500, -100, 300, 1000, -10000000}
	k4 := int32(4)
	assert.Equal(nil, getMaxProfit(arr4, k4), int64(1700))

	arr5 := []int32{-500, -100, -300, -1000, -10000000}
	k5 := int32(4)
	assert.Equal(nil, getMaxProfit(arr5, k5), int64(-100))

	arr6 := []int32{-500}
	k6 := int32(1)
	assert.Equal(nil, getMaxProfit(arr6, k6), int64(-500))

	arr7 := []int32{-500, -1, 2, 3, -4}
	k7 := int32(4)
	assert.Equal(nil, getMaxProfit(arr7, k7), int64(5))

	k8 := int32(5)

	arr8 := []int32{5, 1, -1, -1, 4, 10, -4}
	fmt.Println(getMaxProfit(arr8, k8))
	assert.Equal(nil, getMaxProfit(arr8, k8), int64(14))
}
