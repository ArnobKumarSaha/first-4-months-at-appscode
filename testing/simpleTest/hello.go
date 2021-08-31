package main

import "fmt"

func hello(username string) string {
	if len(username) != 0 {
		return fmt.Sprintf("Hello %v!", username)
	} else{
		return "Hello Dude!"
	}
	return ""
}
