
package main

/*
This code has errors.
Have to look on :
https://www.newline.co/@kchan/building-a-simple-restful-api-with-go-and-chi--5912c411

*/

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	port := "8080"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Mount("/posts", postsResource{}.Routes())

	log.Fatal(http.ListenAndServe(":"+port, r))
}