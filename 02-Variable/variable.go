package main

import(
	"fmt"
)

// 3 types of Visibility of the Variables
var ExportedVariable string = "This string will be exported outside the module."
var packageScopeVariable string = "its scope is within package"
// another visibility scope is within block. for example.. typeOne variable declared inside the main function block.
// Important thing : there is no private scope. No variables is accessible just in this go file.

// it is possible to make a var block, like this.. 
var (
	i int = 12
	j string = "hello"
)

func main()  {
	// There are three types of variable declaration
	var typeOne int = 45
	var typeTwo int
	typeThree := 62
	typeTwo = 23
	fmt.Println(typeOne, typeTwo, typeThree)
	fmt.Println(packageScopeVariable, i,j)
}