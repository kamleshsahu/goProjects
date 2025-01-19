package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//arr := [3]Student{Student{1, "kamlesh"}, Student{2, "nilesh"}, Student{3, "tikesh"}}
	//arr2 := [3]Student{Student{1, "kamlesh"}, Student{2, "nilesh"}, Student{3, "tikes"}}
	//
	//fmt.Println(arr[1] == arr2[1])

	var x int
	arr := []int{3, 5, 2}
	x, arr = arr[0], arr[1:]
	fmt.Println(x, arr)

	//var a, b, c byte
	//fmt.Scan(&a)
	//fmt.Scan(&b)
	//c = a + b
	//fmt.Println(c)

	var i uint8

	i = 255

	i++

	fmt.Println(i)

	fmt.Println(i)
	fmt.Println()
	s := &Student{Age: 10}
	//test1(s)

	fmt.Println(s)

	var m *Student
	var data = []byte(`{"age": 1}`)
	err := json.Unmarshal(data, &m)
	fmt.Println(err)
	fmt.Println("all good")
}

//func test1(s *Student) {
//	s.Age = 13
//
//	fmt.Println(s)
//	var i int
//	for i < 10 {
//		fmt.Println("kamlesh")
//		i++
//	}
//}

type Student struct {
	Age   int
	Namer Namer `json:"name"`
}

func (s Student) String() string {
	return fmt.Sprint(s.Age)
}

type Namer interface {
	GetName()
}
