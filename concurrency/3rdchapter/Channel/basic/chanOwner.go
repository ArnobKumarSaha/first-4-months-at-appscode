package main

import (
	"fmt"
	"time"
)

func main()  {
	outputs := make([]string, 0)

	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		//wg := sync.WaitGroup{}
		//wg.Add(1)


		go func() {
			defer close(resultStream)
			//defer wg.Done()
			for i := 1; i <= 10; i++ {
				fmt.Println("sending data starts", i)
				resultStream <- i
				outputs = append(outputs, fmt.Sprintf("sent `%d`", i))
				fmt.Println("sending data ends", i)
			}
		}()
		//wg.Wait()
		return resultStream
	}


	resultStream := chanOwner()
	//fmt.Println("chanOwner() funciton done !")
	time.Sleep(time.Second*5)
	for result := range resultStream {
		outputs = append(outputs, fmt.Sprintf("recieved `%d`", result))
		fmt.Printf("Received: %d\n", result)
	}
	//fmt.Println("Done receiving!")

	//println("\n")

	//for _, output := range outputs {
	//	fmt.Println(output)
	//}
}