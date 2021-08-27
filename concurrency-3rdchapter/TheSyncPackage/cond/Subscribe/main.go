package main

import (
	"fmt"
	"sync"
	"time"
)

type Button struct {
	Clicked *sync.Cond
}

var subscribe = func(name string , c *sync.Cond, fn func()) {
	var wg sync.WaitGroup
	fmt.Println(name)
	wg.Add(1)
	go func() {
		wg.Done()
		c.L.Lock()
		defer c.L.Unlock()
		c.Wait()  // waiting for Broadcast
		fn()
	}()
	wg.Wait()
}

func main()  {
	button := Button{ Clicked: sync.NewCond(&sync.Mutex{}) }

	// WaitGroup used, only to ensure our program doesn’t exit before our writes to stdout occur
	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)



	subscribe("01",button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	subscribe("02",button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})
	subscribe("03",button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})


	time.Sleep(3 * time.Second) // to visualize what is happening here

	button.Clicked.Broadcast() // Mouse button has been clicked , inform all the awaited routines.
	clickRegistered.Wait()
}
/*
Output :
01
02
03
Mouse clicked.
Maximizing window.
Displaying annoying dialog box!
*/