package main

import "fmt"

func main() {

	fmt.Println(numberOfSubmatrices([][]byte{{'X', 'Y', '.'}, {'Y', '.', '.'}}))
	fmt.Println(numberOfSubmatrices([][]byte{{'.', 'X'}, {'.', 'Y'}}))

}

func numberOfSubmatrices(grid [][]byte) int {

	dp := make([][]*Point, len(grid)+1)

	for i, _ := range dp {
		dp[i] = make([]*Point, len(grid[0])+1)
		for j := 0; j < len(grid[0])+1; j++ {
			dp[i][j] = &Point{0, 0}
		}
	}

	n := len(dp)
	m := len(dp[0])
	count := 0
	//dp[0][0] = &Point{0, 0}
	//dp[0][1] = &Point{0, 0}
	//dp[1][0] = &Point{0, 0}
	//dp[1][1] = &Point{0, 0}
	fmt.Println(dp[0])
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			//if j > 0 {
			dp[i][j].add(dp[i][j-1])
			//}
			//if i > 0 {
			dp[i][j].add(dp[i-1][j])
			//}

			//if i > 0 && j > 0 {
			dp[i][j].sub(dp[i-1][j-1])
			//}

			if grid[i-1][j-1] == 'X' {
				dp[i][j].X++
			} else if grid[i-1][j-1] == 'Y' {
				dp[i][j].Y++
			}

			if dp[i][j].X == 0 || dp[i][j].Y == 0 {
				continue
			}

			if dp[i][j].X == dp[i][j].Y {
				count++
			}
		}

		fmt.Println(dp[i])
	}

	//for i := 1; i < n; i++ {
	//	for j := 1; j < m; j++ {
	//		//fmt.Println(i, j)
	//		for k := 0; k < i; k++ {
	//			for l := 0; l < j; l++ {
	//				if dp[i][j].X == 0 || dp[i][j].Y == 0 {
	//					continue
	//				}
	//				// fmt.Println(dp[i][j], dp[k][l]);
	//
	//				cp := dp[i][j].copy()
	//				fmt.Printf("[%d,%d] - [%d,%d] %s\n", i, j, k, l, cp)
	//				cp.sub(dp[i][l])
	//				cp.sub(dp[k][j])
	//				cp.add(dp[k][l])
	//				fmt.Printf("[%d,%d] - [%d,%d] diff : %s\n", i, j, k, l, cp)
	//				if cp.X > 0 && cp.X == cp.Y {
	//					fmt.Printf("[%d,%d] - [%d,%d]\n", i, j, k, l)
	//					count++
	//				}
	//			}
	//		}
	//	}
	//}

	return count
}

type Point struct {
	X int
	Y int
}

func (p *Point) add(p1 *Point) {
	p.X += p1.X
	p.Y += p1.Y
}

func (p *Point) sub(p1 *Point) {
	p.X -= p1.X
	p.Y -= p1.Y
}

func (p *Point) String() string {
	return fmt.Sprintf("[%d, %d]", p.X, p.Y)
}

func (p *Point) copy() *Point {
	return &Point{p.X, p.Y}
}
