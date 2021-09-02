package main

import (
"fmt"
"sync"
"time"
)

func main() {
	lock := sync.RWMutex{}

	b := map[string]int{}
	b["0"] = 0

	go func(i int) {
		lock.RLock()
		fmt.Printf("RLock: from go routine %d: b = %d\n",i, b["0"])
		time.Sleep(time.Second*3)
		fmt.Printf("RLock: from go routine %d: lock released\n",i)
		lock.RUnlock()
	}(1)

	go func(i int) {
		lock.Lock()
		b["2"] = i
		fmt.Printf("Lock: from go routine %d: b = %d\n",i, b["2"])
		time.Sleep(time.Second*3)
		fmt.Printf("Lock: from go routine %d: lock released\n",i)
		lock.Unlock()
	}(2)

	go func(i int) {
		lock.RLock()
		fmt.Printf("RLock: from go routine %d: b = %d\n",i, b["0"])
		time.Sleep(time.Second*3)
		fmt.Printf("RLock: from go routine %d: lock released\n",i)
		lock.RUnlock()
	}(3)

	go func(i int) {
		lock.RLock()
		fmt.Printf("RLock: from go routine %d: b = %d\n",i, b["0"])
		time.Sleep(time.Second*3)
		fmt.Printf("RLock: from go routine %d: lock released\n",i)
		lock.RUnlock()
	}(4)

	<-time.After(time.Second*8)

	fmt.Println("*************************************8")

	go func(i int) {
		lock.Lock()
		b["3"] = i
		fmt.Printf("Lock: from go routine %d: b = %d\n",i, b["3"])
		time.Sleep(time.Second*3)
		fmt.Printf("Lock: from go routine %d: lock released\n",i)
		lock.Unlock()
	}(5)

	go func(i int) {
		lock.RLock()
		fmt.Printf("RLock: from go routine %d: b = %d\n",i, b["3"])
		time.Sleep(time.Second*3)
		fmt.Printf("RLock: from go routine %d: lock released\n",i)
		lock.RUnlock()
	}(6)

	<-time.After(time.Second*8)
}
