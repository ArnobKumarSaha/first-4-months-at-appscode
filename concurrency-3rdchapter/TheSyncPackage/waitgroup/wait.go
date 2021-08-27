package main

import (
	"fmt"
	"sync"
)

//we call Done() using the defer keyword to ensure that before we exit
//the goroutine’s closure, we indicate to the WaitGroup that we’ve exited.
var hello = func(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Hello from %v!\n", id)
}

func main()  {
	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters) // adding go routines all at once is also possible.
	// But, calling wg.Add() should be just before the go Routine call.
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
