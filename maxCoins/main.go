package main

import (
	"fmt"
	"sort"
)

func maximumCoins(_coins [][]int, k int) int64 {
	sort.Slice(_coins, func(i, j int) bool {
		return _coins[i][start] < _coins[j][start]
	})
	coins = _coins
	ps = make([]int, len(coins))

	for i := 0; i < len(coins); i++ {
		ps[i] += coins[i][coin] * (coins[i][end] - coins[i][start] + 1)
		if i > 0 {
			ps[i] += ps[i-1]
		}
	}

	fmt.Println(coins)
	fmt.Println(ps)

	left := 0
	right := 0
	window := 0

	ans := 0

	for ; left < len(coins); left++ {
		window = coins[right][end] - coins[left][start] + 1
		for window < k && right < len(coins) {
			right++
			if right == len(coins) {
				break
			}
			window = coins[right][end] - coins[left][start] + 1
		}

		if right < len(coins) {
			coinscollected := maxCoins(left, right, k)
			ans = max(ans, coinscollected)
		} else if right >= left {
			coinscollected := maxCoins(left, right-1, k)
			ans = max(ans, coinscollected)
		}
	}

	return int64(ans)

}

const (
	coin  = 2
	start = 0
	end   = 1
)

var coins [][]int
var ps []int

func maxCoins(i, j, k int) int {
	totalCoins := ps[j]
	if i > 0 {
		totalCoins -= ps[i-1]
	}
	totalElements := coins[j][end] - coins[i][start] + 1
	extraElements := max(totalElements-k, 0)
	//fmt.Println("extraElements: ", coins[i][start], coins[j][end], extraElements)
	removeFromRight := totalCoins - min(extraElements, coins[j][end]-coins[j][start]+1)*coins[j][coin]

	removeFromLeft := totalCoins - min(extraElements, coins[i][end]-coins[i][start]+1)*coins[i][coin]
	//fmt.Println("total coins", totalCoins, "removeFromRight: ", removeFromRight, "removeFromLeft: ", removeFromLeft)
	return max(removeFromRight, removeFromLeft)
}

func main() {
	//c := [][]int{{8, 10, 1}, {1, 3, 2}, {5, 6, 4}}
	//fmt.Println(maximumCoins(c, 4))

	//c1 := [][]int{{30, 49, 12}}
	//fmt.Println(maximumCoins(c1, 28))

	c2 := [][]int{{10, 18, 6}, {43, 47, 3}}
	fmt.Println(maximumCoins(c2, 22))
}
