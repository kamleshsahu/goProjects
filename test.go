package main

import (
	"fmt"
	"sort"
)

func maxScore(grid [][]int) int {
	dp = make(map[int]map[int]*int)
	arr := make([][2]int, 0)

	for i, row := range grid {
		for _, col := range row {
			val := [2]int{i, col}
			arr = append(arr, val)
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] < arr[j][0]
		}
		return arr[i][1] < arr[j][1]
	})

	return helper(arr, 0, 0)
}

var dp map[int]map[int]*int

func helper(arr [][2]int, visited int, idx int) int {
	if idx >= len(arr) {
		return 0
	}

	if dp[visited] != nil && dp[visited][idx] != nil {
		return *dp[visited][idx]
	}

	mv := helper(arr, visited, idx+1)
	row := arr[idx][0]

	if ((visited >> row) & 1) == 0 {
		// Skip duplicates
		nextIdx := idx + 1
		for nextIdx < len(arr) && arr[nextIdx][1] == arr[idx][1] {
			nextIdx++
		}

		for visited > 0 && idx < len(arr) && arr[idx][1] == arr[idx-1][1] {
			idx++
		}
		if idx < len(arr) {
			mv = max(mv, helper(arr, visited|(1<<arr[idx][0]), idx+1)+arr[idx][1])
		}
	}

	if dp[visited] == nil {
		dp[visited] = make(map[int]*int)
	}
	dp[visited][idx] = &mv

	return mv
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{
		{8, 7},
		{8, 6},
	}

	fmt.Println("Max Score:", maxScore(grid))
}
