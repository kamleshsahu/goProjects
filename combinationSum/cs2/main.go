package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] > candidates[j]
	})

	newarr := make([]int, 0)
	count := make(map[int]int)
	for _, val := range candidates {
		if len(newarr) > 0 && newarr[len(newarr)-1] == val {
			count[val]++
			continue
		}
		count[val] = 1
		newarr = append(newarr, val)
	}
	ans := getCombinations(newarr, count, 0, target)
	return ans
}

var dp map[int]map[int][][]int

func getCombinations(candidates []int, count map[int]int, i, target int) [][]int {
	var ans [][]int
	if target == 0 {
		ans = append(ans, []int{})
		return ans
	}

	if target < 0 || i >= len(candidates) {
		return ans
	}

	//take
	if count[candidates[i]] > 0 {
		count[candidates[i]]--
		o1 := getCombinations(candidates, count, i, target-candidates[i])
		count[candidates[i]]++
		for _, set := range o1 {
			set = append(set, candidates[i])
			ans = append(ans, set)
		}
	}
	//skip
	o2 := getCombinations(candidates, count, i+1, target)

	for _, set := range o2 {
		ans = append(ans, set)
	}

	return ans
}

func main() {
	ans := combinationSum2([]int{1}, 1)
	fmt.Println(ans)
}
