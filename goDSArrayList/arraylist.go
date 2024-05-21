package main

import (
	"container/list"
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/utils"
)

type Student struct {
	name string
	age  int
}

func StudentComparator(a, b interface{}) int {
	aAsserted := a.(Student)
	bAsserted := b.(Student)
	switch {
	case aAsserted.name > bAsserted.name:
		return 1
	case aAsserted.name < bAsserted.name:
		return -1
	default:
		return 0
	}
}

func main() {
	l5 := list.New()
	list := arraylist.New()
	list.Add("d", "a")                // ["a"]
	list.Add("c", "b")                // ["a","c","b"]
	list.Sort(utils.StringComparator) // ["a","b","c"]

	fmt.Println(list)
	s1 := Student{"kamlesh", 27}
	s2 := Student{"tikesh", 24}
	s3 := Student{"nilesh", 20}

	list2 := arraylist.New()
	UNUSED(s1, s2, s3)
	list2.Sort(StudentComparator)
	fmt.Println(list2)

	it := list2.Iterator()
	//it := list.Iterator()
	for it.Next() {
		index, value := it.Index(), it.Value()
		fmt.Println(index, value)
	}
	_, _ = list.Get(0)                    // "a",true
	_, _ = list.Get(100)                  // nil,false
	_ = list.Contains("a", "b", "c")      // true
	_ = list.Contains("a", "b", "c", "d") // false
	list.Swap(0, 1)                       // ["b","a",c"]
	list.Remove(2)                        // ["b","a"]
	list.Remove(1)                        // ["b"]
	list.Remove(0)                        // []
	list.Remove(0)                        // [] (ignored)
	_ = list.Empty()                      // true
	_ = list.Size()                       // 0
	list.Add("a")                         // ["a"]
	list.Clear()                          // []
	list.Insert(0, "b")                   // ["b"]
	list.Insert(0, "a")                   // ["a","b"]
}

func UNUSED(x ...interface{}) {}
