package main

import (
	"fmt"
	"time"
)

// Problem : doWork() will remain in memory for the lifetime of this process.
// bcz it will not be garbage collected, as it's execution is still on-going

var doWork = func(strings <-chan string) <-chan interface{} {
	completed := make(chan interface{})
	go func() {
		defer fmt.Println("doWork exited.")
		defer close(completed)

		// the strings channel will never actually gets any strings written onto it, as called with nil
		for s := range strings {
			// Do something interesting
			fmt.Println(s)
		}
	}()
	return completed
}

func main()  {

	doWork(nil)
	time.Sleep(10 * time.Second)  // It doesn't matter how much time I sleep here.
	// "doWork exited." will never be printed out.
	fmt.Println("Done.")
}
