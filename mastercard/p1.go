package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//box := &Box{}
	//locker := sync.Mutex{}
	//waiter := sync.WaitGroup{}
	//for i := 1; i < 100000000; i++ {
	//	waiter.Add(1)
	//	go sumOfSquares(i, box, &locker, &waiter)
	//}
	//waiter.Wait()
	//fmt.Println(box.Value)

	//box := &Box{}
	//locker := sync.Mutex{}
	//waiter := sync.WaitGroup{}
	t1 := time.Now()
	channel := make(chan int, 1)
	sum := 0
	for i := 1; i < 1000000; i++ {
		//waiter.Add(1)
		go sumOfSquares2(i, channel)
	}
	//waiter.Wait()
	for i := 1; i < 1000000; i++ {
		sum += <-channel
	}
	fmt.Println(time.Now().Sub(t1))
	fmt.Println(sum)
}

type Box struct {
	Value int
}

func sumOfSquares(val int, box *Box, locker *sync.Mutex, group *sync.WaitGroup) {
	//locker.Lock()
	box.Value += val * val
	//locker.Unlock()
	group.Done()
}

func sumOfSquares2(val int, channel chan int) {
	channel <- val * val
}
