package main

import (
	"fmt"
	"sync"
)

func wrongWay()  {
	fmt.Println("Wrong way of using go routine : ")
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

func rightWay()  {
	fmt.Println("Right way of using go routine : ")
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}

func main()  {
	// This will print "good day" 3 times.
	// What is happening is,  the loop exits before any goroutines being running.

/*	But The Go runtime is observant enough to know that a reference to the salutation
	variable is still being held, and therefore will transfer the memory to the heap
	so that the goroutines can continue to access it.*/
	wrongWay()


	// This will print "hello", "greetings", "good day". And certainly , the order will not be assured.
	// passing the copied value of 'salutation' variable will do the work.
	rightWay()
}
