package main

import (
	"fmt"
	"math"
)

var dp map[int]map[int]int
var n int
var mod int

func numberWays(hats [][]int) int {
	dp = make(map[int]map[int]int)

	hatToPeople := make([][]int, 41)

	for person, hatList := range hats {
		for _, hat := range hatList {
			if hatToPeople[hat] == nil {
				hatToPeople[hat] = []int{}
			}
			hatToPeople[hat] = append(hatToPeople[hat], person)
		}
	}
	n = (1 << len(hats)) - 1
	mod = int(math.Pow(10, 9)) + 7

	return recur(hatToPeople, 1, 0) % mod
}

func recur(hatToPeople [][]int, idx int, selectedPeoples int) int {
	if selectedPeoples == n {
		return 1
	}
	if idx >= 40 {
		return 0
	}
	key := selectedPeoples
	if dp[key] != nil && dp[key][idx] != 0 {
		return dp[key][idx] % mod
	}

	ans := recur(hatToPeople, idx+1, selectedPeoples) % mod
	for i := 0; i < len(hatToPeople[idx]); i++ {
		if notHas(selectedPeoples, hatToPeople[idx][i]) {
			ans = (ans + recur(hatToPeople, idx+1, selectedPeoples|(1<<hatToPeople[idx][i]))) % mod
		}
	}
	if dp[key] == nil {
		dp[key] = make(map[int]int)
	}
	dp[key][idx] = ans
	return ans % mod
}

func notHas(selectedPeoples int, people int) bool {
	return (selectedPeoples & (1 << people)) == 0
}

func main() {
	hats := [][]int{{1, 3, 5, 10, 12, 13, 14, 15, 16, 18, 19, 20, 21, 27, 34, 35, 38, 39, 40},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40},
		{3, 7, 10, 12, 13, 14, 15, 17, 21, 25, 29, 31, 35, 40},
		{2, 3, 7, 8, 9, 11, 12, 14, 15, 16, 17, 18, 19, 20, 22, 24, 25, 28, 29, 32, 33, 34, 35, 36, 38},
		{6, 12, 17, 20, 22, 26, 28, 30, 31, 32, 34, 35},
		{1, 4, 6, 7, 12, 13, 14, 15, 21, 22, 27, 28, 30, 31, 32, 35, 37, 38, 40},
		{6, 12, 21, 25, 38},
		{1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 34, 35, 36, 37, 38, 39, 40}}
	fmt.Println("Number of ways:", numberWays(hats)) // Output: 842465346
}
