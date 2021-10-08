package main

import (
	"fmt"
	called "golang_k8s/basics/calledpackage"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// This block will call the Hellos function declared in calledPackage/greetings.go
	names := []string{"Gladys", "Samantha", "Darrin"} // A slice of names.
	message, err := called.Hellos(names)
	// If an error was returned, print it to the console and exit the program.
	// log.Fatal function call os.Exit(1) internally.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message nto the console.
	fmt.Println("Hellos : ", message)

	// This block will call the Hello function declared in calledPackage/greeting.go
	another_message, error := called.Hello("Arnob")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Hello : ", another_message)
}


// The last line on go.mod file was created when syncing
// In package name, only characters, digits and underscore are allowed.