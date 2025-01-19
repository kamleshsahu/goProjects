package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Point represents a point in the grid with time, row, and column
type Point struct {
	time, row, col int
}

// PriorityQueue implements a min-heap for Point
type PriorityQueue []Point

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].time < pq[j].time }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Point)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// minTimeToReach finds the minimum time to reach the bottom-right corner
func minTimeToReach(moveTime [][]int) int {
	n, m := len(moveTime), len(moveTime[0])
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// Minimum time to reach each cell
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, m)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt64
		}
	}
	dist[0][0] = 0

	// Priority queue for BFS
	pq := &PriorityQueue{}
	heap.Push(pq, Point{time: 0, row: 0, col: 0})

	for pq.Len() > 0 {
		p := heap.Pop(pq).(Point)
		t, row, col := p.time, p.row, p.col

		// If we've reached the destination
		if row == n-1 && col == m-1 {
			return t
		}

		// Explore all four directions
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow >= 0 && newRow < n && newCol >= 0 && newCol < m {
				// Calculate the next possible time to move to this cell
				nextTime := max(t, moveTime[newRow][newCol]) + 1
				if nextTime < dist[newRow][newCol] {
					dist[newRow][newCol] = nextTime
					heap.Push(pq, Point{time: nextTime, row: newRow, col: newCol})
				}
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(minTimeToReach([][]int{{0, 4}, {4, 4}}))
}
