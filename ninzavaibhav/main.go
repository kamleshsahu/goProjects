package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int32, k int32) int64 {
	maxm, sum := int64(math.MinInt64), int64(0)
	start := 0
	count := int32(0)
	for i, num := range nums {
		_num := int64(num)
		if _num > sum+_num {
			start = i
			count = 0
			sum = 0
		}
		count++
		sum += _num

		if count > k {
			count--
			sum -= int64(nums[start])
			start++

			for sum-int64(nums[start]) >= sum {
				count--
				sum -= int64(nums[start])
				start++
			}
		}

		maxm = max(maxm, sum)
	}
	//fmt.Println(maxm)

	return maxm
}

func main() {
	k := int32(5)

	arr := []int32{5, 1, -1, -1, 4, 10, -4}
	//8
	//18
	//13
	//17
	//18
	ans := maxSubArray(arr, k)

	fmt.Println(ans)
}
