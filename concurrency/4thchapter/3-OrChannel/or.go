package main

// Recursive function in go :: 2 main ways

/* Way-1 (function):
func name(parameter) return_type{
	// do something
}
*/

/* Way-2 (closure):
var name func(parameter) return_type

name = func(parameter) return_type{
	// do something
}
*/


import (
	"fmt"
	"time"
)


// When you want to combine one or more done channels into a single done channel

// The select statement could be used in this case.  But if you don't know the total number of done channel , then
// using or-channel is the only solution.

func or(channels  ...<-chan interface{}) <-chan interface{} {
	// Termination criteria. aka Base Cases
	switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
	}

	orDone := make(chan interface{})

	go func() {
		defer close(orDone)
		switch len(channels) {
			case 2:
				select {
					case <-channels[0]:
					case <-channels[1]:
				}
			default:
				// Code is here means, len(channels) >= 3.
				// In that case, we can recursively call
				select {
					case <-channels[0]:
					case <-channels[1]:
					case <-channels[2]:
						// look , we are appending our orDone channel to the other remaining channels
					case <-or(append(channels[3:], orDone)...):
				}
		}
	}()
	return orDone
}

func main()  {
	// waitAPeriod is a very simple function , which will wait a certain period of time.
	waitAPeriod:= func(after time.Duration) <-chan interface{}{
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	
	start := time.Now()
	<-or(  // waiting , for someone to be executed.
		waitAPeriod(2*time.Hour),
		waitAPeriod(5*time.Minute),
		waitAPeriod(1*time.Second),  // this is the winner
		waitAPeriod(1*time.Hour),
		waitAPeriod(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}

