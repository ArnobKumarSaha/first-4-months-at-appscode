package main

// pass it a channel to read from, and it will return two
//separate channels that will get the same value

/*
var tee = func(
	done <-chan interface{},
	in <-chan interface{},
) (_, _ <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})

	go func() {
		defer close(out1)
		defer close(out2)

		// orDone func is in 7-OrDoneChannel
		for val := range orDone(done, in) {
			// We will want to use local versions of out1 and out2 , so we shadow these variables.
			var out1, out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
				case out1 <- val:
					out1 = nil
					// Once we’ve written to a channel, we set its shadowed copy to nil so
					//that further writes will block and the other channel may continue.
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()

	return out1, out2
}
*/

func main()  {
	/*
	done := make(chan interface{})
	defer close(done)

	// repeat & take func are in 5-pipeline/Handygenerator
	out1, out2 := tee(done, take(done, repeat(done, 1, 2), 4))
	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}*/
}