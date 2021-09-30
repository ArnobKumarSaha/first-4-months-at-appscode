package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var a []int
	a = append(a, []int{12, 65, 983, 83 , 23}...)

	// --------------- COPYING ------------------
	b := make([]int, len(a))
	copy(b, a)  // option 1 : with copy() function
	b[2] = 0
	fmt.Println(a, b)


	b = append([]int(nil), a...)  // option 2 : with the empty slice and variadic notation
	b[3] = 0
	fmt.Println(a,b)
	println(reflect.TypeOf(a).String(), reflect.TypeOf(b).String())

	println("-------------------------------------------")


	// --------------- CUT ------------------
	var i,j = 1,3
	a = a[:4]
	fmt.Println(a)
	var c []int
	c  = append(a[:i], a[j:]...)
	fmt.Println(c)

	// --------------- DELETE ------------------
	a = []int{12, 65, 983, 83 , 23, 25}
	i = 3
	copy(a[i:], a[i+1:]) // i+1 th element copied on i th, i+2 th element copied on i+1 th & so on
	a = a[:len(a)-1] // exclude the last index
	fmt.Println(a)

	// --------------- EXPAND ------------------
	a = []int{12, 65, 983, 83 , 23, 25}
	i=3
	j=5  // ith index e j ta zero element dhukabe
	a = append(a[:i], append(make([]int, j), a[i:]...)...) 
	fmt.Println(a)

	// --------------- EXTEND ------------------
	a = append(a, make([]int, j)...)
	fmt.Println(a)

	println("-------------------------------------------")



	// --------------- INSERT ------------------
	a = []int{12, 65, 983, 83 , 23, 25}
	i=4
	a = append(a, 0 /* use the zero value of the element type */)
	copy(a[i+1:], a[i:])
	a[i] = 100
	fmt.Println(a)

	// --------------- INSERT VECTOR------------------
	a = []int{12, 65, 983, 83 , 23, 25}
	b = []int{1,2}
	a = append(a[:i], append(b, a[i:]...)...)
	fmt.Println(a,b)


	// --------------- REVERSING ------------------
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	fmt.Println(a)
}

// For more details , Go to   https://github.com/golang/go/wiki/SliceTricks