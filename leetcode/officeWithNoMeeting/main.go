package main

import "fmt"

var vis [][]bool
var dp [][][]int

func shortestPath(grid [][]int, health int) int {
	m := len(grid)
	n := len(grid[0])

	// Initialize dp array
	dp = make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, health+1)
			for k := 0; k < health+1; k++ {
				dp[i][j][k] = 1_000_000
			}
		}
	}

	// Initialize visited array
	vis = make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	// Start the check from (0, 0)
	ans := check(grid, 0, 0, health)

	if ans >= 1_000_000 {
		return -1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(dp[i][j])
		}
		fmt.Println()
	}

	return ans
}

// Check function to determine if a path is possible
func check(grid [][]int, row, col, hlt int) int {
	if isInvalid(grid, row, col) || hlt < 0 || vis[row][col] {
		return 1_000_000
	}

	if grid[row][col] == 1 && hlt == 0 {
		return 1_000_000
	}

	if row == len(grid)-1 && col == len(grid[0])-1 {
		dp[row][col][hlt-grid[row][col]] = 0
		return 0
	}

	if dp[row][col][hlt] != 1_000_000 {
		return dp[row][col][hlt]
	}

	vis[row][col] = true

	left := check(grid, row, col-1, hlt-grid[row][col])
	right := check(grid, row, col+1, hlt-grid[row][col])
	up := check(grid, row-1, col, hlt-grid[row][col])
	down := check(grid, row+1, col, hlt-grid[row][col])

	vis[row][col] = false

	result := min(min(up, down), min(left, right)) + 1
	dp[row][col][hlt] = result

	return result
}

func isInvalid(grid [][]int, row, col int) bool {
	m := len(grid)
	n := len(grid[0])
	return row >= m || col >= n || row < 0 || col < 0
}

func main() {

	grid := [][]int{{0, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}, {0, 0}, {0, 1}, {0, 1}, {0, 1}, {0, 0}, {1, 0}, {1, 0}, {0, 0}}
	k := 3
	ans := shortestPath(grid, k)
	fmt.Println(ans)
}
