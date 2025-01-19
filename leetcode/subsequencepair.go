package main

import "fmt"

func main() {

	mymap := make(map[int]int)

	val := mymap[5]
	fmt.Println(val)
	val1, found := mymap[5]
	fmt.Println(val1, found)

	//nums := []int{1, 2, 3, 4}
	////k := 8
	//ans := maximumLength(nums, 2)
	//fmt.Println(ans)
}
func maximumLength(nums []int, k int) int {
	reminderIdx = make([][]int, 2*k+1)

	for i, _ := range nums {
		nums[i] %= k
		if reminderIdx[nums[i]] == nil {
			reminderIdx[nums[i]] = make([]int, 0)
		}
		reminderIdx[nums[i]] = append(reminderIdx[nums[i]], i)
	}
	ans := 2
	// fmt.Println(reminderIdx);
	for rem := 0; rem <= 2*k; rem++ {
		for curr := 0; curr <= rem; curr++ {
			if len(reminderIdx[rem-curr]) == 0 {
				continue
			}
			count := 0
			nextIdx := reminderIdx[rem-curr][0]

			for nextIdx < len(nums) {
				newrem := rem - nums[nextIdx]
				nextIdx = getNext(reminderIdx[newrem], nextIdx)
				count++
			}
			ans = max(ans, count)
		}
	}

	return ans
}

var reminderIdx [][]int

func getNext(reminder []int, targetIdx int) int {
	end := len(reminder) - 1
	start := 0
	ans := 10000
	for start <= end {
		mid := start + (end-start)/2
		if reminder[mid] > targetIdx {
			ans = reminder[mid]
			end = mid - 1
		} else {
			start = mid + 1
		}

	}
	return ans
}
