package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var collection *mongo.Collection
var ctx = context.TODO()



func main()  {
	/*
	fmt.Println("Hello world")
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
	r.Mount("/products", productsResource{}.Routes())
	*/

	client = Connect()
	collection = client.Database("golang").Collection("products")

	DeleteAll()
	InsertOne( Product{"samsung j7", 22000, "phone"})

	var prods []interface{}
	prods = append(prods,  Product{"Redme Note4", 14000, "phone"})
	prods = append(prods , Product{"dell 4543", 60000, "Laptop"})
	InsertMany(prods)

	fmt.Println("Inserting Done !! ")
	RetrieveOne("name", "dell 4543")

	UpdateOne()
	RetrieveMany()
	Disconnect()

	//log.Fatal(http.ListenAndServe(":"+port, r))
}


