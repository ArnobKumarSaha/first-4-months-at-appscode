package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error error
	Response *http.Response
}

// Sending http.Response only from the channel doesn't give us detailed information about an error.
// So , in this code, we are sending back a struct to the calling function.
// and let that func to handle it properly.

var checkStatus = func(done <-chan interface{}, urls ...string) <-chan Result {
	results := make(chan Result)
	go func() {
		defer close(results)
		for _, url := range urls {
			var result Result
			resp, err := http.Get(url)
			result = Result{Error: err, Response: resp}
			select {
			case <-done:
				return
			case results <- result:
			}
		}
	}()
	return results
}

/*
we’ve successfully separated the concerns of error handling from our producer goroutine.
*/

func main()  {
	done := make(chan interface{})
	defer close(done)
	errCount := 0
	urls := []string{"a", "https://www.google.com", "b", "https://badhost", "d"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("Too many errors, breaking!")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}

