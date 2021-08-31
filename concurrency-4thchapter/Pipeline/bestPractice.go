package main

import (
	"fmt"
	"time"
)

// Just to write the integer values to a readonly channel.
// it converts a discrete set of values into a stream of data on a channel
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
	time.Sleep(4 * time.Second)
	done <- 5
	close(done)
	time.Sleep(1 * time.Second)
}