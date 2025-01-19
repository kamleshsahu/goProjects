package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{1, 3, 4}
	fmt.Println(arr[3:3])

	//number := 19
	////sqrt := int(math.Sqrt(float64(number)))
	////fmt.Printf("The square root of %.2f is %d\n", number, sqrt)
	//dp = make(map[int]int)
	//ans := solve(number)
	//println(ans)
}

var dp map[int]int

func solve(target int) int {
	if target <= 3 {
		return target
	}
	if val, ok := dp[target]; ok {
		return val
	}
	num := getSmallerSquare(target)

	ans := target
	for i := num; i >= 1; i-- {
		ans = min(ans, solve(target-i*i)+1)
	}

	dp[target] = ans
	return ans
}

func min(o1, o2 int) int {
	if o1 <= o2 {
		return o1
	}
	return o2
}

func getSmallerSquare(number int) int {
	sqrt := int(math.Sqrt(float64(number)))
	return sqrt
}
