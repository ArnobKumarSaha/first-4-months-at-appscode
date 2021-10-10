package main

import (
	"fmt"
	"time"
)

func main() {
	data := make([]int, 4)
	// loopData is the writter function which writes data to the channel
	// It is confined only to write.
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}
	handleData := make(chan int)
	go loopData(handleData)
	// And this is the receiver part
	for num := range handleData {
		fmt.Println(num)
	}





	// EXAMPLE 2  :

	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		fmt.Println("doWork() has started.")

		go func() {  // It is confined only to receive.
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				fmt.Println(s, "TTT")
			}
		}()

		return completed
	}
	stringChan := make(chan string)
	//stringChan <- "a"   // This will make a deadlock ! As there is still no receiver.
	doWork(stringChan)
	// Perhaps more work is done here
	stringChan <- "b"
	fmt.Println("Done.")

	// time.sleep has been used just to see if 'c' has been printed or not, as main() is exiting before that print.
	stringChan <- "c"
	time.Sleep(time.Second/2)
}