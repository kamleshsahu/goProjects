package main

import (
	"fmt"
	"strings"
)

func main() {
	areSentencesSimilar("a b c", "a c")
}
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	a1 := strings.Split(sentence1, " ")
	a2 := strings.Split(sentence2, " ")

	if len(a2) > len(a1) {
		a1, a2 = a2, a1
	}

	i := 0

	for ; i < len(a2); i++ {
		if a1[i] != a2[i] {
			break
		}
	}
	j := len(a2) - 1
	k := len(a1) - 1
	fmt.Println("before ", i, j)
	for ; j >= 0; j-- {
		if a1[k] != a2[j] {
			break
		}
		k--
	}

	fmt.Println("after ", i, j)

	if i > j {
		return true
	}

	return false

}
