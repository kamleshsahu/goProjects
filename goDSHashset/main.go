package main

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/sets/linkedhashset"
)

func main() {
	set := hashset.New()   // empty
	set.Add(1)             // 1
	set.Add(2, 2, 3, 4, 5) // 3, 1, 2, 4, 5 (random order, duplicates ignored)
	set.Remove(4)          // 5, 3, 2, 1 (random order)
	set.Remove(2, 3)       // 1, 5 (random order)
	set.Contains(1)        // true
	set.Contains(1, 5)     // true
	set.Contains(1, 6)     // false
	_ = set.Values()       // []int{5,1} (random order)
	set.Clear()            // empty
	set.Empty()            // true
	set.Size()             // 0

	set1 := linkedhashset.New() // empty
	set1.Add(5)                 // 5
	set1.Add(4, 4, 3, 2, 1)     // 5, 4, 3, 2, 1 (in insertion-order, duplicates ignored)
	set1.Add(4)                 // 5, 4, 3, 2, 1 (duplicates ignored, insertion-order unchanged)
	set1.Remove(4)              // 5, 3, 2, 1 (in insertion-order)
	set1.Remove(2, 3)           // 5, 1 (in insertion-order)
	set1.Contains(1)            // true
	set1.Contains(1, 5)         // true
	set1.Contains(1, 6)         // false
	_ = set1.Values()           // []int{5, 1} (in insertion-order)
	set1.Clear()                // empty
	set1.Empty()                // true
	set1.Size()                 // 0
}
