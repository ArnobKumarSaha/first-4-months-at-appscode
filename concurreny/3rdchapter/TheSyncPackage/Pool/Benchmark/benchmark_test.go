package Benchmark

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

// Not getting the sequence .  Need to look back

func connectToService() interface{} {
	fmt.Println("connectToService() is called.")
	time.Sleep(1*time.Second)
	return struct{}{}
}
func startNetworkDaemon() *sync.WaitGroup {
	fmt.Println("startNetworkDaemon() is called.")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()
		wg.Done()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			connectToService()
			fmt.Fprintln(conn , "connec = ")
			conn.Close()
		}
	}()
	return &wg
}
func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}
func BenchmarkNetworkRequest(b *testing.B) {
	fmt.Println("BenchmarkNetworkRequest() is called.")
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}
