package advanced

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

/* This function is run multiple times as go routines */
func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // When done decrement the wait group counter by 1, indicates done

	fmt.Println("started task", id)
	time.Sleep(time.Duration(id) * time.Second) // Simulate an expensive task
	fmt.Println("finished task", id)
}

/* WaitGroups - to wait for multiple goroutines to finish, we can use a wait group. */
func execute() {
	const NUM_OF_ROUTINES = 4
	var wg sync.WaitGroup   // Create wait group for all the goroutines launched here to finish
	wg.Add(NUM_OF_ROUTINES) // Add counter of how many routines it has to wait for

	for i := 1; i <= NUM_OF_ROUTINES; i++ {
		go task(i, &wg) // Spawn multiple go routines
	}

	wg.Wait() // Wait until counter goes to zero i.e all routines notified as done
	fmt.Println("finished all tasks")
}

func Advanced8() {
	execute()
}

/* Rate limiting - mechanism for controlling resource utilization and maintaining quality of service.
Go elegantly supports rate limiting with goroutines, channels, and tickers.
Below is a basic rate limiter implementation. */
func Advanced9() {
	// Populate the channel to serve all requests
	requests := make(chan string, 10)
	for req := 1; req <= 10; req++ {
		requests <- "request " + strconv.Itoa(req)
	}
	close(requests)

	// Limiter will receive value after every 200ms
	limiter := time.NewTicker(200 * time.Millisecond)
	for req := range requests {
		start := time.Now()

		// By blocking on a receive from the limiter channel before serving each
		// request, we limit ourselves to 1 request every 200 milliseconds.
		<-limiter.C
		fmt.Println("process", req, "	time diff from last request is", time.Since(start))

		// We can observe that the difference between each request is limited to 200ms.
	}
}

/* Atomic Counters */
func Advanced10() {
	var value uint64
	var wg sync.WaitGroup

	// Start 50 goroutines that each increment the counter exactly 1000 times
	for r := 1; r <= 50; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for count := 1; count <= 1000; count++ {
				atomic.AddUint64(&value, 1) // Atomically increment the counter

				// value++
				// Had we used the non-atomic ops++ to increment the counter, we’d likely get a different
				// number every time, changing between runs, because the goroutines would interfere with each other.
				// Moreover, we’d get data race failures when running with the -race flag.
			}
		}()
	}
	wg.Wait()

	// It’s safe to access ops now because we know no other goroutine is writing to it
	fmt.Println("value is", value)
	// Reading atomics safely while they are being updated is also possible, using
	// functions like atomic.LoadUint64.
}
