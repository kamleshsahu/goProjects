package main

import (
	"fmt"
	"time"
)

func main() {
	test()
}

func test() {
	i := 0
	ch := make(chan struct{})
	go func() {
		i++
		close(ch)
		time.Sleep(1000 * time.Microsecond)
		i++
	}()

	<-ch
	fmt.Println(i)
}
