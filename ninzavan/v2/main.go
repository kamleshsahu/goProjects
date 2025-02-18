package main

import (
	"fmt"
	"github.com/go-playground/assert/v2"
)

// We'll use the deque to store indices (of prefixSums) with the property that
// prefixSums[deque[0]] is the minimum value in the current window.
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
	//fmt.Println(getMaxProfit(arr8, k8))
	assert.Equal(nil, getMaxProfit(arr8, k8), int64(14))
}

func getMaxProfit(arr []int32, k int32) int64 {
	d := int(k)
	n := len(arr)

	prefixSums := make([]int32, n+1)
	for i := 0; i < n; i++ {
		prefixSums[i+1] = prefixSums[i] + arr[i]
	}

	if d <= 0 {
		return -1
	}

	ans := int32(-1 << 31)

	deque := make([]int, 0, n+1)
	deque = append(deque, 0)

	for j := 1; j <= n; j++ {
		if len(deque) > 0 && deque[0] < j-d {
			deque = deque[1:]
		}

		currentSum := prefixSums[j] - prefixSums[deque[0]]
		ans = max(ans, currentSum)

		for len(deque) > 0 && prefixSums[deque[len(deque)-1]] >= prefixSums[j] {
			deque = deque[:len(deque)-1]
		}
		// Append the current index j.
		deque = append(deque, j)
	}

	return int64(ans)
}
