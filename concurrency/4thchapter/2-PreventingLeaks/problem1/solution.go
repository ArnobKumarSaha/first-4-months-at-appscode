package main

import (
	"fmt"
	"time"
)
// `done <-chan interface{}` is the first parameter by convention

var doSomeWork = func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
	endConnection := make(chan interface{})
	go func() {
		defer fmt.Println("doWork exited.")
		defer close(endConnection)
		for {
			select {
			case s := <-strings:
				// Do something interesting
				fmt.Println(s)
			case <-done:
				return
				// its returning means calling close(endConnection) function
			}
		}
	}()
	return endConnection
}

// This system is simple,
// use a channel to give the exit-message from parent to the child process  by passing that channel to the child.
// use another channel, from child to parent , to wait the execution in parent.

func main() {
	// Another channel is being used to control the child routine from the parent routine.
	done := make(chan interface{})
	terminated := doSomeWork(done, nil)
	go func() {
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)  // this close message will be received in the for select loop.  [[ case <- done ]]
	}()

	// Code execution will wait in this line
	<-terminated

	fmt.Println("Done.")
	//time.Sleep(4 * time.Second)
}
