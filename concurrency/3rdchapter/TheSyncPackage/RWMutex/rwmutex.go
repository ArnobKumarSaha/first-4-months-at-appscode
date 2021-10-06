package main

// A RWMutex is a reader/writer mutual exclusion lock. The lock can be held by an arbitrary number of readers or a single writer.
// In other words, readers don't have to wait for each other. They only have to wait for writers holding the lock.

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

//type sync.Locker is a interface. It has two methods, Lock and Unlock , which the Mutex and RWMutex types satisfy.
// So basically , RwMutex is a sync.Locker

var producer = func(wg *sync.WaitGroup, l sync.Locker) {
	defer wg.Done()
	for i := 5; i > 0; i-- {
		l.Lock()
		l.Unlock()
		time.Sleep(1)  // to make less active compared to observer
	}
}
var observer = func(wg *sync.WaitGroup, l sync.Locker) {
	defer wg.Done()
	l.Lock()
	defer l.Unlock()
}
var test = func(count int, mutex, rwMutex sync.Locker) time.Duration {
	var wg sync.WaitGroup
	wg.Add(count+1) // count times for observer & 1 time for producer
	beginTestTime := time.Now()
	go producer(&wg, mutex)
	for i := count; i > 0; i-- {
		go observer(&wg, rwMutex)
	}
	wg.Wait()
	return time.Since(beginTestTime)
}

//You can request a lock for reading, in which case you will be
//granted access unless the lock is being held for writing. This means that an
//arbitrary number of readers can hold a reader lock so long as nothing else is
//holding a writer lock

func main()  {
	// tabwriter has been used just make the output text properly aligned.
	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}



/*
Readers  RWMutex      Mutex
1        19.281µs      2.253µs
2        2.719µs       3.029µs
4        24.23µs       16.066µs
8        5.892µs       3.561µs
16       7.748µs       4.942µs
32       34.752µs      27.278µs
64       97.714µs      65.045µs
128      108.574µs     39.022µs
256      74.42µs       65.818µs
512      156.359µs     160.922µs
1024     1.067258ms    309.361µs
2048     1.60263ms     637.024µs
4096     1.698222ms    1.330635ms
8192     2.485068ms    2.098831ms
16384    3.976958ms    4.073832ms
32768    7.861785ms    7.601204ms
65536    15.223873ms   15.254524ms
131072   30.706914ms   32.050487ms
262144   65.319664ms   60.652329ms
524288   121.834185ms  125.398681ms
*/
// ugi-agzbn-7o