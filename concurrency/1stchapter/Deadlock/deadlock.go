package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu sync.Mutex
	val int
}
var wg sync.WaitGroup

func main()  {
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()
		time.Sleep(2*time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Printf("sum=%v\n", v1.val + v2.val)
	}
	var a, b value

	wg.Add(2)

	/*
	our first call to printSum locks a and then attempts to lock b , but in the meantime our
	second call to printSum has locked b and has attempted to lock a . Both goroutines wait infinitely on each other.
	*/
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}

/*  COFFMAN conditions for Deadlocks :
1) Mutual Exclusion
A concurrent process holds exclusive rights to a resource at any one time.
2) Wait For Condition
A concurrent process must simultaneously hold a resource and be waiting for an additional resource.
3) No Preemption
A resource held by a concurrent process can only be released by that process, so it fulfills this condition.
4) Circular Wait
A concurrent process (P1) must be waiting on a chain of other concurrent processes (P2), which are in turn waiting on it (P1),
so it fulfills this final condition too.
 */

