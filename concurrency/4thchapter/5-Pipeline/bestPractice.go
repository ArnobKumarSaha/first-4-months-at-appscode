package main

/*
There are two types of pipelines :
Batch processing,  stream processing
 */

import (
	"fmt"
	"time"
)

// The generator function takes in a variadic slice of integers, constructs a
//buffered channel of integers with a length equal to the incoming integer slice,
//starts a goroutine, and returns the constructed channel. Then, on the goroutine
//that was created, generator ranges over the variadic slice that was passed in
//and sends the slices’ values on the channel it created.

// In short, it converts a discrete set of values into a stream of data on a channel
var generator = func(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, i := range integers {
			select {
			case <-done:
				fmt.Println("Done Generating")
				return
			case intStream <- i:
			}
		}
	}()
	return intStream
}

// generator takes slice of integers , and gives int stream
// where, multiply takes and gives int stream
var multiply = func(done <-chan interface{}, intStream <-chan int, multiplier int, ) <-chan int {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for i := range intStream {
			select {
			case <-done:
				fmt.Println("Done Multiplying")
				return
			case multipliedStream <- i*multiplier:
			}
		}
	}()
	return multipliedStream
}

// Almost everything, except the functionality, is same in multiply and add function
var add = func(done <-chan interface{}, intStream <-chan int, additive int, ) <-chan int {
	addedStream := make(chan int)
	go func() {
		defer close(addedStream)
		for i := range intStream {
			select {
			case <-done:
				fmt.Println("Done Adding")
				return
			case addedStream <- i+additive:
			}
		}
	}()
	return addedStream
}

func main() {
	done := make(chan interface{})
	//defer close(done)
	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for v := range pipeline {
		fmt.Println(v)
	}
	fmt.Println("main() function completed ! ")
	close(done)
	time.Sleep(2 * time.Second)
	// // done <- 5
}

/*
Two important things :
1) We can use range statement to extract values.
2) each stage of the pipeline is executing concurrently
 */