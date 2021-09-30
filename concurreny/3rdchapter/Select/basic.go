package main

import (
	"fmt"
	"time"
)

func main()  {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(2*time.Second)
		close(c)
	}()
	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	// what if we want to continue if nothing is received through channel in 1 second , and get out from the select statement.
	// this is the elegant way , Go support to do that
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.")
	}
}
