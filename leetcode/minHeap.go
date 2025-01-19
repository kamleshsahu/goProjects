package main

import "container/heap"

func maxSpending(_values [][]int) int64 {
	values = _values
	data := Data{}

	for shopId, items := range values {
		data = append(data, [2]int{shopId, len(items) - 1})
	}

	heap.Init(&data)
	ans := 0
	d := 1
	for len(data) > 0 {
		c := heap.Pop(&data)
		curr := c.([2]int)
		if curr[1] >= 0 {
			heap.Push(&data, [2]int{curr[0], curr[1] - 1})
		}
		ans += values[curr[0]][curr[1]] * d
		d++
	}

	return int64(ans)
}

var values [][]int

type Data [][2]int

func (h *Data) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Data) Pop() any {
	hd := *h
	l := len(hd)
	last := (hd)[l-1]
	hd = hd[:l-1]
	h = &hd
	return last
}

func (h *Data) Len() int {
	return len(*h)
}

func (h *Data) Push(val interface{}) {
	d := val.([2]int)
	*h = append(*h, d)
}

func (h *Data) Less(i, j int) bool {
	hd := *h
	a, b := hd[i], hd[j]
	return values[a[0]][a[1]] < values[b[0]][b[1]]
}
