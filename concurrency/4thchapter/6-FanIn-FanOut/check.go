package main

/*
This code has nothing to do with fan-in-fan-out
I was just checking if calling sleep from main can have any affect on other go routines or not...
The answer is no.
 */
import (
	"fmt"
	"time"
)

func fun()  {
	for{
		fmt.Println("having lot of fun ! ")
	}
}
func main()  {
	go fun()
	c1 , c2 := 0, 0
	for i:=0; i<10000; i+=1 {
		c1 += 1
	}
	time.Sleep(3 * time.Second)
	fmt.Println(c1, "****************************************************************************************************************")
	for i:=0; i<100000; i+=1 {
		c2 += 1
	}
	fmt.Println(c2, "****************************************************************************************************************")
}