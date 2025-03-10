package main

import "fmt"

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
		fmt.Println(arr)
	}

	swap(arr, i+1, high)
	return i + 1
}

func quickSort(arr []int, low, high int) {

	if low < high {
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {

	arr := []int{11, 14, 5, 1, 6}

	quickSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}
