package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func listen(name string, a map[string]int, c *sync.Cond) {
	c.L.Lock()
	c.Wait()
	fmt.Println(name, " age:", a["T"])
	c.L.Unlock()
}

func broadcast(name string, a map[string]int, c *sync.Cond) {
	// time.Sleep(time.Second) is added inside the broadcast() to ensure that the cond.Broadcast() is called
	// after the cond.Wait() functions are called first in the listen() functions which are called in separate goroutines.
	time.Sleep(5 * time.Second)
	c.L.Lock()
	a["T"] = 25
	c.Broadcast()  // Broadcast wakes all goroutines waiting on c.
	c.L.Unlock()
}

func main() {
	// As, map are reference type, it is perfect for this example, otherwise a copy of map would be passed.
	var age = make(map[string]int)

	m := sync.Mutex{}
	cond := sync.NewCond(&m)

	// 2 listeners
	go listen("lis1", age, cond)
	go listen("lis2", age, cond)

	// these listeners will wait forever for cond.Broadcast
	// It has been achieved by c.Wait()

	go broadcast("b1", age, cond)

	// os.Interupt used, so that we can end the program from the terminal using ctrl+c
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}