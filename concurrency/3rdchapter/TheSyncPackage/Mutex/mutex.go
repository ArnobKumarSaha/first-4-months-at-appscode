package main

import (
	"fmt"
	"sync"
)

// Mutex stands for “mutual exclusion”
// It provides a concurrent-safe way to express exclusive access to these shared resources

var count int
var lock sync.Mutex

var increment = func() {
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Printf("Incrementing: %d\n", count)
}
var decrement = func() {
	lock.Lock()
	defer lock.Unlock()
	count--
	fmt.Printf("Decrementing: %d\n", count)
}

func main()  {
	var arithmetic sync.WaitGroup

	// Scheduling 6 Increment() here.
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// Scheduling 6 Decrement() here.
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
	// Though , increment go routines are called before the decrement go routines ,
	// But we can't say surely an increment routine comes before a decrement routine or not.
	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")

}
