package main

import (
	"fmt"
	"sync"
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

func main() {
	evenChan := make(chan int)
	oddChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go sendNumbers(evenChan, oddChan, &wg)
	go printNumbers(evenChan, oddChan, &wg)

	wg.Wait()
}
