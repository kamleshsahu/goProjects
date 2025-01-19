package main

import "fmt"

func main() {

	ans := findLength([]int{1, 1, 0, 0, 1, 1}, []int{0, 0})

	fmt.Println(ans)
}
func findLength(nums1 []int, nums2 []int) int {
	mod := 1_000_000_007
	hashes := make(map[int]map[int]map[int]bool)

	size := 102

	for i := 0; i < len(nums1); i++ {
		hash := 0
		for j := i; j < len(nums1); j++ {

			hash = (hash*size%mod + nums1[j]) % mod
			size := j - i + 1

			if hashes[size] == nil {
				hashes[size] = make(map[int]map[int]bool)
			}
			if hashes[size][hash] == nil {
				hashes[size][hash] = make(map[int]bool)
			}
			hashes[size][hash][i] = true
		}
	}

	hashes2 := make(map[int]map[int]map[int]bool)

	for i := 0; i < len(nums2); i++ {
		hash := 0
		for j := i; j < len(nums2); j++ {
			hash = (hash*size%mod + nums2[j]) % mod
			size := j - i + 1

			if hashes2[size] == nil {
				hashes2[size] = make(map[int]map[int]bool)
			}
			if hashes2[size][hash] == nil {
				hashes2[size][hash] = make(map[int]bool)
			}
			hashes2[size][hash][i] = true
		}
	}

	low := 1
	high := min(len(nums1), len(nums2))

	mx := 0

	for low <= high {
		guessedSize := low + (high-low)/2
		fmt.Println("hash for size:", guessedSize, hashes[guessedSize], hashes2[guessedSize])
		found := false
		if hashes[guessedSize] != nil && hashes2[guessedSize] != nil {
			for hash, idxs := range hashes[guessedSize] {
				if _, exist := hashes2[guessedSize][hash]; exist {
					for idxi, _ := range idxs {
						for idxj, _ := range hashes2[guessedSize][hash] {
							if isequal(nums1, nums2, idxi, idxj, guessedSize) {
								mx = guessedSize
								found = true
								break
							}
						}
					}
				}
				if found {
					break
				}
			}
			if found {
				low = guessedSize + 1
			}
		}
		if !found {
			high = guessedSize - 1
		}
	}

	return mx
}

func isequal(nums1, nums2 []int, i, j, size int) bool {

	for k, l := i, j; k < i+size; k++ {
		if nums1[k] != nums2[l] {
			return false
		}
		l++
	}

	return true
}
