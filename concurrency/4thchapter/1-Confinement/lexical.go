package main

// Lexical confinement involves using lexical scope to expose only the correct data and
// concurrency primitives for multiple concurrent processes to use. It makes it impossible to do the wrong thing.

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main()  {
	// Call the functions here
}

func printDataDemo()  {
	// Not concurrent safe
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		// As bytes.Buffer is not concurrent safe, we are using it to demontrate the example
		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}
	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3]) // passing a slice containing first 3 bytes
	go printData(&wg, data[3:]) // passing a slice containing last 3 bytes
	wg.Wait()
}

func ownerConsumerDemo() {
	// Look at the return type .  who receives it's return value .. are only readers !!
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	// Here we receive a read-only copy of an int channel. By declaring that the only
	// usage we require is read access, we confine usage of the channel within the consume function to only reads
	comsumer := func(results <-chan int) {
		for result := range results {
			fmt.Println("Received: %d\n", result)
		}
		fmt.Println("Done Receiving!")
	}

	results := chanOwner()
	comsumer(results)
}

func blockOnAttemptingToWriteToChannel() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure existed.")
			// This will never be printed for the infinite for-loop
			// The main function will always exit first before closing randStream channel.
			defer close(randStream)
			cnt := 1
			for {
				randStream <- rand.Int()
				fmt.Println("I have generated ", cnt, " int.")
				cnt += 1
			} // "I have generated 4 int" will never be printed , no matter how much time is set to sleep().
			 // bcz there is a block on attempting to write , by the below for-loop
		}()
		return randStream
	}
	randStream := newRandStream()
	fmt.Println("3 random ints:")

	// Read only 3 int from the infinite stream & exit.
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	time.Sleep(time.Second * 4)
}

func fixBlockOnAttemptingToWriteToChannel() {
	// The solution, just like for the receiving case, is to provide the
	// producer goroutine with a channel informing it to exit
	d := make(chan interface{})
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			// this will be printed. As the for loop has break-condition now.
			defer fmt.Println("newRandStream closure existed.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}
	randStream := newRandStream(d)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(d)
	// This sleep() is called to show the "clousure exited" messages.
	// otherwise main() exits first
	time.Sleep(1 * time.Second)
}
