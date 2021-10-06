package main

import (
	"sync"
	"testing"
)

// Actually , I have no idea what's going on here.
// I just have taken the code from page 82 of the concurrency book.

// A look back to this code is required.

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})
	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}
	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}

// Command to run this code :
// go test -bench=. -cpu=1

/* OUTPUT
goos: linux
goarch: amd64
pkg: learning-golang/concurrency-3rdchapter/SwitchingTime
cpu: Intel(R) Core(TM) i3-4150 CPU @ 3.50GHz
BenchmarkContextSwitch   7080342               167.0 ns/op
PASS
ok      learning-golang/concurrency-3rdchapter/SwitchingTime    1.356s

*/
