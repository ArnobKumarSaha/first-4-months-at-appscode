package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Say, We have a queue of fixed length 2, and 10 items we want to push onto the
queue. We want to enqueue items as soon as there is room, so we want to be
notified as soon as there’s room in the queue.
*/
func main()  {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)

		c.L.Lock() // for critical section
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()

		// cond has only 2 methods : Signal() to send signal in one GoRoutine & Broadcast() to all go routines.
		c.Signal()  // inform that , something has occurred
	}
	for i := 0; i < 10; i++{
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1*time.Second)
		c.L.Unlock()
	}
}
