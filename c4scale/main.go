package main

import (
	"fmt"
	"sync"
)

// 3 list with radnomised number
var (
	GroupA = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	GroupB = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	GroupC = []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
)

// input will be a number
// output should be index of number and group
// numbers are unsorted
// 100 elements
var groupMapA, groupMapB, groupMapC map[int]int
var groupMaps []map[int]int

func main() {

	groupMaps = make([]map[int]int, 3)
	groupIdToName = []string{"GroupA", "GroupB", "GroupC"}
	go initMap(0, GroupA)
	go initMap(1, GroupB)
	go initMap(2, GroupC)

	fmt.Println(groupMaps)
	fmt.Println(getIndexAndGroup(1))
	fmt.Println(getIndexAndGroup(29))
	fmt.Println(getIndexAndGroup(16))
	fmt.Println(getIndexAndGroup(101))

}

func initMap(idx int, data []int) {
	if groupMaps[idx] == nil {
		groupMaps[idx] = make(map[int]int)
	}
	for i, num := range data {
		groupMaps[idx][num] = i
	}
}

type Response struct {
	Index int
	Value int
	Group string
}

var groupIdToName []string

func getIndexAndGroup(number int) (*Response, error) {
	waitGroup := sync.WaitGroup{}
	var ans *Response
	waitGroup.Add(len(groupMaps))
	for i := 0; i < len(groupMaps); i++ {
		go func() {
			idx, found := groupMaps[i][number]
			if found {
				ans = &Response{Index: idx, Value: number, Group: groupIdToName[i]}
			}
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	if ans != nil {
		return ans, nil
	}

	return nil, fmt.Errorf("number not found in any group")
}
