package main

import (
	"fmt"
	"sync"
)

/*
Get will first check whether there are any available instances within the pool to return to the
caller, and if not, call its New member variable to create a new one
 */

//When finished, callers call Put to place the instance they were working with back in the pool for use by other processes

var myPool = &sync.Pool{
	New: func() interface{} {
		fmt.Println("Creating new instance.")
		return struct{}{}
	},
}
func main() {
	myPool.Get() // Garbage collector removes this instances
	// that's why a second call to New() will be occurred.
	fmt.Println("After 1st call to Get().")

	instance := myPool.Get()
	myPool.Put(instance)


	fmt.Println("After 2nd call to Get().")
	myPool.Get()
}
