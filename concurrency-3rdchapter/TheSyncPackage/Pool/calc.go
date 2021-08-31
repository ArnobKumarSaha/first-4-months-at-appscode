package main

import (
	"fmt"
	"sync"
)

var numCalcsCreated int

var calcPool = &sync.Pool {
	New: func() interface{} {
		numCalcsCreated += 1
		mem := make([]byte, 1024)
		return &mem
	},
}


func main()  {
	// Seed the pool with 4KB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	const numWorkers = 1024*1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}
	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
	// Only 4 calculators created.
	// Though Get() called 1024*1024 times.
}
