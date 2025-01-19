package main

import (
	"container/heap"
	"fmt"
)

func main() {
	ans := clearStars("bcb**")

	fmt.Println(ans)

	//ans := clearStars("*");
	//
	//fmt.Println(ans);
}

func clearStars(s string) string {
	pq := make(PQ, 0)
	byteArr := []byte(s)
	heap.Init(&pq)
	toDel := make(map[int]bool)
	for i := 0; i < len(byteArr); i++ {
		if byteArr[i] == '*' {
			if pq.Len() == 0 {
				continue
			}
			idx := heap.Pop(&pq).([]byte)
			toDel[int(idx[1])] = true
			toDel[i] = true
		} else {
			heap.Push(&pq, []byte{byteArr[i], byte(i)})
		}
	}

	ans := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if toDel[i] {
			continue
		}
		ans = append(ans, s[i])
	}
	return string(ans)
}

type PQ [][]byte

func (P *PQ) Len() int {
	return len(*P)
}

func (P *PQ) Less(i, j int) bool {
	if (*P)[i][0] == (*P)[j][0] {
		return (*P)[i][1] > (*P)[j][1]
	}

	return (*P)[i][0] < (*P)[j][0]
}

func (P *PQ) Swap(i, j int) {
	(*P)[i], (*P)[j] = (*P)[j], (*P)[i]
}

func (P *PQ) Push(x any) {
	val := x.([]byte)
	*P = append(*P, val)
}

func (P *PQ) Pop() any {
	pv := *P
	lastIdx := len(pv) - 1
	last := pv[lastIdx]
	pv = pv[:lastIdx]
	*P = pv
	return last
}
