package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InsertOne(obj Product) {
	res, err := collection.InsertOne(ctx, obj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted document: ", res.InsertedID)
}

func InsertMany( prods []interface{}) {
	insertManyResult, err := collection.InsertMany(context.TODO(), prods)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

func RetrieveOne(key string, value string) {
	var result Product
	filter := bson.D{{key, value}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
}

func RetrieveMany()  {
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(2)

	// Here's an array in which you can store the decoded documents
	var results []*Product

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var elem Product
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
		fmt.Println("count++ ", elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(ctx)

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
}

func DeleteAll()  {
	deleteResult, err := collection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the products collection\n", deleteResult.DeletedCount)
}

func UpdateOne()  {
	filter := bson.D{{"name", "dell 4543"}}
	update := bson.D{
		{"$inc", bson.D{
			{"price", 100},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
