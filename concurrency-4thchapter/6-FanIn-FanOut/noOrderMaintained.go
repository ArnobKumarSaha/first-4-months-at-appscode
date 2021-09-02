package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// This is to combine tha working channels
var fanIn = func(
	done <-chan interface{},
	channels ...<-chan interface{},
) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})

	// it writes values come from multiple channels , in the multiplexed Stream.
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}
	// Select from all the channels
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}
	// Wait for all the reads to complete
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}

// Make a stream of random integers by internally executing rand() function defined in main
var repeatFn = func(
	done <-chan interface{},
	fn func() interface{},
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

var take = func(
	done <-chan interface{},
	valueStream <-chan interface{},
	num int,
) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

// Simply typecasting the general valueStream to int
var toInt = func(
	done <-chan interface{},
	valueStream <-chan interface{},
) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case intStream <- v.(int):
			}
		}
	}()
	return intStream
}

var primeFinder = func(done <- chan interface{} , intStream <- chan int) <- chan interface{}{
	primeStream := make(chan interface{})

	go func() {
		defer close(primeStream)
		for v := range intStream{
			select {
			case <-done:
				return
			default:
				// Running a heavy primefinding algorithm here.
				flag := true
				for i:=2 ; i<v; i+=1{
					if v%i==0 {
						flag = false
					}
				}
				// writting the value as a prime in the primeStream
				if flag {
					primeStream <- v
				}
			}
		}
	}()

	return primeStream
}

func main() {
	// IN EFFICIENT WAY

	/*
	done := make(chan interface{})
	defer close(done)
	start := time.Now()
	rand := func() interface{} { return rand.Intn(50000000) }
	randIntStream := toInt(done, repeatFn(done, rand))

	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
*/

	// EFFICIENT WAY
	done := make(chan interface{})
	defer close(done)
	start := time.Now()
	rand := func() interface{} { return rand.Intn(50000000) }
	randIntStream := toInt(done, repeatFn(done, rand))
	// Above code is same in both way

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}
	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}
/*
In my case ,  inefficient way takes 41.6s & efficient way takes 15.7s.  Ran on 4 CPU.
 */