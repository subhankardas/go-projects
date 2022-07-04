package advanced

import (
	"fmt"
	"time"
)

func Advanced1() {
	/* Select - lets you wait on multiple channel operations */
	chncount := 2
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- "msg1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "msg2"
	}()

	for i := 1; i <= chncount; i++ {
		// Use select to await both of these values simultaneously
		select {
		case result1 := <-chan1:
			fmt.Printf("channel 1 > %v\n", result1)
		case result2 := <-chan2:
			fmt.Printf("channel 2 > %v\n", result2)
		}
	}

	/* Timeouts - important for programs that connect to external resources
	or that otherwise need to bound execution time */
	timeout := 3 * time.Second // 3 seconds

	go func() {
		fmt.Println("starting slow I/O process...")
		time.Sleep(10 * time.Second)
		chan1 <- "completed"
	}()

	select {
	case result := <-chan1:
		fmt.Println(result)
	case <-time.After(timeout):
		fmt.Println("error timeout")
	}
}

func Advanced2() {
	/* Non-Blocking Channel Operations - sends and receives on channels are blocking.
	However, we can use select with a default clause to implement non-blocking sends, receives,
	and even non-blocking multi-way selects. */
	msgs := make(chan string)

	select {
	case result := <-msgs:
		fmt.Println("received " + result)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case msgs <- msg:
		fmt.Println("sent " + msg)
	default:
		fmt.Println("no message sent")
	}
}

func Advanced3() {
	/* Closing a channel indicates that no more values will be sent on it */
	ch := make(chan string)

	ch <- "msg 1"
	ch <- "msg 2"

	close(ch)

	ch <- "msg 3" // FATAL ERROR - wont be able to send msg as channel has been closed
	fmt.Println("completed")
}

func Advanced4() {
	/* Range over Channels */
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	/* Since we closed the channel above, loop terminates
	after 2 elements from the closed channel. it’s possible to
	close a non-empty channel but still have the remaining values
	be received. */
	for elem := range queue {
		fmt.Println(elem)
	}
}
