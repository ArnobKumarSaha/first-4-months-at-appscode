package main

import (
	"fmt"
	"time"
)

func  main()  {
	done := make(chan int)
	go func() {
		time.Sleep(3*time.Second)
		close(done)
		//done <- 100
	}()
	workCounter := 0




loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		// Simulate work
		workCounter++
		time.Sleep(1*time.Second)
	}


	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
