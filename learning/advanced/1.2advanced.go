package advanced

import (
	"fmt"
	"time"
)

func Advanced5() {
	/* Timers represent a single event in the future */
	timer1 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer1.C
		fmt.Println("timer 1 fired")
	}()

	stop := false
	// stop = timer1.Stop() // Timers can be stopped before it fires
	if stop {
		fmt.Println("timer 1 stopped")
	}

	time.Sleep(3 * time.Second)
}

func Advanced6() {
	/* Tickers are for when you want to do something repeatedly at regular intervals */
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case tick := <-ticker.C:
				fmt.Println("ticked at", tick.Format("2006-01-02 15:04:05"))
			}
		}
	}()

	time.Sleep(3 * time.Second) // Delay of 3 secs will tick 3 times after 1 sec
	ticker.Stop()               // Tickers can be stopped like timers

	done <- true
	fmt.Println("ticker stopped")
}

/* These workers will receive work on the jobs channel and send the corresponding results on results */
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		time.Sleep(time.Second) // Simulate an expensive task
		results <- job * 2      // Send the result for this job
		fmt.Println("worker", id, "finished job", job)
	}
}

/* Worker pool using goroutines and channels
Observation: Using worker pool we wil be able to execute 5 jobs (each of 1 sec) within 2 secs instead
of 5 secs. This is done by running 3 worker go routines in parallel. */
func Advanced7() {
	start := time.Now()
	const NUM_OF_WORKERS = 3
	const NUM_OF_JOBS = 5

	jobs := make(chan int, NUM_OF_JOBS)
	results := make(chan int, NUM_OF_JOBS)

	// Create 3 worker routines to finish the jobs, wait for the jobs
	for wr := 1; wr <= NUM_OF_WORKERS; wr++ {
		go worker(wr, jobs, results)
	}

	// Send all the jobs to the channel and close
	for job := 1; job <= NUM_OF_JOBS; job++ {
		jobs <- job
	}
	close(jobs)

	// Collect the results from the channel
	for res := 1; res <= NUM_OF_JOBS; res++ {
		fmt.Println("result =", <-results)
	}

	end := time.Now()
	fmt.Println("jobs executed in", end.Sub(start))
}
