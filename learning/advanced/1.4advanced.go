package advanced

import (
	"fmt"
	"os"
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

/* A panic typically means something went unexpectedly wrong. Mostly we use it to
fail fast on errors that shouldn’t occur during normal operation, or that we aren’t
prepared to handle gracefully. */
func Advanced12() {
	_, err := os.Create("/tmp/file")
	if err != nil {
		// Abort if a function returns an error value that we don’t want to handle.
		panic("unable to create file")
	}
}

/* Defer is used to ensure that a function call is performed later in a program’s
execution, usually for purposes of cleanup. */
func Advanced13() {
	file := createFile()
	writeFile(file)

	// This will be executed at the end of the enclosing function, after writeFile has finished
	defer closeFile(file)
}

func createFile() *os.File {
	f, err := os.Create("file.txt")
	if err != nil {
		panic("unable to create file")
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Fprintln(f, "file data")
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		panic("unable to close file")
	}
}

/* Recover - from a panic, by using the recover built-in function. A recover can stop a
panic from aborting the program and let it continue with execution instead. */
func doPanic() {
	panic("some error")
}

func Advanced14() {
	defer func() {
		// Recover must be called within a deferred function. When the enclosing function
		// panics, the defer will activate and a recover call within it will catch the panic.
		if r := recover(); r != nil {
			fmt.Println("recover", r) // Return value of recover is the error raised in the call to panic
		}
	}()

	doPanic()                           // Simulate some logic that will panic
	fmt.Println("after panic occurred") // This is never printed as panic will trigger the defer function
}
