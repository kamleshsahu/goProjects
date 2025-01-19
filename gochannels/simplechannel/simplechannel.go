package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan []string)
	done := make(chan bool)

	// Goroutine that sends messages and then closes the channel
	go func() {
		fmt.Println("Sending messages...")
		for i := 0; i < 7; i++ {
			channel <- []string{"ping", ""}
			fmt.Println("sent ", i, time.Now())
		}
		close(channel) // Close the channel to signal that no more messages will be sent
	}()

	// Main loop to process messages
	go func() {
		count := 0
		for _ = range channel {
			count++
			fmt.Println("received ", count, time.Now())
		}
		done <- true // Signal that all messages have been processed
	}()

	// Wait for done signal to proceed
	<-done

	// Code that should be executed after all messages are processed
	fmt.Println("All messages processed, pong")
}
