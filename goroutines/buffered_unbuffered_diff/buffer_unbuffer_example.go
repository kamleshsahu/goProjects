package main

import (
	"fmt"
	"time"
)

func main() {

	bufferedChannel := make(chan string, 1)
	unbufferedChannel := make(chan string)

	fmt.Println("buffered example : ")
	work(bufferedChannel)
	fmt.Println()
	fmt.Println("unbuffered example : ")
	work(unbufferedChannel)
}

func work(messages chan string) {
	t1 := time.Now()

	go func() {
		messages <- "kamlesh: "
		fmt.Println(time.Now().Sub(t1))
		fmt.Println("sent")
	}()

	time.Sleep(5 * time.Second)
	fmt.Println(<-messages)

	fmt.Println(time.Now().Sub(t1))
	fmt.Println("received")
}
