package main

import (
	"fmt"
	"sync"
	"time"
)
/*
writes to a channel block if a channel is full, and
reads from a channel block if the channel is empty
 */

func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)  // Here we ensure that the channel is closed before we exit the goroutine.
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()
	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
	fmt.Println("Another example starts !")
	example()
}

/*
If you have n goroutines waiting on a single channel, instead
of writing n times to the channel to unblock each goroutine, you can simply
close the channel. Since a closed channel can be read from an infinite number
of times, it doesn’t matter how many goroutines are waiting on it
 */
func example()  {
	begin := make(chan interface{})
	var wg sync.WaitGroup


	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin  // Here the goroutine waits until it is told it can continue. (This is a 'Read from channel')
			//begin <- i
			fmt.Printf("%v has begun\n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutines...")
	/*fmt.Println(<-begin)
	fmt.Println(<-begin)
	fmt.Println(<-begin)
	fmt.Println(<-begin)
	fmt.Println(<-begin)
	fmt.Println(<-begin)
	fmt.Println()*/
	time.Sleep(5 * time.Second)
	close(begin) // all the goroutines will wait for this to be done.
	fmt.Println("channel already closed.")
	wg.Wait()
}