package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
)

var cache redis.Conn

func main() {
	//initCache()
	conn, err := redis.DialURL("redis://localhost")

	fmt.Println("dffsdfsdf", conn)
	if err != nil {
		panic(err)
	}
	cache = conn

	// "Signin" and "Signup" are handler that we will implement
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func initCache() {
	conn, err := redis.DialURL("redis://localhost:8000")
	if err != nil {
		panic(err)
	}
	cache = conn
}