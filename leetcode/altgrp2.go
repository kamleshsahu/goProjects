package main

import "fmt"

func numberOfAlternatingGroups(colors []int, k int) int {

	count := 0

	n := len(colors)

	dp := make([]int, n)

	for i, _ := range colors {
		if i == 0 {
			continue
		}
		if colors[i] == colors[i-1] {
			dp[i] = 0
		} else {
			dp[i] = dp[i-1] + 1
		}

	}

	fmt.Println(dp)

	for i := 0; i < n; i++ {
		ans := iscont(colors, dp, i, k, n)
		fmt.Println(i, ": ", ans)
		if ans {
			count++
		}
	}

	return count
}

func iscont(colors []int, dp []int, i, k int, n int) bool {

	if i < k {
		p1 := dp[i] == i
		last := (i - k + 1 + n) % n
		p2 := dp[n-1]-dp[last] == (n-1)-last

		return p1 && p2
	}

	last := i - k
	return (dp[i] - dp[last]) == (last - (i))
}
