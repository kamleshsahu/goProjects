package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(waysToReachStair(2097132))
}

func waysToReachStair(k int) int {
	state = make(map[string]*int)
	return jump(k, 1, 0, 1)
}

var state map[string]*int

func jump(k, curr, j, last int) int {
	if curr > k+1 || curr < 0 {
		return 0
	}
	hash := string(curr) + "#" + string(j) + "#" + string(last)
	fmt.Println(hash)
	if state[hash] != nil {
		return *state[hash]
	}

	o1 := 0
	// down
	if last > 0 {
		o1 = jump(k, curr-1, j, -1)
	}
	// up
	o2 := jump(k, curr+(int)(math.Pow(2, float64(j))), j+1, 1)

	ans := o1 + o2
	if curr == k {
		ans++
	}

	state[hash] = &ans

	return ans
}
