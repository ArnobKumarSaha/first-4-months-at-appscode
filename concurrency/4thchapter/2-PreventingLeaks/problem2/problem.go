package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This problem is almost similar with problem 1 , except
// in problem-1 , the child process was a reader, that waits for receiving
// here, the child process is a writer, that waits for sending

// so, the solution is also too similar

var newRandStream = func() <-chan int {
	randStream := make(chan int)
	go func() {
		// This will never be printed.
		defer fmt.Println("newRandStream closure exited.")
		defer close(randStream)
		for {
			randStream <- rand.Int()
		}
	}()
	return randStream
}
func main()  {
	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	time.Sleep(3 * time.Second)
}
