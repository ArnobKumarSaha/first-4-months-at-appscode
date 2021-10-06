package main

import (
	"fmt"
	"runtime"
	"sync"
)


func main() {

	memConsumed := func() uint64 {
		// GC runs a garbage collection and blocks the caller until the
		// garbage collection is complete. It may also block the entire program.
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <-chan interface{}
	var wg sync.WaitGroup

	// this goroutine won’t exit until the process is finished
	noop := func() { wg.Done(); <-c }


	// Now , we will call noop 1e5 times , & compare the difference between before doing those calls & after doing those calls.
	const numGoroutines = 1e5
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Println(before , " " , after)
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)

	// In my run ... They print
	//71453456   283240504
	//2.118kb⏎
}