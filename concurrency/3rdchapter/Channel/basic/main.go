package main

import "fmt"

// channels in Go are said to be blocking

func main()  {
	var receiveChan <-chan interface{} // receiving channels are to "read from channel" , arrow position left side
	var sendChan chan<- interface{}  // sending channels are to "write to channel" , arrow position right side

	dataStream := make(chan interface{}) // bi-directional channel

	// Valid statements: () as making bi-directional channel uni-directional is valid in go.
	receiveChan = dataStream
	sendChan = dataStream

	go func() {
		sendChan <- "Hello channels!"
	}()
	fmt.Println("In between sending & receiving.")
	fmt.Println(<-receiveChan) // wait until receive
}
