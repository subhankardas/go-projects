package advanced

import (
	"fmt"
	"sync"
)

/* Container holds a map of counters; since we want to update it concurrently from multiple goroutines,
we add a Mutex to synchronize access. Note that mutexes must not be copied, so if this struct is passed
around, it should be done by pointer. */
type Container struct {
	mutex    sync.Mutex
	counters map[string]int
}

func (con *Container) increment(name string) {
	con.mutex.Lock() // Lock the mutex before accessing counters, unlock it at the end
	con.counters[name]++

	defer con.mutex.Unlock()
}

/* Mutexes - for more complex state we can use a mutex to safely
access data across multiple goroutines. */
func Advanced11() {
	con := Container{counters: map[string]int{"a": 0, "b": 10}}
	var wg sync.WaitGroup

	incrementCounterBy := func(name string, times int) {
		for time := 1; time <= times; time++ {
			con.increment(name)
		}
		wg.Done()
	}

	// Run several goroutines concurrently to access the same Container,
	// and two of them access the same counter.
	wg.Add(3)
	go incrementCounterBy("a", 1000)
	go incrementCounterBy("b", 1000)
	go incrementCounterBy("a", 1000)

	wg.Wait()
	fmt.Println("counters", con.counters)
}
