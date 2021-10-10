package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This is a basic generator function which we have already seen in bestPractice.go
var repeat = func(
	done <-chan interface{},
	values ...interface{},
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					fmt.Println("Done called in repeat().")
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

// If num=10 in take(), only 10+1 number will be generated in repeat function.
var take = func(
	done <-chan interface{},
	valueStream <-chan interface{},
	num int,
) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				fmt.Println("Done called in take().")
				return
				// it means, Get the value from valueStream, then pass it to takeStream
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

// repeatFn calls fn() infinitely, until stopped
var repeatFn = func(
	done <-chan interface{},
	fn func() interface{},
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				fmt.Println("Done called in repeatFn().")
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

// This is simply a converter. it will take a channel of interface{] , & make it type-casted.
var toString = func(
	done <-chan interface{},
	valueStream <-chan interface{},
) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for v := range valueStream {
			select {
			case <-done:
				fmt.Println("Done called in toString().")
				return
			case stringStream <- v.(string):
			}
		}
	}()
	return stringStream
}

func main() {
	done := make(chan interface{})
	//defer close(done)

	fmt.Println("Using Repeat function : ")
	for num := range take(done, repeat(done, 1,2,3), 10) {
		fmt.Printf("%v ", num)
	}

	fmt.Println("\nUsing RepeatFn function : ")
	rand := func() interface{} {
		// if the function is expensive, it takes a lot of time , as this function is not running in concurrently.
		// This is running like .... valueStream <- fn()
		// time.Sleep(5 * time.Second)
		return rand.Int()
	}
	// passing our rand function as fn on the repeatFn()
	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Println(num)
	}

	fmt.Println("Using toString function : ")
	var message string
	for token := range toString(done, take(done, repeat(done, "I", "am."), 10)) {
		message += token
	}
	fmt.Printf("message: %s...\n", message)
	close(done)
	time.Sleep(4 * time.Second)
}

/* Why "Done called in repeat()" only printed (comment out from line number 95 to 112)? answer from stackoverflow ::

because it takes only 10 elements (line 115) then it closes its stream (line 38) which in turn closes the loop of toString (line 80)
which closes the loop in main (line 115)which in turn allows for close(done) to happen, at that point only the repeater is still trying to push.
 */