package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/homepage", trackVisits(homepageHandler))
	http.HandleFunc("/projects", trackVisits(projectsHandler))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
var globalVisit int

func trackVisits(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	totalVisit := 0
	return func(writer http.ResponseWriter, request *http.Request) {
		// track the visit
		totalVisit++
		globalVisit++
		fmt.Println("local visit : ", totalVisit, ", Global visit ", globalVisit)
		// call the original handler
		handler(writer, request)
	}
}

func homepageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to my homepage")
}
func projectsHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "These are my projects")
}
//http.HandleFunc("/projects", authenticate(trackVisits(projectsHandler)))
