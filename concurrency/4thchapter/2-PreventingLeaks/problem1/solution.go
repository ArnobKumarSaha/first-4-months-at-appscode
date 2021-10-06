package main

import (
	"fmt"
	"time"
)
// `done <-chan interface{}` is the first parameter by convention

var doSomeWork = func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
	terminated := make(chan interface{})
	go func() {
		defer fmt.Println("doWork exited.")
		defer close(terminated)
		for {
			select {
			case s := <-strings:
				// Do something interesting
				fmt.Println(s)
			case <-done:
				return
			}
		}
	}()
	return terminated
}

func main() {
	// Another channel is being used to control the child routine from the parent routine.
	done := make(chan interface{})
	terminated := doSomeWork(done, nil)
	go func() {
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()
	<-terminated
	fmt.Println("Done.")
}
