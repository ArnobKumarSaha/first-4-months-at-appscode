package main

import (
	"fmt"
	"math/rand"
)

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
}
