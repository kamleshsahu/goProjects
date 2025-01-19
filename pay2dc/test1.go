package main

func main() {
	solution([]int{12}, []int{123})
}

func solution(firstArray []int, secondArray []int) int {

	trie := &Trie{
		Value:    0,
		Children: make([]*Trie, 10),
	}

	for _, val := range firstArray {
		add(trie, val)
	}

	ans := 0

	for _, val := range secondArray {
		ans = max(ans, search(trie, val))
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type Trie struct {
	Value    int
	Children []*Trie
}

func add(trie *Trie, number int) {

	arr := reverse(number)

	for i := len(arr) - 1; i >= 0; i-- {
		val := arr[i]
		if trie.Children[val] == nil {
			trie.Children[val] = &Trie{
				Value:    val,
				Children: make([]*Trie, 10),
			}
		}
		trie = trie.Children[val]
	}
}

func reverse(number int) []int {
	arr := make([]int, 0)
	for number > 0 {
		rem := number % 10
		arr = append(arr, rem)
		number = number / 10
	}
	return arr
}

func search(trie *Trie, number int) int {
	arr := reverse(number)
	count := 0
	for i := len(arr) - 1; i >= 0; i-- {
		val := arr[i]
		if trie.Children[val] == nil {
			return count
		}
		count++
		trie = trie.Children[val]
	}
	return count
}

//firstArray = [25, 288, 2655, 54546, 54, 555] and secondArray = [2, 255, 266, 244, 26, 5, 54547], the output should be solution(firstArray, secondArray) = 4.
//firstArray = [25, 288, 2655, 544, 54, 555] and secondArray = [2, 255, 266, 244, 26, 5, 5444444]
//firstArray = [817, 99] and secondArray = [1999, 1909], the output should be solution(firstArray, secondArray) = 0.
