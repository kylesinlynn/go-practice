package main

import (
	"fmt"
	"time"
)

// Function to simulate a task and send data through a channel
func task1(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Task 1 complete"
}

// Function to simulate another task and send data through a channel
func task2(ch chan<- string) {
	time.Sleep(1 * time.Second)
	ch <- "Task 2 complete"
}

func main() {
	// Create a channel to communicate between goroutines
	ch := make(chan string)

	// Start two tasks as goroutines
	go task1(ch)
	go task2(ch)

	// Receive and print the messages from the tasks
	msg1 := <-ch
	fmt.Println(msg1)

	msg2 := <-ch
	fmt.Println(msg2)

	fmt.Println("All tasks complete")
}
