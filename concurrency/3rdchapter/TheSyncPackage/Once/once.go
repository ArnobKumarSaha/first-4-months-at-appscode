package main

import (
	"fmt"
	"sync"
)

/*
the only thing sync.Once guarantees is that your functions are only called once.
Sometimes this is done by deadlocking your program and exposing the flaw
in your logic — in this case a circular reference
 */

func example1()  {
	var count int
	increment := func() {
		count++
	}
	var once sync.Once
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(increment) // Calling it multiple times .. will not affect
		}()
	}
	wg.Wait()
	fmt.Printf("Count is %d\n", count) // Count will remain 1
}
func example2()  {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }
	var once sync.Once

	// Even calling different functions won't have an affect !
	once.Do(increment)
	once.Do(decrement)
	fmt.Printf("Count: %d\n", count)
}
func main()  {
	example1()
	example2()
}