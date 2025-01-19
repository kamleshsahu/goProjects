package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func sendNumbers(evenChan, oddChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(evenChan)
	defer close(oddChan)

	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			evenChan <- i
		} else {
			oddChan <- i
		}
	}
}

func printNumbers(evenChan, oddChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		select {
		case even := <-evenChan:
			fmt.Println("Even:", even)
		case odd := <-oddChan:
			fmt.Println("Odd:", odd)
		}
	}
}

//func main() {
//	evenChan := make(chan int)
//	oddChan := make(chan int)
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	go sendNumbers(evenChan, oddChan, &wg)
//	go printNumbers(evenChan, oddChan, &wg)
//
//	wg.Wait()
//}
//
//package main
//
//import "fmt"

//func main() {
//
//	messages := make(chan string)
//
//	go func() {
//		messages <- "ping"
//
//	}()
//
//	go func() {
//		messages <- "pong"
//	}()
//
//	go func() {
//		messages <- "p"
//	}()
//
//	msg := <-messages
//	fmt.Println(msg)
//
//	msg2 := <-messages
//	fmt.Println(msg2)
//
//	//msg3 := <-messages
//	//fmt.Println(msg3)
//	//msg4 := <-messages
//	//fmt.Println(msg4)
//}

func main() {

	messages := make(chan string, 1)
	t1 := time.Now()

	go func() {
		messages <- "kamlesh: " + strconv.Itoa(1)
		fmt.Println(time.Now().Sub(t1))
		fmt.Println("sent")
	}()

	time.Sleep(5 * time.Second)
	fmt.Println(<-messages)

	fmt.Println(time.Now().Sub(t1))
	fmt.Println("received")
}
