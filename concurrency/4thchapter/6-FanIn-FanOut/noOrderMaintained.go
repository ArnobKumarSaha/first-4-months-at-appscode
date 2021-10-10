package main

/*
Fans-in-fans-out ===  Reuse a single stage of pipeline on multiple
goroutines in an attempt to parallelize pulls from an upstream stage

fans-out == the process of starting multiple goroutines to handle input from the pipeline
fans-in == the process of combining multiple results into one channel.
*/

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
				fmt.Println("Done called in fanIn().")
				return
			case multiplexedStream <- i:
			}
		}
	}
	// Collect results from all the channels
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}
	// Wait for all the reads to complete
	go func() {
		wg.Wait()
		fmt.Println("multiplexedStream is being closed.")
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
				fmt.Println("Done called in repeatFn().")
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
				fmt.Println("Done called in take().")
				return
			case takeStream <- <-valueStream:
			}
		}
		fmt.Println("TakeStream is being closed.")
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
				fmt.Println("Done called in toInt().")
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
				fmt.Println("Done called in primeFinder().")
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
		fmt.Println("PrimeStream is being closed.")
	}()

	return primeStream
}

func main() {
	done := make(chan interface{})
	//defer close(done)
	start := time.Now()
	rand := func() interface{} { return rand.Intn(50000000) }
	randIntStream := toInt(done, repeatFn(done, rand))

/*   // INEFFICIENT WAY
	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
*/

	// EFFICIENT WAY
	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)   // a slice of interface{}-typed reader-channels
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}
	// In the inefficient method, We called , take(done, primeFinder(done, randIntStream), 10 )
	// But now, as there are several primeFinders, we have to combine those result into one single channel first.
	for prime := range take(done, fanIn(done, finders...), 5) {
		fmt.Printf("\t%d\n", prime)
	}
	c2 := 0
	for i:=0; i<10000000000; i+=1 {
		c2 += 1
	}
	fmt.Println(c2)
	time.Sleep(7 * time.Second)
	fmt.Printf("Search took: %v\n", time.Since(start))
	close(done)
	time.Sleep(5 * time.Second)
	// "Done called" printed in toInt() & repeatFn() respectively.
}

/*
In my case ,  inefficient way takes 41.6s & efficient way takes 15.7s.  Ran on 4 CPU.
 */

/* OUTPUT

Spinning up 4 prime finders.
Primes:
        19727887
        43516159
        38043721
        45071563
TakeStream is being closed.
        49509107
10000000000
Search took: 25.94634619s
Done called in fanIn().
Done called in fanIn().
Done called in repeatFn().
Done called in toInt().
Done called in fanIn().
Done called in fanIn().
multiplexedStream is being closed.
PrimeStream is being closed.
PrimeStream is being closed.

 */