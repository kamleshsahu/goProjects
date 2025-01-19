package main

import (
	"fmt"
)

func main() {
	ans := longestSubarray([]int{8, 2, 4, 14, 5, 16, 7, 8}, 4)
	fmt.Println(ans)
}

func longestSubarray(_nums []int, limit int) int {
	minstack := &MinStack{}
	maxstack := &MaxStack{}
	nums = _nums

	ans := 1
	start := 0
	for end := 0; end < len(nums); end++ {
		minstack.Add(end)
		maxstack.Add(end)
		fmt.Println("min stack after :", start, "-", end, minstack, maxstack)

		diff := abs(maxstack.First() - minstack.First())
		for start < end && diff > limit {
			minstack.RemoveFirst(start)
			maxstack.RemoveFirst(start)
			diff = abs(maxstack.First() - minstack.First())
			start++
		}
		if diff <= limit {
			ans = max(ans, end-start+1)
		}
		fmt.Println(end, ": stack after :", start, "-", end, minstack, maxstack)
	}

	return ans
}

type MinStack []int
type MaxStack []int

var nums []int

func (ms *MinStack) Add(id int) {
	stack := *ms
	for len(stack) > 0 && nums[id] < nums[stack[len(stack)-1]] {
		stack = stack[:len(stack)-1]
	}

	stack = append(stack, id)
	*ms = stack
}

func (ms *MaxStack) Add(id int) {
	stack := *ms
	for len(stack) > 0 && nums[id] >= nums[stack[len(stack)-1]] {
		stack = stack[:len(stack)-1]
	}

	stack = append(stack, id)
	*ms = stack
}

func (ms *MaxStack) RemoveFirst(id int) {
	stack := *ms
	if stack[0] <= id {
		stack = stack[1:]
	}
	*ms = stack
}

func (ms *MinStack) RemoveFirst(id int) {
	stack := *ms
	if stack[0] <= id {
		stack = stack[1:]
	}
	*ms = stack
}

func (ms *MaxStack) First() int {
	stack := *ms
	return nums[stack[0]]
}

func (ms *MinStack) First() int {
	stack := *ms
	return nums[stack[0]]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
