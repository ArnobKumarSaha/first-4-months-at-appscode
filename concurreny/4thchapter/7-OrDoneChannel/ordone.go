package main

/*
// The old approach.  (it can't prevent goroutine leaking)
for val := range myChan {
		// Do something with val
	}


// A cumbersome approach.
// It can prevent , but code gets horrible. As we have to write for-select statements wherever we need to prevent goroutine leak.
loop:
	for {
		select {
		case <-done:
			break loop
		case maybeVal, ok := <-myChan:
			if ok == false {
				return // or maybe break from for
			}
			// Do something with val
		}
	}

*/

// A very standard Approach for preventing GoRoutine leaks
var orDone = func(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func main() {
	// Now , We can simply call like this .. , as we did in the older approaches ..
	/*
	for val := range orDone(done, myChan) {
		// Do something with val
	}
	*/
}
