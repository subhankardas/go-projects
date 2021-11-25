package main

import (
	"fmt"
	"strconv"
	"time"
)

func routine(rname string) {
	for run := 1; run <= 5; run++ {
		fmt.Print("running [", rname, "] for ", run, "\n")
	}
}

/*
 A Goroutine is a function or method which executes independently
 and simultaneously in connection with any other Goroutines present
 in your program.

 Goroutines are cheaper than threads.
 Goroutine are stored in the stack and the size of the stack
 can grow and shrink according to the requirement of the program.
 But in threads, the size of the stack is fixed.
*/
func main() {
	routine("simple function") // Simple synchronous function call

	go routine("go routine") // Concurrent new goroutine call

	go func(name string) { // Goroutine anonymous function
		fmt.Print("running [", name, "]\n")
	}("anonymous routine")

	time.Sleep(time.Second) // Wait for goroutines to finish

	/* Channels - pipes that connect concurrent goroutines */
	messages := make(chan string)

	// Producer goroutine
	go func(name string) {
		for run := 1; run <= 3; run++ {
			messages <- "msg" + strconv.Itoa(run) // Send a message to the channel
		}
	}("produce")

	// Consumer goroutine
	go func(name string) {
		for run := 1; run <= 3; run++ {
			msg := <-messages
			fmt.Println("received message is", msg) // Receive the message from the channel
		}
	}("consume")

	time.Sleep(time.Second)

	/* Channel buffering - send multiple values into the channel
	without a corresponding concurrent receive */
	messagebuff := make(chan string, 2)

	messagebuff <- "some message"
	messagebuff <- "more message"

	fmt.Println(<-messagebuff)
	fmt.Println(<-messagebuff)

	/* Channel synchronization */
	worker := func(done chan bool) {
		fmt.Println("working...")
		time.Sleep(time.Second)
		fmt.Println("done")

		done <- true // Notify other goroutines using boolean channel
	}

	done := make(chan bool, 1)
	go worker(done) // Start routine with notify channel

	/* Wait until we receive from done channel.
	Without the <-done line, program would exit even before
	the worker execution started */
	fmt.Println("resuming", <-done)

	/* Channel directions */
	send := func(pings chan<- string, msg string) {
		pings <- msg // Pings defined only to receive values (chan<- type)
		// ms := <-pings // COMPILE-TIME ERROR, cannot receive from send only
	}

	ping := make(chan string, 1)
	send(ping, "message")

	fmt.Println("completed")
}
