package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
Have you ever been in a hallway walking toward another person? She moves to one side to let you pass, but you’ve just done the same. So you move to the
other side, but she’s also done the same. Imagine this going on forever, and you understand livelocks
 */

func main() {
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	//each person must move at the same rate of speed, or cadence. takeStep simulates a constant cadence between all parties.
	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	// A common function used by both tryLeft & tryRight
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)
		atomic.AddInt32(dir, 1)
		fmt.Fprintf(out , " * %v", *dir)
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!")
			return true
		}
		takeStep()

		//the person realizes they cannot go in this direction and gives up. We indicate this by decrementing that direction by one
		atomic.AddInt32(dir, -1)
		return false
	}
	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool {
		return tryDir("left", &left, out)
	}
	tryRight := func(out *bytes.Buffer) bool {
		return tryDir("right", &right, out)
	}

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() { fmt.Println(out.String()) }()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}


	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()

}