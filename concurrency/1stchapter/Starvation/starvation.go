package main

import (
	"fmt"
	"sync"
	"time"
)
/*
Starvation is any situation where a concurrent process cannot get all the resources it needs to perform work.
*/

var wg sync.WaitGroup
var sharedLock sync.Mutex
const runtime = 1*time.Second

// The greedy worker greedily holds onto the shared lock for the entirety of its work loop
var greedyWorker = func() {
	defer wg.Done()
	var count int
	for begin := time.Now(); time.Since(begin) <= runtime; {
		sharedLock.Lock()
		time.Sleep(3*time.Nanosecond)
		sharedLock.Unlock()
		count++
	}
	fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
}

// the polite worker attempts to only lock when it needs to.
var politeWorker = func() {
	defer wg.Done()
	var count int
	for begin := time.Now(); time.Since(begin) <= runtime; {
		sharedLock.Lock()
		time.Sleep(1*time.Nanosecond)
		sharedLock.Unlock()
		sharedLock.Lock()
		time.Sleep(1*time.Nanosecond)
		sharedLock.Unlock()
		sharedLock.Lock()
		time.Sleep(1*time.Nanosecond)
		sharedLock.Unlock()
		count++
	}
	fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
}


func main()  {
	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}

/*
we can conclude that the greedy worker has unnecessarily expanded its hold on the shared lock beyond its critical section
and is preventing (via starvation) the polite worker’s goroutine from performing work efficiently.
 */

