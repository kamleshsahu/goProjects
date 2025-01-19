package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	ans := maxTotalReward([]int{1, 6, 4, 3, 2})

	fmt.Println(ans)
}

func maxTotalReward(rewardValues []int) int {

	sort.Ints(rewardValues)

	slices.Reverse(rewardValues)

	ru := make([]int, 0)

	for _, reward := range rewardValues {
		if len(ru) > 0 && ru[len(ru)-1] == reward {
			continue
		}
		ru = append(ru, reward)
	}

	dp = make(map[int]*[2001]*int)
	return maximise(ru, 0, 0)
}

var dp map[int]*[2001]*int

func maximise(rewards []int, i int, cr int) int {
	if i >= len(rewards) {
		return cr
	}

	nextIdx := next(rewards, i, cr)

	if nextIdx >= len(rewards) {
		return cr
	}

	if dp[cr] != nil && dp[cr][i] != nil {
		return *dp[cr][i]
	}

	ans := 0
	for j := nextIdx; j < len(rewards); j++ {
		ans = max(ans, maximise(rewards, j, cr+rewards[j]))
	}

	if dp[cr] == nil {
		dp[cr] = &[2001]*int{}
	}

	dp[cr][i] = &ans

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func next(rewards []int, i, cr int) int {

	low := i
	high := len(rewards) - 1
	ans := len(rewards)
	for low <= high {
		mid := low + (high-low)/2.0

		if rewards[mid] >= cr {
			high = mid - 1
		} else {
			ans = mid
			low = mid + 1
		}

	}

	return ans
}
