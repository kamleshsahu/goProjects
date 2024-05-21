package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	dp = make(map[int][][]int)
	ans := getCombinations(candidates, 0, target)
	return ans
}

var dp map[int][][]int

func getCombinations(candidates []int, i, target int) [][]int {
	var ans [][]int
	if target == 0 {
		ans = append(ans, []int{})
		return ans
	}

	if target < 0 || i >= len(candidates) {
		return ans
	}

	if dp[target] != nil {
		return dp[target]
	}

	//take
	o1 := getCombinations(candidates, i, target-candidates[i])
	//skip
	o2 := getCombinations(candidates, i+1, target)

	for _, set := range o1 {
		set = append(set, candidates[i])
		ans = append(ans, set)
	}

	for _, set := range o2 {
		ans = append(ans, set)
	}
	if dp[target] == nil {
		dp[target] = [][]int{}
	}
	dp[target] = ans
	return ans
}

func main() {
	ans := combinationSum([]int{7, 3, 2}, 18)

	fmt.Println(ans)
}
