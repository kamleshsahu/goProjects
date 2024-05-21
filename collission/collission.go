package main

import "fmt"

func collission(direction, strength []int) (ans []int) {
	if len(strength) == 0 {
		return ans
	}
	curr := 0

	for curr < len(strength) {

		// collission
		if len(ans) > 0 && direction[ans[len(ans)-1]] > 0 && direction[curr] < 0 {

			if strength[ans[len(ans)-1]] > strength[curr] {
				// last is greater then curr, skip curr
				curr++
			} else if strength[ans[len(ans)-1]] == strength[curr] {
				// both equal strength , remove the last from stack , also skip curr
				ans = ans[:len(ans)-1]
				curr++
			} else {
				// last is smaller then curr, pop last
				ans = ans[:len(ans)-1]
			}

		} else {
			// no collission , add the curr to stack
			ans = append(ans, curr)
			curr++
		}
	}

	return ans
}

func main() {
	dir1 := []int{1, 1, -1, -1, 1}
	strength1 := []int{2, 3, 4, 5, 6}

	fmt.Println(collission(dir1, strength1))
	// expected ans: [2,3,4]

	dir2 := []int{-1, -1, -1}
	strength2 := []int{2, 3, 1}

	fmt.Println(collission(dir2, strength2))
	// expected ans: [0,1,2]

	dir3 := []int{1, 1, 1}
	strength3 := []int{2, 3, 1}

	fmt.Println(collission(dir3, strength3))
	// expected ans: [0,1,2]

	dir4 := []int{1, -1}
	strength4 := []int{2, 2}

	fmt.Println(collission(dir4, strength4))
	// expected ans: []

}
