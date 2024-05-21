package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

var seq *linkedliststack.Stack

func findOrder(numCourses int, prerequisites [][]int) []int {
	visiting := make(map[int]bool)
	visited := make(map[int]bool)

	n := numCourses
	matrix := make([][]int, n)
	dependencyCount := make([]int, n)
	for _, prereq := range prerequisites {
		dependent, dependency := prereq[0], prereq[1]
		if matrix[dependency] == nil {
			matrix[dependency] = make([]int, n)
		}

		matrix[dependency][dependent] = 1
		dependencyCount[dependent]++
	}

	seq = linkedliststack.New()
	for i, _ := range matrix {
		dfs(matrix, visiting, visited, i)
	}
	temp := make([]int, 0)

	iter := seq.Iterator()

	for iter.Next() {
		temp = append(temp, iter.Value().(int))
	}

	return temp
}

func dfs(matrix [][]int, visiting, visited map[int]bool, dependency int) bool {
	if visited[dependency] {
		return true
	}

	if visiting[dependency] {
		return false
	}
	visiting[dependency] = true
	ans := true
	for dependent, value := range matrix[dependency] {
		if value != 0 {
			ans = ans && dfs(matrix, visiting, visited, dependent)
		}
	}
	delete(visiting, dependency)
	if !ans {
		return ans
	}
	visited[dependency] = true
	seq.Push(dependency)
	return ans
}

func main() {

	numCourses := 3
	prerequisites := [][]int{{0, 2}, {1, 2}, {2, 0}}
	ans := findOrder(numCourses, prerequisites)
	fmt.Println(ans)
}
