package main

import (
	"fmt"
	"time"
)

/*
// The old approach.  (it can't prevent goroutine leaking)
for val := range myChan {
		// Do something with val
	}


// A cumbersome approach.
// It can prevent , but code gets horrible. As we have to write for-select statements wherever we need to prevent goroutine leak.
loop:
	for {
		select {
		case <-done:
			break loop
		case maybeVal, ok := <-myChan:
			if ok == false {
				return // or maybe break from for
			}
			// Do something with val
		}
	}

*/

var repeat = func(
	done <-chan interface{},
	limit int,
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for v:=1 ; v<=limit ; v+=1{
			select {
			case <-done:
				fmt.Println("Done called in repeat().")
				return
			case valueStream <- v:
			}
		}
	}()
	return valueStream
}

// A very standard Approach for preventing GoRoutine leaks
var orDone = func(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func main() {
	// Now , We can simply call like this .. , as we did in the older approaches ..
	/*
	for val := range orDone(done, myChan) {
		// Do something with val
	}
	*/


	// This example demonstrates the visualization
	done := make(chan interface{})
	go func() {
		x := 0
		for i:=0; i<1000000000; i+=1 {
			x+=1
		}
		fmt.Println(x)
		close(done)
	}()

	/*
	for val := range repeat(done, 10) {
		fmt.Println(val)
	}
	*/

	for val := range orDone(done, repeat(done, 10)) {
		fmt.Println(val)
	}
	fmt.Println("range loop completed")
	time.Sleep(2 * time.Second)
}
