package main

import (
	"fmt"
)

// Define an interface for types that can be ordered
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// Max returns the maximum element in a slice of ordered elements
func Max[T Ordered](slice []T) T {
	if len(slice) == 0 {
		panic("empty slice")
	}
	max := slice[0]
	for _, value := range slice[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Max of intSlice:", Max(intSlice))

	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("Max of floatSlice:", Max(floatSlice))

	stringSlice := []string{"apple", "banana", "cherry"}
	fmt.Println("Max of stringSlice:", Max(stringSlice))
}
