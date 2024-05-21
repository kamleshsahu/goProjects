package main

import (
	"fmt"
	"github.com/emirpasic/gods/queues/arrayqueue"
)

func main() {
	//redEdges := [][]int{{0, 1}}
	//blueEdges := [][]int{{2, 1}}

	redEdges := [][]int{{0, 1}}
	blueEdges := [][]int{{1, 2}}

	//var redneighbours [5][5]int;
	redneighbours := make([][]int, 5)
	for _, edge := range redEdges {
		if redneighbours[edge[0]] == nil {
			redneighbours[edge[0]] = make([]int, 0)
		}
		redneighbours[edge[0]] = append(redneighbours[edge[0]], edge[1])
	}

	blueneighbours := make([][]int, 5)
	for _, edge := range blueEdges {
		if blueneighbours[edge[0]] == nil {
			blueneighbours[edge[0]] = make([]int, 0)
		}
		blueneighbours[edge[0]] = append(blueneighbours[edge[0]], edge[1])
	}

	ans := shortestPath([][][]int{redneighbours, blueneighbours}, 5)
	fmt.Println(ans)
}

func shortestPath(neighbours [][][]int, n int) []int {
	distance := make([]int, n)
	for i := 0; i < len(distance); i++ {
		distance[i] = 1000
	}

	distance[0] = 0
	queue := arrayqueue.New()

	queue.Enqueue(0)

	curr := 0
	for !queue.Empty() {
		size := queue.Size()
		for i := 0; i < size; i++ {
			current, _ := queue.Dequeue()
			for _, neighbour := range neighbours[curr][current.(int)] {
				distance[neighbour] = min(distance[current.(int)]+1, distance[neighbour])
				queue.Enqueue(neighbour)
			}
		}
		if curr == 0 {
			curr = 1
		} else {
			curr = 0
		}
	}

	curr = 1
	for !queue.Empty() {
		size := queue.Size()
		for i := 0; i < size; i++ {
			current, _ := queue.Dequeue()
			for _, neighbour := range neighbours[curr][current.(int)] {
				distance[neighbour] = min(distance[current.(int)]+1, distance[neighbour])
				queue.Enqueue(neighbour)
			}
		}
		if curr == 0 {
			curr = 1
		} else {
			curr = 0
		}
	}

	return distance
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
