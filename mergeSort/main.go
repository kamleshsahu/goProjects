package main

import (
	"fmt"
	"sort"
)

const (
	VAL   = iota // 0
	COUNT        // 1
	INDEX        // 2
)

func main() {
	//arr := []int{4, 2, 3, 7, 8}
	//arr := []int{5,2,6,1}
	//arr := []int{-1,-1}
	//arr := []int{0, 2, 1}
	arr := []int{26, 78, 27, 100, 33, 67, 90, 23, 66, 5, 38, 7, 35, 23, 52, 22, 83, 51, 98, 69, 81, 32, 78, 28, 94, 13, 2, 97, 3, 76, 99, 51, 9, 21, 84, 66, 65, 36, 100, 41}
	fmt.Println(arr)
	ans := countSmaller(arr)
	fmt.Println(ans)
}

func countSmaller(nums []int) []int {
	copyArr := make([]int, len(nums))
	copy(copyArr, nums)
	sort.Ints(copyArr)
	fmt.Println(copyArr)
	arr := make([][]int, len(nums))

	for i, val := range nums {
		arr[i] = []int{val, 0, i}
	}

	sorted, _ := mergeSort(arr, 0, len(arr))
	fmt.Println(sorted)
	ans := make([]int, len(sorted))
	for i := 0; i < len(sorted); i++ {
		ans[sorted[i][INDEX]] = sorted[i][COUNT]
	}
	return ans
}

func mergeSort(arr [][]int, left, right int) ([][]int, int) {
	if right-left <= 1 {
		return arr[left:right], 0
	}

	mid := left + (right-left)/2
	leftArr, invl := mergeSort(arr, left, mid)
	rightArr, invr := mergeSort(arr, mid, right)

	var merge [][]int
	leftPtr := 0
	rightPtr := 0
	inv := invr + invl
	for i := 0; i < len(leftArr)+len(rightArr); i++ {
		if leftArr[leftPtr][VAL] <= rightArr[rightPtr][VAL] {
			merge = append(merge, leftArr[leftPtr])
			leftArr[leftPtr][COUNT] += rightPtr
			inv += rightPtr
			leftPtr++
		} else {
			merge = append(merge, rightArr[rightPtr])

			rightPtr++
		}
		if leftPtr == len(leftArr) || rightPtr == len(rightArr) {
			break
		}
	}

	for i := leftPtr; i < len(leftArr); i++ {
		merge = append(merge, leftArr[i])
		leftArr[leftPtr][COUNT] += rightPtr
	}
	for i := rightPtr; i < len(rightArr); i++ {
		merge = append(merge, rightArr[i])
	}

	return merge, inv
}
