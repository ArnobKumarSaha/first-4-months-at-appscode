package main

import "fmt"

func Add(x , y int) int {
	var v int
	v = x+y
	return v
}
func main()  {
	ret := Add(3,6)
	fmt.Println(ret)
}
