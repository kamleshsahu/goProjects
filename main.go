package main

import (
	"awesomeProject/lld/elevator"
	"fmt"
)

func main() {
	//impl.Splitwise()

	//d := time.Duration(6000000000)
	//
	//fmt.Println(reflect.TypeOf(d))

	//fileSystem.Run()

	//calculator.Run()

	//minWindow("a", "aa")
	elevator.Runner()
}

func minWindow(s string, t string) string {

	tmap := make(map[int32]int)

	for _, val := range t {
		tmap[int32(val)]++
	}

	smap := make(map[int32]int)
	fmt.Println(tmap)
	i := 0
	j := 0
	minAns := 100001
	ans := ""
	haveCount := 0
	for i < len(s) {
		fmt.Println(haveCount, i, j, smap)
		if i < j && haveCount == len(tmap) {
			fmt.Println("inside i:", i)
			if j-i < minAns {
				ans = s[i:j]
				minAns = min(minAns, j-i)
			}
			smap[int32(s[i])]--
			if smap[int32(s[i])] < tmap[int32(s[i])] {
				haveCount--
			}
			i++
		} else if j < len(s) {
			fmt.Println("inside j:", j, smap)
			smap[int32(s[j])]++
			if smap[int32(s[j])] == tmap[int32(s[j])] {
				fmt.Println("inside havecount", haveCount, smap[int32(j)], smap[int32(j)])
				haveCount++
			}
			fmt.Println("after inside j:", j, smap, haveCount)
			j++
		} else {
			break
		}
	}

	return ans

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//func comp(smap, tmap map[int32]int) bool {
//
//	for key, val := range tmap {
//		if smap[key] < val {
//			return false
//		}
//	}
//	return true
//}
