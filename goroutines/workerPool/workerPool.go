package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		//fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Nanosecond)
		//fmt.Println("worker", id, "finished job", j)
		results <- j * j
	}
}

func main() {

	const numJobs = 10000000
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	t1 := time.Now()
	for w := 1; w <= 1000; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	sum := 0
	for a := 1; a <= numJobs; a++ {
		sum += <-results
	}

	fmt.Println(time.Now().Sub(t1))
	fmt.Println(sum)
}
