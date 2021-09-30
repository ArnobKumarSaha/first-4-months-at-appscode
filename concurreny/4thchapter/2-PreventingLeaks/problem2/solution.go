package main

import (
	"fmt"
	"math/rand"
	"time"
)

var anotherRandStream = func(done <-chan interface{}) <-chan int {
	randStream := make(chan int)
	go func() {
		defer fmt.Println("newRandStream closure exited.")
		defer close(randStream)
		for {
			select {
			case randStream <- rand.Int():
			case <-done: // It will be the case , after executing line number 32
				return
			}
		}
	}()
	return randStream
}

/*
If a goroutine is responsible for creating a goroutine, it is also
responsible for ensuring it can stop the goroutine.
 */
func main()  {
	done := make(chan interface{})
	randStream := anotherRandStream(done)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	// Simulate ongoing work
	time.Sleep(1 * time.Second)
}

